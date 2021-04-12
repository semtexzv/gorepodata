package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"time"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	inner, _ := DB.DB()

	sql, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		panic(err)
	}
	_, err = inner.Exec(string(sql))
	if err != nil {
		panic(err)
	}
}

type Base struct {
	ID int `gorm:"primaryKey"`
}

type ContentSet struct {
	Base
	Label string
	Name  string
}

type Repo struct {
	Base
	ContentSetID int
	Url          string
	Revision     int
}

type PackageName struct {
	Base
	Name string
}

type Evr struct {
	Base
	EvrData
}

type EvrData struct {
	Epoch   int
	Version string
	Release string
}

type Package struct {
	Base
	NameID int
	EvrID  int
	ArchID int
}

type PackageRepo struct {
	PackageID int `gorm:"primarykey"`
	RepoID    int `gorm:"primarykey"`
}

type AdvisoryType struct {
	Base
	Name string
}

type AdvisorySeverity struct {
	Base
	Name string
}

type Advisory struct {
	Base
	Name string

	TypeID     int
	SeverityID int

	Synopsis    string
	Summary     string
	Description string
	Solution    string

	Issued  time.Time
	Updated time.Time
}

type AdvisoryRepo struct {
	AdvisoryID int `gorm:"primarykey"`
	RepoID     int `gorm:"primarykey"`
}

type AdvisoryPackage struct {
	AdvisoryID int `gorm:"primarykey"`
	PackageID  int `gorm:"primarykey"`
	StreamID   int `gorm:"primarykey"`
}

type Module struct {
	Base
	RepoID int
	Name   string
	ArchID int
}

type ModuleStream struct {
	Base
	ModuleID int
	Name     string
	Version  int
	Context  string
}
