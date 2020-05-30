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

type Slice_NgingDbSyncLog []*NgingDbSyncLog

func (s Slice_NgingDbSyncLog) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingDbSyncLog) RangeRaw(fn func(m *NgingDbSyncLog) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingDbSyncLog) GroupBy(keyField string) map[string][]*NgingDbSyncLog {
	r := map[string][]*NgingDbSyncLog{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingDbSyncLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingDbSyncLog) KeyBy(keyField string) map[string]*NgingDbSyncLog {
	r := map[string]*NgingDbSyncLog{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingDbSyncLog) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingDbSyncLog) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingDbSyncLog) FromList(data interface{}) Slice_NgingDbSyncLog {
	values, ok := data.([]*NgingDbSyncLog)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingDbSyncLog{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingDbSyncLog 数据表同步日志
type NgingDbSyncLog struct {
	base    factory.Base
	objects []*NgingDbSyncLog

	Id             uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	SyncId         uint   `db:"sync_id" bson:"sync_id" comment:"同步方案ID" json:"sync_id" xml:"sync_id"`
	Created        uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Result         string `db:"result" bson:"result" comment:"结果报表" json:"result" xml:"result"`
	ChangeTables   string `db:"change_tables" bson:"change_tables" comment:"被更改的表" json:"change_tables" xml:"change_tables"`
	ChangeTableNum uint   `db:"change_table_num" bson:"change_table_num" comment:"被更改的表的数量" json:"change_table_num" xml:"change_table_num"`
	Elapsed        uint64 `db:"elapsed" bson:"elapsed" comment:"总共耗时" json:"elapsed" xml:"elapsed"`
	Failed         uint   `db:"failed" bson:"failed" comment:"失败次数" json:"failed" xml:"failed"`
}

// - base function

func (a *NgingDbSyncLog) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingDbSyncLog) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingDbSyncLog) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingDbSyncLog) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingDbSyncLog) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingDbSyncLog) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingDbSyncLog) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingDbSyncLog) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingDbSyncLog) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingDbSyncLog) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingDbSyncLog) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingDbSyncLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingDbSyncLog) Objects() []*NgingDbSyncLog {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingDbSyncLog) XObjects() Slice_NgingDbSyncLog {
	return Slice_NgingDbSyncLog(a.Objects())
}

func (a *NgingDbSyncLog) NewObjects() factory.Ranger {
	return &Slice_NgingDbSyncLog{}
}

func (a *NgingDbSyncLog) InitObjects() *[]*NgingDbSyncLog {
	a.objects = []*NgingDbSyncLog{}
	return &a.objects
}

func (a *NgingDbSyncLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingDbSyncLog) Short_() string {
	return "nging_db_sync_log"
}

func (a *NgingDbSyncLog) Struct_() string {
	return "NgingDbSyncLog"
}

func (a *NgingDbSyncLog) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingDbSyncLog) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingDbSyncLog) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingDbSyncLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingDbSyncLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbSyncLog(*v))
		case []*NgingDbSyncLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbSyncLog(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingDbSyncLog) GroupBy(keyField string, inputRows ...[]*NgingDbSyncLog) map[string][]*NgingDbSyncLog {
	var rows Slice_NgingDbSyncLog
	if len(inputRows) > 0 {
		rows = Slice_NgingDbSyncLog(inputRows[0])
	} else {
		rows = Slice_NgingDbSyncLog(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingDbSyncLog) KeyBy(keyField string, inputRows ...[]*NgingDbSyncLog) map[string]*NgingDbSyncLog {
	var rows Slice_NgingDbSyncLog
	if len(inputRows) > 0 {
		rows = Slice_NgingDbSyncLog(inputRows[0])
	} else {
		rows = Slice_NgingDbSyncLog(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingDbSyncLog) AsKV(keyField string, valueField string, inputRows ...[]*NgingDbSyncLog) param.Store {
	var rows Slice_NgingDbSyncLog
	if len(inputRows) > 0 {
		rows = Slice_NgingDbSyncLog(inputRows[0])
	} else {
		rows = Slice_NgingDbSyncLog(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingDbSyncLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingDbSyncLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbSyncLog(*v))
		case []*NgingDbSyncLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingDbSyncLog(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingDbSyncLog) Add() (pk interface{}, err error) {
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
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingDbSyncLog) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingDbSyncLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingDbSyncLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

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

func (a *NgingDbSyncLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
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
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
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

func (a *NgingDbSyncLog) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingDbSyncLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingDbSyncLog) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingDbSyncLog) Reset() *NgingDbSyncLog {
	a.Id = 0
	a.SyncId = 0
	a.Created = 0
	a.Result = ``
	a.ChangeTables = ``
	a.ChangeTableNum = 0
	a.Elapsed = 0
	a.Failed = 0
	return a
}

func (a *NgingDbSyncLog) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["SyncId"] = a.SyncId
	r["Created"] = a.Created
	r["Result"] = a.Result
	r["ChangeTables"] = a.ChangeTables
	r["ChangeTableNum"] = a.ChangeTableNum
	r["Elapsed"] = a.Elapsed
	r["Failed"] = a.Failed
	return r
}

func (a *NgingDbSyncLog) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "sync_id":
			a.SyncId = param.AsUint(value)
		case "created":
			a.Created = param.AsUint(value)
		case "result":
			a.Result = param.AsString(value)
		case "change_tables":
			a.ChangeTables = param.AsString(value)
		case "change_table_num":
			a.ChangeTableNum = param.AsUint(value)
		case "elapsed":
			a.Elapsed = param.AsUint64(value)
		case "failed":
			a.Failed = param.AsUint(value)
		}
	}
}

func (a *NgingDbSyncLog) Set(key interface{}, value ...interface{}) {
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
			a.Id = param.AsUint64(vv)
		case "SyncId":
			a.SyncId = param.AsUint(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Result":
			a.Result = param.AsString(vv)
		case "ChangeTables":
			a.ChangeTables = param.AsString(vv)
		case "ChangeTableNum":
			a.ChangeTableNum = param.AsUint(vv)
		case "Elapsed":
			a.Elapsed = param.AsUint64(vv)
		case "Failed":
			a.Failed = param.AsUint(vv)
		}
	}
}

func (a *NgingDbSyncLog) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["sync_id"] = a.SyncId
	r["created"] = a.Created
	r["result"] = a.Result
	r["change_tables"] = a.ChangeTables
	r["change_table_num"] = a.ChangeTableNum
	r["elapsed"] = a.Elapsed
	r["failed"] = a.Failed
	return r
}

func (a *NgingDbSyncLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingDbSyncLog) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
