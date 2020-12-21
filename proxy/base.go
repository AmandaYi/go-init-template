package proxy

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
)

type proxy interface {
	Get(uuid string) (interface{}, error)
	GetOne(query interface{}) (interface{}, error)
	GetMany(query interface{}, page int, limit int) (interface{}, error)
	Create(body interface{}) error
	Update(uuid string, body interface{}) error
	Delete(uuid string) error
	Count(query interface{}) (int, error)
}

var _ proxy = &baseProxy{}

type baseProxy struct {
	table *gorm.DB
	model reflect.Type
}

func (p *baseProxy) Get(uuid string) (interface{},error) {
	var err error
	var data = reflect.New(p.model).Interface()
	p.table.Where("uuid = " + uuid).First(data)
	return reflect.ValueOf(data).Elem().Interface(), err
}

func (p *baseProxy) GetOne(query interface{}) (interface{}, error) {
	panic("implement me")
}

func (p *baseProxy) GetMany(query interface{}, page int, limit int) (interface{}, error) {
	panic("implement me")
}

func (p *baseProxy) Create(body interface{}) error {
	panic("implement me")
}

func (p *baseProxy) Update(uuid string, body interface{}) error {
	panic("implement me")
}

func (p *baseProxy) Delete(uuid string) error {
	panic("implement me")
}

func (p *baseProxy) Count(query interface{}) (int, error) {
	panic("implement me")
}


























































//func (p *baseProxy) GetOne(query interface{}) (interface{}, error) {
//	var data = reflect.New(p.model).Interface()
//	err := p.coll.Find(query).One(data)
//	return reflect.ValueOf(data).Elem().Interface(), err
//}
//
//func (p *baseProxy) GetMany(query interface{}, page int, limit int) (interface{}, error) {
//	// fmt.Println(strconv.Itoa(page))
//
//	var data = reflect.MakeSlice(reflect.SliceOf(p.model), 0, 10).Interface()
//	err := p.coll.Find(query).Sort("-_id").Limit(limit).Skip((page - 1) * limit).All(&data)
//	return data, err
//}
//
//func (p *baseProxy) Create(body interface{}) error {
//	v := reflect.ValueOf(body).Elem()
//	now := time.Now()
//	v.FieldByName("CreatedAt").Set(reflect.ValueOf(now))
//	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(now))
//	v.FieldByName("ID").Set(reflect.ValueOf(bson.NewObjectId()))
//	err := p.coll.Insert(body)
//	return err
//}
//
//func (p *baseProxy) Update(id string, body interface{}) error {
//	v := reflect.ValueOf(body).Elem()
//	now := time.Now()
//	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(now))
//	err := p.coll.UpdateId(bson.ObjectIdHex(id), bson.M{
//		"$set": body,
//	})
//	return err
//}
//
//func (p *baseProxy) Delete(id string) error {
//	err := p.coll.RemoveId(bson.ObjectIdHex(id))
//	return err
//}
//
//func (p *baseProxy) Count(query interface{}) (int, error) {
//	n, err := p.coll.Find(query).Count()
//	return n, err
//}
//
