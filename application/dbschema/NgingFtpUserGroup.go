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

type Slice_NgingFtpUserGroup []*NgingFtpUserGroup

func (s Slice_NgingFtpUserGroup) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFtpUserGroup) RangeRaw(fn func(m *NgingFtpUserGroup) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFtpUserGroup) GroupBy(keyField string) map[string][]*NgingFtpUserGroup {
	r := map[string][]*NgingFtpUserGroup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFtpUserGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFtpUserGroup) KeyBy(keyField string) map[string]*NgingFtpUserGroup {
	r := map[string]*NgingFtpUserGroup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFtpUserGroup) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFtpUserGroup) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFtpUserGroup) FromList(data interface{}) Slice_NgingFtpUserGroup {
	values, ok := data.([]*NgingFtpUserGroup)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFtpUserGroup{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingFtpUserGroup FTP用户组
type NgingFtpUserGroup struct {
	base    factory.Base
	objects []*NgingFtpUserGroup

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Name        string `db:"name" bson:"name" comment:"组名称" json:"name" xml:"name"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Banned      string `db:"banned" bson:"banned" comment:"是否禁止组内用户连接" json:"banned" xml:"banned"`
	Directory   string `db:"directory" bson:"directory" comment:"授权目录" json:"directory" xml:"directory"`
	IpWhitelist string `db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个)" json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist string `db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个)" json:"ip_blacklist" xml:"ip_blacklist"`
}

// - base function

func (a *NgingFtpUserGroup) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingFtpUserGroup) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFtpUserGroup) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFtpUserGroup) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFtpUserGroup) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFtpUserGroup) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFtpUserGroup) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFtpUserGroup) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFtpUserGroup) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingFtpUserGroup) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFtpUserGroup) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingFtpUserGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingFtpUserGroup) Objects() []*NgingFtpUserGroup {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFtpUserGroup) XObjects() Slice_NgingFtpUserGroup {
	return Slice_NgingFtpUserGroup(a.Objects())
}

func (a *NgingFtpUserGroup) NewObjects() factory.Ranger {
	return &Slice_NgingFtpUserGroup{}
}

func (a *NgingFtpUserGroup) InitObjects() *[]*NgingFtpUserGroup {
	a.objects = []*NgingFtpUserGroup{}
	return &a.objects
}

func (a *NgingFtpUserGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFtpUserGroup) Short_() string {
	return "nging_ftp_user_group"
}

func (a *NgingFtpUserGroup) Struct_() string {
	return "NgingFtpUserGroup"
}

func (a *NgingFtpUserGroup) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingFtpUserGroup) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFtpUserGroup) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingFtpUserGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFtpUserGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUserGroup(*v))
		case []*NgingFtpUserGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUserGroup(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFtpUserGroup) GroupBy(keyField string, inputRows ...[]*NgingFtpUserGroup) map[string][]*NgingFtpUserGroup {
	var rows Slice_NgingFtpUserGroup
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpUserGroup(inputRows[0])
	} else {
		rows = Slice_NgingFtpUserGroup(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFtpUserGroup) KeyBy(keyField string, inputRows ...[]*NgingFtpUserGroup) map[string]*NgingFtpUserGroup {
	var rows Slice_NgingFtpUserGroup
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpUserGroup(inputRows[0])
	} else {
		rows = Slice_NgingFtpUserGroup(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFtpUserGroup) AsKV(keyField string, valueField string, inputRows ...[]*NgingFtpUserGroup) param.Store {
	var rows Slice_NgingFtpUserGroup
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpUserGroup(inputRows[0])
	} else {
		rows = Slice_NgingFtpUserGroup(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFtpUserGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingFtpUserGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUserGroup(*v))
		case []*NgingFtpUserGroup:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpUserGroup(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFtpUserGroup) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Banned) == 0 {
		a.Banned = "N"
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

func (a *NgingFtpUserGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Banned) == 0 {
		a.Banned = "N"
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

func (a *NgingFtpUserGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFtpUserGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["banned"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["banned"] = "N"
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

func (a *NgingFtpUserGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Banned) == 0 {
			a.Banned = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Banned) == 0 {
			a.Banned = "N"
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

func (a *NgingFtpUserGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingFtpUserGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFtpUserGroup) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFtpUserGroup) Reset() *NgingFtpUserGroup {
	a.Id = 0
	a.Name = ``
	a.Created = 0
	a.Updated = 0
	a.Disabled = ``
	a.Banned = ``
	a.Directory = ``
	a.IpWhitelist = ``
	a.IpBlacklist = ``
	return a
}

func (a *NgingFtpUserGroup) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Name"] = a.Name
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["Disabled"] = a.Disabled
	r["Banned"] = a.Banned
	r["Directory"] = a.Directory
	r["IpWhitelist"] = a.IpWhitelist
	r["IpBlacklist"] = a.IpBlacklist
	return r
}

func (a *NgingFtpUserGroup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "banned":
			a.Banned = param.AsString(value)
		case "directory":
			a.Directory = param.AsString(value)
		case "ip_whitelist":
			a.IpWhitelist = param.AsString(value)
		case "ip_blacklist":
			a.IpBlacklist = param.AsString(value)
		}
	}
}

func (a *NgingFtpUserGroup) Set(key interface{}, value ...interface{}) {
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
		case "Name":
			a.Name = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Banned":
			a.Banned = param.AsString(vv)
		case "Directory":
			a.Directory = param.AsString(vv)
		case "IpWhitelist":
			a.IpWhitelist = param.AsString(vv)
		case "IpBlacklist":
			a.IpBlacklist = param.AsString(vv)
		}
	}
}

func (a *NgingFtpUserGroup) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["name"] = a.Name
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["disabled"] = a.Disabled
	r["banned"] = a.Banned
	r["directory"] = a.Directory
	r["ip_whitelist"] = a.IpWhitelist
	r["ip_blacklist"] = a.IpBlacklist
	return r
}

func (a *NgingFtpUserGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFtpUserGroup) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
