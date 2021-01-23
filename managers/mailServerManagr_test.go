package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

func TestC2Manager_GetMailServer(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n

	res := m.GetMailServer()
	fmt.Println("GetMailServer: ", *res)
	if res.SenderEmail != "test@test.com" {
		t.Fail()
	}
}
