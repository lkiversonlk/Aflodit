package dao

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
)

const (
	maximux_record_count = 1000
	ImageCollection = "images"
)

type Database struct {
	db *mgo.Database
}

type Iter struct {
	*mgo.Iter
}

func NewDatabase() *Database{
	return &Database{}
}

func (d *Database) Connect(url, db string) (err error){
	mgoSession, err := mgo.Dial(url)
	d.db = mgoSession.DB(db)
	return err
}

func (d *Database) Query(model string, query map[string] interface{}, sel interface{}, start, count int) ([]bson.M, bool, error){
	var result []bson.M
	if count == 0 {
		count = maximux_record_count
	}

	mongoQuery := d.db.C(model).Find(query).Skip(start).Limit(count)

	if sel != nil {
		mongoQuery = mongoQuery.Select(sel)
	}

	err := mongoQuery.All(&result)

	if err != nil {
		return nil, false, err
	}else {
		more := false
		getCount := len(result)
		if getCount == count {
			more = true
		}
		return result, more, err
	}
}

func (d *Database) Count(model string, query map[string] interface{}) (int, error){
	return d.db.C(model).Find(query).Count()
}

