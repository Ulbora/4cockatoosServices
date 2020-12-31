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

func TestC2Handler_AddUserToNote(t *testing.T) {
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

	cdb.MockAddNoteUserSuc = true

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u

	var fn db.Note
	fn.ID = 4
	cdb.MockNote = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"ownerEmail":"test@test.com", "userEmail": "test2@test.com", "noteId": 4}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUserToNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}


func TestC2Handler_AddUserToNoteMedia(t *testing.T) {
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

	cdb.MockAddNoteUserSuc = true

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u

	var fn db.Note
	fn.ID = 4
	cdb.MockNote = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"ownerEmail":"test@test.com", "userEmail": "test2@test.com", "noteId": 4}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUserToNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}


func TestC2Handler_AddUserToNoteReq(t *testing.T) {
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

	cdb.MockAddNoteUserSuc = true

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u

	var fn db.Note
	fn.ID = 4
	cdb.MockNote = &fn

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"ownerEmail":"test@test.com", "userEmail": "test2@test.com", "noteId": 4}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUserToNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}



func TestC2Handler_AddUserToNoteAuth(t *testing.T) {
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

	cdb.MockAddNoteUserSuc = true

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u

	var fn db.Note
	fn.ID = 4
	cdb.MockNote = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"ownerEmail":"test@test.com", "userEmail": "test2@test.com", "noteId": 4}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUserToNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}


func TestC2Handler_AddUserToNoteFail(t *testing.T) {
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

	//cdb.MockAddNoteUserSuc = true

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u

	var fn db.Note
	fn.ID = 4
	cdb.MockNote = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"ownerEmail":"test@test.com", "userEmail": "test2@test.com", "noteId": 4}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddUserToNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}
