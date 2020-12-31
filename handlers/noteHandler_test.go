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
	"github.com/gorilla/mux"
)

func TestC2Handler_AddNote(t *testing.T) {
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

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 3
	cdb.MockAddNoteUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteMedia(t *testing.T) {
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

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 3
	cdb.MockAddNoteUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteAuth(t *testing.T) {
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

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 3
	cdb.MockAddNoteUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteReq(t *testing.T) {
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

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 3
	cdb.MockAddNoteUserSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_AddNoteFail(t *testing.T) {
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

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 3
	//cdb.MockAddNoteUserSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.AddNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNote(t *testing.T) {

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

	cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteMedia(t *testing.T) {

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

	cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteReq(t *testing.T) {

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

	cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteAuth(t *testing.T) {

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

	cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_UpdateNoteFail(t *testing.T) {

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

	//cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.UpdateNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

func TestC2Handler_GetNote(t *testing.T) {

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
	cdb.MockNote = &fn

	var ni1 db.CheckboxNoteItem
	ni1.ID = 1
	ni1.Checked = true
	ni1.NoteID = 4
	ni1.Text = "test1"

	var ni2 db.CheckboxNoteItem
	ni2.ID = 1
	ni2.Checked = true
	ni2.NoteID = 4
	ni2.Text = "test1"

	var nlst []db.CheckboxNoteItem
	nlst = append(nlst, ni1)
	nlst = append(nlst, ni2)
	cdb.MockCheckboxNoteItemList = &nlst

	//cdb.MockUpdateNoteSuc = true

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

	h.GetNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_GetNoteReq(t *testing.T) {

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
	cdb.MockNote = &fn

	var ni1 db.CheckboxNoteItem
	ni1.ID = 1
	ni1.Checked = true
	ni1.NoteID = 4
	ni1.Text = "test1"

	var ni2 db.CheckboxNoteItem
	ni2.ID = 1
	ni2.Checked = true
	ni2.NoteID = 4
	ni2.Text = "test1"

	var nlst []db.CheckboxNoteItem
	nlst = append(nlst, ni1)
	nlst = append(nlst, ni2)
	cdb.MockCheckboxNoteItemList = &nlst

	//cdb.MockUpdateNoteSuc = true

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

	h.GetNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_GetNoteReq2(t *testing.T) {

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
	cdb.MockNote = &fn

	var ni1 db.CheckboxNoteItem
	ni1.ID = 1
	ni1.Checked = true
	ni1.NoteID = 4
	ni1.Text = "test1"

	var ni2 db.CheckboxNoteItem
	ni2.ID = 1
	ni2.Checked = true
	ni2.NoteID = 4
	ni2.Text = "test1"

	var nlst []db.CheckboxNoteItem
	nlst = append(nlst, ni1)
	nlst = append(nlst, ni2)
	cdb.MockCheckboxNoteItemList = &nlst

	//cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "2s",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_GetNoteAuth(t *testing.T) {

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
	cdb.MockNote = &fn

	var ni1 db.CheckboxNoteItem
	ni1.ID = 1
	ni1.Checked = true
	ni1.NoteID = 4
	ni1.Text = "test1"

	var ni2 db.CheckboxNoteItem
	ni2.ID = 1
	ni2.Checked = true
	ni2.NoteID = 4
	ni2.Text = "test1"

	var nlst []db.CheckboxNoteItem
	nlst = append(nlst, ni1)
	nlst = append(nlst, ni2)
	cdb.MockCheckboxNoteItemList = &nlst

	//cdb.MockUpdateNoteSuc = true

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

	h.GetNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_GetUsersNotes(t *testing.T) {

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
	cdb.MockNote = &fn

	var nlst []db.Note
	nlst = append(nlst, fn)

	cdb.MockNoteList = &nlst

	//cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetUsersNotes(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_GetUsersNotesReq(t *testing.T) {

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
	cdb.MockNote = &fn

	var nlst []db.Note
	nlst = append(nlst, fn)

	cdb.MockNoteList = &nlst

	//cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		//"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetUsersNotes(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_GetUsersNotesAuth(t *testing.T) {

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
	cdb.MockNote = &fn

	var nlst []db.Note
	nlst = append(nlst, fn)

	cdb.MockNoteList = &nlst

	//cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetUsersNotes(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_DeleteNote(t *testing.T) {

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

	cdb.MockDeleteNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id":    "2",
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_DeleteNoteReq(t *testing.T) {

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

	cdb.MockDeleteNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "2",
		//"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_DeleteNoteReq2(t *testing.T) {

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

	cdb.MockDeleteNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id":    "2s",
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_DeleteNoteAuth(t *testing.T) {

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

	cdb.MockDeleteNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id":    "2",
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_DeleteNoteFail(t *testing.T) {

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

	//cdb.MockDeleteNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id":    "2",
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.DeleteNote(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}
