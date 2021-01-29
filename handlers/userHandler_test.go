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
	ml "github.com/Ulbora/go-mail-sender"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

func TestC2Handler_GetUser(t *testing.T) {

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

	var fn db.User
	fn.Email = "test@test.com"
	fn.Password = "eerrt"

	cdb.MockUser = &fn

	//cdb.MockUpdateNoteSuc = true

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "title":"test", "type": "checkbox", "ownerEmail": "test@test.com"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"email": "test@tst.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_GetUserReq(t *testing.T) {

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
		//"email": "test@tst.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.GetUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_GetUserAuth(t *testing.T) {

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
		"email": "test@tst.com",
	}
	r = mux.SetURLVars(r, vars)

	r.Header.Set("apiKey", "123456a")
	w := httptest.NewRecorder()

	h.GetUser(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_Login(t *testing.T) {
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

	var fn db.User
	fn.Email = "test@test.com"
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	if err == nil {
		fn.Password = string(hashedPw)
	}

	cdb.MockUser = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.Login(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_LoginAuth(t *testing.T) {
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
	sh.APIKey = "123456q"
	sh.Log = &l

	var fn db.User
	fn.Email = "test@test.com"
	fn.Password = "eerrt"

	cdb.MockUser = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.Login(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_LoginMedia(t *testing.T) {
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

	var fn db.User
	fn.Email = "test@test.com"
	fn.Password = "eerrt"

	cdb.MockUser = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	//r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.Login(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_LoginReq(t *testing.T) {
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

	var fn db.User
	fn.Email = "test@test.com"
	fn.Password = "eerrt"

	cdb.MockUser = &fn

	h := sh.GetNew()
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.Login(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_LoginFail(t *testing.T) {
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

	var fn db.User
	fn.Email = "test@test.com"
	fn.Password = "eerrt"

	cdb.MockUser = &fn

	h := sh.GetNew()
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test@test.com", "password": "test"}`))
	//aJSON, _ := json.Marshal(robj)
	//fmt.Println("aJSON: ", aJSON)
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("storeName", "TestStore")
	r.Header.Set("localDomain", "test.domain")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apiKey", "123456")
	w := httptest.NewRecorder()

	h.Login(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_ResetPassword(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n
	cdb.MockUpdateUserSuc = true
	var mkusr db.User
	mkusr.Email = "test@test.com"
	cdb.MockUser = &mkusr

	var msender ml.MockSecureSender
	msender.MockSuccess = true
	c2m.MailSender = msender.GetNew()

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

	h.ResetPassword(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestC2Handler_ResetPasswordMedia(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n
	cdb.MockUpdateUserSuc = true
	var mkusr db.User
	mkusr.Email = "test@test.com"
	cdb.MockUser = &mkusr

	var msender ml.MockSecureSender
	msender.MockSuccess = true
	c2m.MailSender = msender.GetNew()

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

	h.ResetPassword(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 415 {
		t.Fail()
	}
}

func TestC2Handler_ResetPasswordReq(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n
	cdb.MockUpdateUserSuc = true
	var mkusr db.User
	mkusr.Email = "test@test.com"
	cdb.MockUser = &mkusr

	var msender ml.MockSecureSender
	msender.MockSuccess = true
	c2m.MailSender = msender.GetNew()

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

	h.ResetPassword(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestC2Handler_ResetPasswordAuth(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n
	cdb.MockUpdateUserSuc = true
	var mkusr db.User
	mkusr.Email = "test@test.com"
	cdb.MockUser = &mkusr

	var msender ml.MockSecureSender
	msender.MockSuccess = true
	c2m.MailSender = msender.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456q"
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

	h.ResetPassword(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 401 {
		t.Fail()
	}
}

func TestC2Handler_ResetPasswordFail(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m m.C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n
	cdb.MockUpdateUserSuc = true
	var mkusr db.User
	mkusr.Email = "test@test.com"
	cdb.MockUser = &mkusr

	var msender ml.MockSecureSender
	msender.MockSuccess = true
	c2m.MailSender = msender.GetNew()

	var sh C2Handler
	sh.Manager = m
	sh.APIKey = "123456"
	sh.Log = &l

	var u db.User
	u.Email = "test@test.com"
	cdb.MockUser = &u
	cdb.MockUpdateUserSuc = false

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

	h.ResetPassword(w, r)

	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}
