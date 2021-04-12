package main

import (
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"github.com/ulikunitz/xz"
	"gopkg.in/yaml.v3"
	"io"
	"repodata/db"
	"strconv"

	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "repodata/db"
	"repodata/repodata"
	"strings"
)

func Unwrap(err error) {
	if err != nil {
		panic(err)
	}
}

type SyncHandler interface {
	Metadata(md repodata.RepoMD, ctx Syncer) error
	Primary(primary repodata.Primary, ctx Syncer) error
	Updateinfo(updateinfo repodata.Updateinfo, ctx Syncer) error
	ModuleItem(mod repodata.ModuleItem, ctx Syncer) error
}

type Syncer struct {
	base string
}

func NewSyncer(base string) Syncer {
	return Syncer{base}
}
func (ctx Syncer) getItem(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Faile %s", ctx.base)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid")
	}
	stream := res.Body
	if strings.HasSuffix(url, "gz") {
		stream, err = gzip.NewReader(res.Body)
	} else if strings.HasSuffix(url, "bz2") {
		// TODO:
		stream = ioutil.NopCloser(bzip2.NewReader(res.Body))
	} else if strings.HasSuffix(url, "xz") {
		var rdr io.Reader
		rdr, err = xz.NewReader(res.Body)
		stream = ioutil.NopCloser(rdr)
	}
	if stream != nil {
		return ioutil.ReadAll(stream)
	}
	return nil, err
}

func (ctx Syncer) findUrl(md repodata.RepoMD, typ string) string {
	for _, item := range md.Data {
		if item.Type == typ {
			return fmt.Sprintf("%s/%s", ctx.base, item.Location.Href)
		}
	}
	return ""
}

func (ctx Syncer) SyncMetadata(handler SyncHandler) error {
	data, err := ctx.getItem(fmt.Sprintf("%s/repodata/repomd.xml", ctx.base))
	if err != nil {
		return fmt.Errorf("metadata: %v", err)
	}

	var md repodata.RepoMD
	Unwrap(xml.Unmarshal(data, &md))
	return handler.Metadata(md, Syncer{ctx.base})
}

func (ctx Syncer) SyncPrimary(md repodata.RepoMD, handler SyncHandler) error {
	url := ctx.findUrl(md, "primary")
	if url == "" {
		return errors.New("primary not found")
	}

	data, err := ctx.getItem(url)
	if err != nil {
		return fmt.Errorf("primary: %v", err)
	}

	var prim repodata.Primary
	Unwrap(xml.Unmarshal(data, &prim))
	return handler.Primary(prim, ctx)
}

func (ctx Syncer) SyncUpdates(md repodata.RepoMD, handler SyncHandler) error {
	url := ctx.findUrl(md, "updateinfo")
	if url == "" {
		return errors.New("updateinfo not found")
	}

	data, err := ctx.getItem(url)
	if err != nil {
		return fmt.Errorf("updateinfo: %v", err)
	}

	var up repodata.Updateinfo

	if err = xml.Unmarshal(data, &up); err != nil {
		return err
	}
	return handler.Updateinfo(up, ctx)
}

func (ctx Syncer) SyncModules(md repodata.RepoMD, h handler) error {
	url := ctx.findUrl(md, "modules")
	if url == "" {
		return errors.New("modules not found")
	}

	data, err := ctx.getItem(url)
	if err != nil {
		return fmt.Errorf("modules: %v", err)
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	var item repodata.ModuleItem
	for dec.Decode(&item) == nil {
		if err = h.ModuleItem(item, ctx); err != nil {
			return err
		}
	}

	return nil
}

type handler struct {
	repo db.Repo
}

func (h handler) Metadata(md repodata.RepoMD, ctx Syncer) error {
	err := ctx.SyncPrimary(md, h)
	if err != nil {
		return err
	}
	err = ctx.SyncUpdates(md, h)
	if err != nil {
		return err
	}
	err = ctx.SyncModules(md, h)
	if err != nil {
		return err
	}
	return nil
}

func (h handler) Primary(primary repodata.Primary, ctx Syncer) error {
	names := []string{}
	evrs := map[db.EvrData]bool{}

	for _, p := range primary.Package {
		names = append(names, p.Name)
		epoch, _ := strconv.Atoi(p.Version.Epoch)
		evrs[db.EvrData{
			Epoch:   epoch,
			Version: p.Version.Ver,
			Release: p.Version.Rel,
		}] = true
	}
	var nameArr []db.PackageName
	db.DB.Find(&nameArr, "name in (?)", names)

	println("primary done")
	return nil
}
func (h handler) Updateinfo(updateinfo repodata.Updateinfo, ctx Syncer) error {
	return nil
}

func (h handler) ModuleItem(m repodata.ModuleItem, ctx Syncer) error {
	fmt.Printf("Module item %+v", m)
	return nil
}

func main() {
	repo, err := ioutil.ReadFile("./testdata/repolist.json")
	Unwrap(err)
	var repolist repodata.Repolist
	Unwrap(json.Unmarshal(repo, &repolist))
	fmt.Printf("%+v\n", repolist)

	repodata.SyncRepolist(repolist)

	var repos []db.Repo
	db.DB.Find(&repos)
	for _, repo := range repos {
		ctx := NewSyncer(repo.Url)
		err = ctx.SyncMetadata(handler{repo})
		if err != nil {
			fmt.Printf("Error occured: %v \n", err)

		}
	}

}
