// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingDbAccount []*NgingDbAccount

func (s Slice_NgingDbAccount) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingDbAccount) RangeRaw(fn func(m *NgingDbAccount) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingDbAccount) GroupBy(keyField string) map[string][]*NgingDbAccount {
	r := map[string][]*NgingDbAccount{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingDbAccount{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingDbAccount) KeyBy(keyField string) map[string]*NgingDbAccount {
	r := map[string]*NgingDbAccount{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingDbAccount) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingDbAccount) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingDbAccount) FromList(data interface{}) Slice_NgingDbAccount {
	values, ok := data.([]*NgingDbAccount)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingDbAccount{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingDbAccount 数据库账号
type NgingDbAccount struct {
	base    factory.Base
	objects []*NgingDbAccount

	Id       uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Title    string `db:"title" bson:"title" comment:"标题" json:"title" xml:"title"`
	Uid      uint   `db:"uid" bson:"uid" comment:"UID" json:"uid" xml:"uid"`
	Engine   string `db:"engine" bson:"engine" comment:"数据库引擎" json:"engine" xml:"engine"`
	Host     string `db:"host" bson:"host" comment:"服务器地址" json:"host" xml:"host"`
	User     string `db:"user" bson:"user" comment:"用户名" json:"user" xml:"user"`
	Password string `db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Name     string `db:"name" bson:"name" comment:"数据库名称" json:"name" xml:"name"`
	Options  string `db:"options" bson:"options" comment:"其它选项(JSON)" json:"options" xml:"options"`
	Created  uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated  uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingDbAccount) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingDbAccount) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingDbAccount) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingDbAccount) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingDbAccount) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingDbAccount) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingDbAccount) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingDbAccount) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingDbAccount) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingDbAccount) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingDbAccount) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingDbAccount) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingDbAccount) Objects() []*NgingDbAccount {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingDbAccount) XObjects() Slice_NgingDbAccount {
	return Slice_NgingDbAccount(a.Objects())
}

func (a *NgingDbAccount) NewObjects() factory.Ranger {
	return &Slice_NgingDbAccount{}
}

func (a *NgingDbAccount) InitObjects() *[]*NgingDbAccount {
	a.objects = []*NgingDbAccount{}
	return &a.objects
}

func (a *NgingDbAccount) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingDbAccount) Short_() string {
	return "nging_db_account"
}

func (a *NgingDbAccount) Struct_() string {
	return "NgingDbAccount"
}

func (a *NgingDbAccount) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingDbAccount) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingDbAccount) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingDbAccount) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingDbAccount:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbAccount(*v))
		case []*NgingDbAccount:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbAccount(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingDbAccount) GroupBy(keyField string, inputRows ...[]*NgingDbAccount) map[string][]*NgingDbAccount {
	var rows Slice_NgingDbAccount
	if len(inputRows) > 0 {
		rows = Slice_NgingDbAccount(inputRows[0])
	} else {
		rows = Slice_NgingDbAccount(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingDbAccount) KeyBy(keyField string, inputRows ...[]*NgingDbAccount) map[string]*NgingDbAccount {
	var rows Slice_NgingDbAccount
	if len(inputRows) > 0 {
		rows = Slice_NgingDbAccount(inputRows[0])
	} else {
		rows = Slice_NgingDbAccount(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingDbAccount) AsKV(keyField string, valueField string, inputRows ...[]*NgingDbAccount) param.Store {
	var rows Slice_NgingDbAccount
	if len(inputRows) > 0 {
		rows = Slice_NgingDbAccount(inputRows[0])
	} else {
		rows = Slice_NgingDbAccount(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingDbAccount) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingDbAccount:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbAccount(*v))
		case []*NgingDbAccount:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbAccount(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingDbAccount) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Engine) == 0 {
		a.Engine = "mysql"
	}
	if len(a.Host) == 0 {
		a.Host = "localhost:3306"
	}
	if len(a.User) == 0 {
		a.User = "root"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingDbAccount) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Engine) == 0 {
		a.Engine = "mysql"
	}
	if len(a.Host) == 0 {
		a.Host = "localhost:3306"
	}
	if len(a.User) == 0 {
		a.User = "root"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingDbAccount) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingDbAccount) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["engine"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["engine"] = "mysql"
		}
	}
	if val, ok := kvset["host"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["host"] = "localhost:3306"
		}
	}
	if val, ok := kvset["user"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["user"] = "root"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingDbAccount) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Engine) == 0 {
			a.Engine = "mysql"
		}
		if len(a.Host) == 0 {
			a.Host = "localhost:3306"
		}
		if len(a.User) == 0 {
			a.User = "root"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Engine) == 0 {
			a.Engine = "mysql"
		}
		if len(a.Host) == 0 {
			a.Host = "localhost:3306"
		}
		if len(a.User) == 0 {
			a.User = "root"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingDbAccount) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingDbAccount) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingDbAccount) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingDbAccount) Reset() *NgingDbAccount {
	a.Id = 0
	a.Title = ``
	a.Uid = 0
	a.Engine = ``
	a.Host = ``
	a.User = ``
	a.Password = ``
	a.Name = ``
	a.Options = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingDbAccount) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Title"] = a.Title
	r["Uid"] = a.Uid
	r["Engine"] = a.Engine
	r["Host"] = a.Host
	r["User"] = a.User
	r["Password"] = a.Password
	r["Name"] = a.Name
	r["Options"] = a.Options
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	return r
}

func (a *NgingDbAccount) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "title":
			a.Title = param.AsString(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "engine":
			a.Engine = param.AsString(value)
		case "host":
			a.Host = param.AsString(value)
		case "user":
			a.User = param.AsString(value)
		case "password":
			a.Password = param.AsString(value)
		case "name":
			a.Name = param.AsString(value)
		case "options":
			a.Options = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingDbAccount) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Title":
			a.Title = param.AsString(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "Engine":
			a.Engine = param.AsString(vv)
		case "Host":
			a.Host = param.AsString(vv)
		case "User":
			a.User = param.AsString(vv)
		case "Password":
			a.Password = param.AsString(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Options":
			a.Options = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingDbAccount) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["title"] = a.Title
	r["uid"] = a.Uid
	r["engine"] = a.Engine
	r["host"] = a.Host
	r["user"] = a.User
	r["password"] = a.Password
	r["name"] = a.Name
	r["options"] = a.Options
	r["created"] = a.Created
	r["updated"] = a.Updated
	return r
}

func (a *NgingDbAccount) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingDbAccount) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
