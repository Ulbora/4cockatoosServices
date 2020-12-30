package cocka2db

/*
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	dbi "github.com/Ulbora/dbinterface"
	"strconv"
)

//NotesDB NotesDB
type NotesDB interface {
	AddUser(u *User) bool
	UpdateUser(u *User) bool
	GetUser(email string) *User
	DeleteUser(email string) bool

	AddNote(n *Note) (bool, int64)
	UpdateNote(n *Note) bool
	GetNote(id int64) *Note
	GetUsersNotes(email string) *[]Note
	DeleteNote(id int64) bool

	AddCheckboxItem(ni *CheckboxNoteItem) (bool, int64)
	UpdateCheckboxItem(ni *CheckboxNoteItem) bool
	// //GetCheckboxItem(id int64) *CheckboxNoteItem
	GetCheckboxItemList(noteID int64) *[]CheckboxNoteItem
	DeleteCheckboxItem(id int64) bool

	AddNoteItem(ni *NoteItem) (bool, int64)
	UpdateNoteItem(ni *NoteItem) bool
	// //GetNoteItem(id int64) *NoteItem
	GetNoteItemList(noteID int64) *[]NoteItem
	DeleteNoteItem(id int64) bool

	AddNoteUser(nu *NoteUsers) bool
	DeleteNoteUser(nu *NoteUsers) bool
}

//C2DB C2DB
type C2DB struct {
	DB  dbi.Database
	Log *lg.Logger
}

//GetNew GetNew
func (c *C2DB) GetNew() NotesDB {
	return c
}

func (c *C2DB) testConnection() bool {
	c.Log.Debug("in testConnection")
	var rtn = false
	var a []interface{}
	c.Log.Debug("c.DB: ", fmt.Sprintln(c.DB))
	rowPtr := c.DB.Test(noteTest, a...)
	c.Log.Debug("rowPtr", *rowPtr)
	c.Log.Debug("after testConnection test", *rowPtr)
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		//log.Print("Records found during test ")
		//log.Println("Records found during test :", int64Val)
		if err != nil {
			c.Log.Error(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

// go mod init github.com/Ulbora/cocka2notesServices
