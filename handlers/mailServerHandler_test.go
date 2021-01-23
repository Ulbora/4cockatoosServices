package handlers

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	m "github.com/Ulbora/cocka2notesServices/managers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestC2Handler_GetMailServer(t *testing.T) {

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

	var fn db.MailServer
	fn.ID = 4
	fn.SenderEmail = "test@test.com"
	cdb.MockMailServer = &fn

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetMailServer(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_GetMailServerAuth(t *testing.T) {

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

	var fn db.MailServer
	fn.ID = 4
	fn.SenderEmail = "test@test.com"
	cdb.MockMailServer = &fn

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)

	r.Header.Set("apiKey", "123456q")
	w := httptest.NewRecorder()

	h.GetMailServer(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}
