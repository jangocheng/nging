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

type Slice_NgingFrpGroup []*NgingFrpGroup

func (s Slice_NgingFrpGroup) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpGroup) RangeRaw(fn func(m *NgingFrpGroup) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFrpGroup) GroupBy(keyField string) map[string][]*NgingFrpGroup {
	r := map[string][]*NgingFrpGroup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFrpGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFrpGroup) KeyBy(keyField string) map[string]*NgingFrpGroup {
	r := map[string]*NgingFrpGroup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFrpGroup) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFrpGroup) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFrpGroup) FromList(data interface{}) Slice_NgingFrpGroup {
	values, ok := data.([]*NgingFrpGroup)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFrpGroup{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFrpGroup FRP服务组
type NgingFrpGroup struct {
	base    factory.Base
	objects []*NgingFrpGroup

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid         uint   `db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name        string `db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingFrpGroup) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFrpGroup) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFrpGroup) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFrpGroup) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFrpGroup) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFrpGroup) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFrpGroup) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFrpGroup) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFrpGroup) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingFrpGroup) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFrpGroup) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFrpGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFrpGroup) Objects() []*NgingFrpGroup {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFrpGroup) XObjects() Slice_NgingFrpGroup {
	return Slice_NgingFrpGroup(a.Objects())
}

func (a *NgingFrpGroup) NewObjects() factory.Ranger {
	return &Slice_NgingFrpGroup{}
}

func (a *NgingFrpGroup) InitObjects() *[]*NgingFrpGroup {
	a.objects = []*NgingFrpGroup{}
	return &a.objects
}

func (a *NgingFrpGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFrpGroup) Short_() string {
	return "nging_frp_group"
}

func (a *NgingFrpGroup) Struct_() string {
	return "NgingFrpGroup"
}

func (a *NgingFrpGroup) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFrpGroup) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFrpGroup) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingFrpGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFrpGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpGroup(*v))
		case []*NgingFrpGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpGroup(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpGroup) GroupBy(keyField string, inputRows ...[]*NgingFrpGroup) map[string][]*NgingFrpGroup {
	var rows Slice_NgingFrpGroup
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpGroup(inputRows[0])
	} else {
		rows = Slice_NgingFrpGroup(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFrpGroup) KeyBy(keyField string, inputRows ...[]*NgingFrpGroup) map[string]*NgingFrpGroup {
	var rows Slice_NgingFrpGroup
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpGroup(inputRows[0])
	} else {
		rows = Slice_NgingFrpGroup(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFrpGroup) AsKV(keyField string, valueField string, inputRows ...[]*NgingFrpGroup) param.Store {
	var rows Slice_NgingFrpGroup
	if len(inputRows) > 0 {
		rows = Slice_NgingFrpGroup(inputRows[0])
	} else {
		rows = Slice_NgingFrpGroup(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFrpGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFrpGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpGroup(*v))
		case []*NgingFrpGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFrpGroup(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFrpGroup) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
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

func (a *NgingFrpGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
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

func (a *NgingFrpGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFrpGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	kvset["updated"] = uint(time.Now().Unix())
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

func (a *NgingFrpGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
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

func (a *NgingFrpGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFrpGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFrpGroup) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFrpGroup) Reset() *NgingFrpGroup {
	a.Id = 0
	a.Uid = 0
	a.Name = ``
	a.Description = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingFrpGroup) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Uid"] = a.Uid
	r["Name"] = a.Name
	r["Description"] = a.Description
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	return r
}

func (a *NgingFrpGroup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingFrpGroup) Set(key interface{}, value ...interface{}) {
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
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingFrpGroup) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["uid"] = a.Uid
	r["name"] = a.Name
	r["description"] = a.Description
	r["created"] = a.Created
	r["updated"] = a.Updated
	return r
}

func (a *NgingFrpGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFrpGroup) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
