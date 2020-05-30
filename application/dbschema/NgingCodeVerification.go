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

type Slice_NgingCodeVerification []*NgingCodeVerification

func (s Slice_NgingCodeVerification) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCodeVerification) RangeRaw(fn func(m *NgingCodeVerification) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCodeVerification) GroupBy(keyField string) map[string][]*NgingCodeVerification {
	r := map[string][]*NgingCodeVerification{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCodeVerification{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCodeVerification) KeyBy(keyField string) map[string]*NgingCodeVerification {
	r := map[string]*NgingCodeVerification{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCodeVerification) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCodeVerification) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCodeVerification) FromList(data interface{}) Slice_NgingCodeVerification {
	values, ok := data.([]*NgingCodeVerification)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCodeVerification{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

// NgingCodeVerification 验证码
type NgingCodeVerification struct {
	base    factory.Base
	objects []*NgingCodeVerification

	Id         uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Code       string `db:"code" bson:"code" comment:"验证码" json:"code" xml:"code"`
	Created    uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	OwnerId    uint64 `db:"owner_id" bson:"owner_id" comment:"所有者ID" json:"owner_id" xml:"owner_id"`
	OwnerType  string `db:"owner_type" bson:"owner_type" comment:"所有者类型" json:"owner_type" xml:"owner_type"`
	Used       uint   `db:"used" bson:"used" comment:"使用时间" json:"used" xml:"used"`
	Purpose    string `db:"purpose" bson:"purpose" comment:"目的" json:"purpose" xml:"purpose"`
	Start      uint   `db:"start" bson:"start" comment:"有效时间" json:"start" xml:"start"`
	End        uint   `db:"end" bson:"end" comment:"失效时间" json:"end" xml:"end"`
	Disabled   string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	SendMethod string `db:"send_method" bson:"send_method" comment:"发送方式(mobile-手机;email-邮箱)" json:"send_method" xml:"send_method"`
	SendTo     string `db:"send_to" bson:"send_to" comment:"发送目标" json:"send_to" xml:"send_to"`
}

// - base function

func (a *NgingCodeVerification) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingCodeVerification) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCodeVerification) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCodeVerification) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCodeVerification) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCodeVerification) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCodeVerification) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCodeVerification) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCodeVerification) Namer() func(string) string {
	return a.base.Namer()
}

func (a *NgingCodeVerification) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCodeVerification) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingCodeVerification) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingCodeVerification) Objects() []*NgingCodeVerification {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCodeVerification) XObjects() Slice_NgingCodeVerification {
	return Slice_NgingCodeVerification(a.Objects())
}

func (a *NgingCodeVerification) NewObjects() factory.Ranger {
	return &Slice_NgingCodeVerification{}
}

func (a *NgingCodeVerification) InitObjects() *[]*NgingCodeVerification {
	a.objects = []*NgingCodeVerification{}
	return &a.objects
}

func (a *NgingCodeVerification) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCodeVerification) Short_() string {
	return "nging_code_verification"
}

func (a *NgingCodeVerification) Struct_() string {
	return "NgingCodeVerification"
}

