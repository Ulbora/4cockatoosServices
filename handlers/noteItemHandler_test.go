package handlers

import (
	"github.com/gorilla/mux"
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

func TestC2Handler_AddNoteItem(t *testing.T) {
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

	cdb.MockAddNoteItemSuc = true
	cdb.MockNoteItemID = 2

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"text":"test",  "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteItemReq(t *testing.T) {
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

	cdb.MockAddNoteItemSuc = true
	cdb.MockNoteItemID = 2

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"text":"test",  "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteItemMedia(t *testing.T) {
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

	cdb.MockAddNoteItemSuc = true
	cdb.MockNoteItemID = 2

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"text":"test",  "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteItemAuth(t *testing.T) {
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

	cdb.MockAddNoteItemSuc = true
	cdb.MockNoteItemID = 2

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"text":"test",  "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteItemFail(t *testing.T) {
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

	cdb.MockAddNoteItemSuc = true
	//cdb.MockNoteItemID = 2

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"text":"test",  "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteItem(t *testing.T) {

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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	cdb.MockNote = &fn

	cdb.MockUpdateNoteItemSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "text":"test", "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteItemReq(t *testing.T) {

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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	cdb.MockNote = &fn

	cdb.MockUpdateNoteItemSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "text":"test", "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteItemMedia(t *testing.T) {

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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	cdb.MockNote = &fn

	cdb.MockUpdateNoteItemSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "text":"test", "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteItemAuth(t *testing.T) {

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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	cdb.MockNote = &fn

	cdb.MockUpdateNoteItemSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "text":"test", "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteItemFail(t *testing.T) {

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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	cdb.MockNote = &fn

	//cdb.MockUpdateNoteItemSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "text":"test", "noteId": 5}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestC2Handler_DeleteNoteItem(t *testing.T) {
	
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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	fn.Type = "checkbox"
	fn.OwnerEmail = "test@test.com"
	cdb.MockNote = &fn

	cdb.MockDeleteNoteItemSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "2",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}



func TestC2Handler_DeleteNoteItemReq(t *testing.T) {
	
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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	fn.Type = "checkbox"
	fn.OwnerEmail = "test@test.com"
	cdb.MockNote = &fn

	cdb.MockDeleteNoteItemSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"id": "2",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}



func TestC2Handler_DeleteNoteItemReq2(t *testing.T) {
	
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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	fn.Type = "checkbox"
	fn.OwnerEmail = "test@test.com"
	cdb.MockNote = &fn

	cdb.MockDeleteNoteItemSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "2a",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}



func TestC2Handler_DeleteNoteItemAuth(t *testing.T) {
	
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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	fn.Type = "checkbox"
	fn.OwnerEmail = "test@test.com"
	cdb.MockNote = &fn

	cdb.MockDeleteNoteItemSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "2",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}



func TestC2Handler_DeleteNoteItemFail(t *testing.T) {
	
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

	var fn db.Note
	fn.ID = 4
	fn.Title = "test444"
	fn.Type = "checkbox"
	fn.OwnerEmail = "test@test.com"
	cdb.MockNote = &fn

	//cdb.MockDeleteNoteItemSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "2",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNoteItem(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}
