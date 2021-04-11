package bindb

import bolt "go.etcd.io/bbolt"

type Db struct {
	db bolt.DB
}

func (db *Db) name()  {

}