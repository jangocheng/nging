//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type UserU2f struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*UserU2f
	
	Id     	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid    	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Token  	string  	`db:"token" bson:"token" comment:"签名" json:"token" xml:"token"`
	Type   	string  	`db:"type" bson:"type" comment:"类型" json:"type" xml:"type"`
	Extra  	string  	`db:"extra" bson:"extra" comment:"扩展设置" json:"extra" xml:"extra"`
	Created	uint    	`db:"created" bson:"created" comment:"绑定时间" json:"created" xml:"created"`
}

func (this *UserU2f) Trans() *factory.Transaction {
	return this.trans
}

func (this *UserU2f) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *UserU2f) Objects() []*UserU2f {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *UserU2f) NewObjects() *[]*UserU2f {
	this.objects = []*UserU2f{}
	return &this.objects
}

func (this *UserU2f) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("user_u2f").SetModel(this)
}

func (this *UserU2f) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *UserU2f) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *UserU2f) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *UserU2f) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *UserU2f) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *UserU2f) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *UserU2f) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Update()
}

func (this *UserU2f) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *UserU2f) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *UserU2f) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}
