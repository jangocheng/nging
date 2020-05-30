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

type Slice_NgingCollectorRule []*NgingCollectorRule

func (s Slice_NgingCollectorRule) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorRule) RangeRaw(fn func(m *NgingCollectorRule) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorRule) GroupBy(keyField string) map[string][]*NgingCollectorRule {
	r := map[string][]*NgingCollectorRule{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCollectorRule{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCollectorRule) KeyBy(keyField string) map[string]*NgingCollectorRule {
	r := map[string]*NgingCollectorRule{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCollectorRule) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCollectorRule) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCollectorRule) FromList(data interface{}) Slice_NgingCollectorRule {
	values, ok := data.([]*NgingCollectorRule)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCollectorRule{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingCollectorRule 页面中的元素采集规则
type NgingCollectorRule struct {
	base    factory.Base
	objects []*NgingCollectorRule

	Id      uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	PageId  uint   `db:"page_id" bson:"page_id" comment:"页面ID" json:"page_id" xml:"page_id"`
	Name    string `db:"name" bson:"name" comment:"保存匹配结果的变量名" json:"name" xml:"name"`
	Rule    string `db:"rule" bson:"rule" comment:"规则" json:"rule" xml:"rule"`
	Type    string `db:"type" bson:"type" comment:"数据类型" json:"type" xml:"type"`
	Filter  string `db:"filter" bson:"filter" comment:"过滤器" json:"filter" xml:"filter"`
	Created uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Sort    int    `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
}

// - base function

func (a *NgingCollectorRule) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingCollectorRule) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCollectorRule) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCollectorRule) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCollectorRule) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCollectorRule) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCollectorRule) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCollectorRule) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCollectorRule) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingCollectorRule) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCollectorRule) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingCollectorRule) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingCollectorRule) Objects() []*NgingCollectorRule {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCollectorRule) XObjects() Slice_NgingCollectorRule {
	return Slice_NgingCollectorRule(a.Objects())
}

func (a *NgingCollectorRule) NewObjects() factory.Ranger {
	return &Slice_NgingCollectorRule{}
}

func (a *NgingCollectorRule) InitObjects() *[]*NgingCollectorRule {
	a.objects = []*NgingCollectorRule{}
	return &a.objects
}

func (a *NgingCollectorRule) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCollectorRule) Short_() string {
	return "nging_collector_rule"
}

func (a *NgingCollectorRule) Struct_() string {
	return "NgingCollectorRule"
}

func (a *NgingCollectorRule) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCollectorRule) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCollectorRule) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingCollectorRule) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCollectorRule:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorRule(*v))
		case []*NgingCollectorRule:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorRule(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorRule) GroupBy(keyField string, inputRows ...[]*NgingCollectorRule) map[string][]*NgingCollectorRule {
	var rows Slice_NgingCollectorRule
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorRule(inputRows[0])
	} else {
		rows = Slice_NgingCollectorRule(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCollectorRule) KeyBy(keyField string, inputRows ...[]*NgingCollectorRule) map[string]*NgingCollectorRule {
	var rows Slice_NgingCollectorRule
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorRule(inputRows[0])
	} else {
		rows = Slice_NgingCollectorRule(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCollectorRule) AsKV(keyField string, valueField string, inputRows ...[]*NgingCollectorRule) param.Store {
	var rows Slice_NgingCollectorRule
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorRule(inputRows[0])
	} else {
		rows = Slice_NgingCollectorRule(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCollectorRule) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCollectorRule:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorRule(*v))
		case []*NgingCollectorRule:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorRule(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorRule) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Type) == 0 {
		a.Type = "string"
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

func (a *NgingCollectorRule) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.Type) == 0 {
		a.Type = "string"
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

func (a *NgingCollectorRule) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCollectorRule) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "string"
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

func (a *NgingCollectorRule) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.Type) == 0 {
			a.Type = "string"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Type) == 0 {
			a.Type = "string"
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

func (a *NgingCollectorRule) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingCollectorRule) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCollectorRule) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCollectorRule) Reset() *NgingCollectorRule {
	a.Id = 0
	a.PageId = 0
	a.Name = ``
	a.Rule = ``
	a.Type = ``
	a.Filter = ``
	a.Created = 0
	a.Sort = 0
	return a
}

func (a *NgingCollectorRule) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["PageId"] = a.PageId
	r["Name"] = a.Name
	r["Rule"] = a.Rule
	r["Type"] = a.Type
	r["Filter"] = a.Filter
	r["Created"] = a.Created
	r["Sort"] = a.Sort
	return r
}

func (a *NgingCollectorRule) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "page_id":
			a.PageId = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "rule":
			a.Rule = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "filter":
			a.Filter = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "sort":
			a.Sort = param.AsInt(value)
		}
	}
}

func (a *NgingCollectorRule) Set(key interface{}, value ...interface{}) {
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
		case "PageId":
			a.PageId = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Rule":
			a.Rule = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Filter":
			a.Filter = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Sort":
			a.Sort = param.AsInt(vv)
		}
	}
}

func (a *NgingCollectorRule) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["page_id"] = a.PageId
	r["name"] = a.Name
	r["rule"] = a.Rule
	r["type"] = a.Type
	r["filter"] = a.Filter
	r["created"] = a.Created
	r["sort"] = a.Sort
	return r
}

func (a *NgingCollectorRule) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCollectorRule) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