func (a *NgingCodeVerification) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCodeVerification) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCodeVerification) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingCodeVerification) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCodeVerification:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCodeVerification(*v))
		case []*NgingCodeVerification:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCodeVerification(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCodeVerification) GroupBy(keyField string, inputRows ...[]*NgingCodeVerification) map[string][]*NgingCodeVerification {
	var rows Slice_NgingCodeVerification
	if len(inputRows) > 0 {
		rows = Slice_NgingCodeVerification(inputRows[0])
	} else {
		rows = Slice_NgingCodeVerification(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCodeVerification) KeyBy(keyField string, inputRows ...[]*NgingCodeVerification) map[string]*NgingCodeVerification {
	var rows Slice_NgingCodeVerification
	if len(inputRows) > 0 {
		rows = Slice_NgingCodeVerification(inputRows[0])
	} else {
		rows = Slice_NgingCodeVerification(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCodeVerification) AsKV(keyField string, valueField string, inputRows ...[]*NgingCodeVerification) param.Store {
	var rows Slice_NgingCodeVerification
	if len(inputRows) > 0 {
		rows = Slice_NgingCodeVerification(inputRows[0])
	} else {
		rows = Slice_NgingCodeVerification(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCodeVerification) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCodeVerification:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCodeVerification(*v))
		case []*NgingCodeVerification:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCodeVerification(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCodeVerification) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.OwnerType) == 0 {
		a.OwnerType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.SendMethod) == 0 {
		a.SendMethod = "mobile"
	}
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

func (a *NgingCodeVerification) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.OwnerType) == 0 {
		a.OwnerType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.SendMethod) == 0 {
		a.SendMethod = "mobile"
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

func (a *NgingCodeVerification) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCodeVerification) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["owner_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["owner_type"] = "user"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["send_method"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["send_method"] = "mobile"
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

func (a *NgingCodeVerification) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.OwnerType) == 0 {
			a.OwnerType = "user"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.SendMethod) == 0 {
			a.SendMethod = "mobile"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.OwnerType) == 0 {
			a.OwnerType = "user"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.SendMethod) == 0 {
			a.SendMethod = "mobile"
		}
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

func (a *NgingCodeVerification) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingCodeVerification) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCodeVerification) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCodeVerification) Reset() *NgingCodeVerification {
	a.Id = 0
	a.Code = ``
	a.Created = 0
	a.OwnerId = 0
	a.OwnerType = ``
	a.Used = 0
	a.Purpose = ``
	a.Start = 0
	a.End = 0
	a.Disabled = ``
	a.SendMethod = ``
	a.SendTo = ``
	return a
}

func (a *NgingCodeVerification) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Code"] = a.Code
	r["Created"] = a.Created
	r["OwnerId"] = a.OwnerId
	r["OwnerType"] = a.OwnerType
	r["Used"] = a.Used
	r["Purpose"] = a.Purpose
	r["Start"] = a.Start
	r["End"] = a.End
	r["Disabled"] = a.Disabled
	r["SendMethod"] = a.SendMethod
	r["SendTo"] = a.SendTo
	return r
}

func (a *NgingCodeVerification) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "code":
			a.Code = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "owner_id":
			a.OwnerId = param.AsUint64(value)
		case "owner_type":
			a.OwnerType = param.AsString(value)
		case "used":
			a.Used = param.AsUint(value)
		case "purpose":
			a.Purpose = param.AsString(value)
		case "start":
			a.Start = param.AsUint(value)
		case "end":
			a.End = param.AsUint(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "send_method":
			a.SendMethod = param.AsString(value)
		case "send_to":
			a.SendTo = param.AsString(value)
		}
	}
}

func (a *NgingCodeVerification) Set(key interface{}, value ...interface{}) {
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
		case "Code":
			a.Code = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "OwnerId":
			a.OwnerId = param.AsUint64(vv)
		case "OwnerType":
			a.OwnerType = param.AsString(vv)
		case "Used":
			a.Used = param.AsUint(vv)
		case "Purpose":
			a.Purpose = param.AsString(vv)
		case "Start":
			a.Start = param.AsUint(vv)
		case "End":
			a.End = param.AsUint(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "SendMethod":
			a.SendMethod = param.AsString(vv)
		case "SendTo":
			a.SendTo = param.AsString(vv)
		}
	}
}

func (a *NgingCodeVerification) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["code"] = a.Code
	r["created"] = a.Created
	r["owner_id"] = a.OwnerId
	r["owner_type"] = a.OwnerType
	r["used"] = a.Used
	r["purpose"] = a.Purpose
	r["start"] = a.Start
	r["end"] = a.End
	r["disabled"] = a.Disabled
	r["send_method"] = a.SendMethod
	r["send_to"] = a.SendTo
	return r
}

func (a *NgingCodeVerification) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCodeVerification) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
