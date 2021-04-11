package repodata

import (
	"encoding/json"
	"strings"
)

type MultiStr []string

func (ms *MultiStr) UnmarshalJSON(b []byte) error {
	var err error
	var s string
	if err = json.Unmarshal(b, &s); err == nil {
		*ms = []string{s}
		return nil
	}

	var strings []string
	if err = json.Unmarshal(b, &strings); err == nil {
		*ms = strings
		return nil
	}

	return err
}

type RepoContentSet struct {
	Name       string   `json:"name"`
	Baseurl    MultiStr `json:"baseurl"`
	Basearch   MultiStr `json:"basearch"`
	Releasever MultiStr `json:"releasever"`
	ThirdParty bool     `json:"third_party"`
}

type RepoProduct struct {
	ContentSets map[string]RepoContentSet `json:"content_sets"`
}

type RepoEntry struct {
	Products map[string]RepoProduct `json:"products"`
}

type Repolist []RepoEntry

func GetUrls(repolist Repolist) map[string][]string {
	repos := map[string][]string{}

	for _, r := range repolist {
		for _, prod := range r.Products {
			for label, cset := range prod.ContentSets {
				candidates := map[string]bool{}
				for _, url := range cset.Baseurl {
					candidates[url] = true
					for _, arch := range cset.Basearch {
						withArch := strings.ReplaceAll(url, "$basearch", arch)
						candidates[withArch] = true
						for _, rv := range cset.Releasever {
							candidates[strings.ReplaceAll(withArch, "$releasever", rv)] = true
						}
					}
				}

				var urls []string
				for url := range candidates {
					if !strings.Contains(url, "$basearch") && !strings.Contains(url, "$releasever") {
						urls = append(urls, url)
					}
				}
				repos[label] = urls
			}
		}
	}
	return repos
}
