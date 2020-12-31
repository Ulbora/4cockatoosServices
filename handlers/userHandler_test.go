package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	m "github.com/Ulbora/cocka2notesServices/managers"
)

func TestC2Handler_AddUser(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	cdb.MockAddUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_AddUserMedia(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	cdb.MockAddUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUser(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_AddUserReq(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	cdb.MockAddUserSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_AddUserAuth(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456a"
	sh.Log = &l

	cdb.MockAddUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_AddUserFail(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	//cdb.MockAddUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestC2Handler_UpdateUser(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u
	cdb.MockUpdateUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}


func TestC2Handler_UpdateUserMedia(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u
	cdb.MockUpdateUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}


func TestC2Handler_UpdateUserReq(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u
	cdb.MockUpdateUserSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}


func TestC2Handler_UpdateUserAuth(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456a"
	sh.Log = &l

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u
	cdb.MockUpdateUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}


func TestC2Handler_UpdateUserFail(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u
	//cdb.MockUpdateUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test", "webEnabled": true}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}
