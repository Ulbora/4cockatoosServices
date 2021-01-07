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
	lg "github.com/Ulbora/Level_Logger"
	dbi "github.com/Ulbora/dbinterface"
)

//MockC2DB MockC2DB
type MockC2DB struct {
	DB  dbi.Database
	Log *lg.Logger

	MockAddUserSuc    bool
	MockUpdateUserSuc bool
	MockUser          *User
	MockDeleteUserSuc bool

	MockAddNoteSuc    bool
	MockAddNoteID     int64
	MockUpdateNoteSuc bool
	MockNote          *Note
	MockNoteList      *[]Note
	MockDeleteNoteSuc bool

	MockAddNoteItemSuc    bool
	MockNoteItemID        int64
	MockUpdateNoteItemSuc bool
	MockNoteItemList      *[]NoteItem
	MockDeleteNoteItemSuc bool

	MockAddCheckboxNoteItemSuc    bool
	MockCheckboxNoteItemID        int64
	MockUpdateCheckboxNoteItemSuc bool
	MockCheckboxNoteItemList      *[]CheckboxNoteItem
	MockDeleteCheckboxNoteItemSuc bool

	MockAddNoteUserSuc    bool
	MockDeleteNoteUserSuc bool
	MockNoteUserList      *[]string
}

//GetNew GetNew
func (c *MockC2DB) GetNew() NotesDB {
	return c
}

//AddUser AddUser
func (c *MockC2DB) AddUser(u *User) bool {
	return c.MockAddUserSuc
}

//UpdateUser UpdateUser
func (c *MockC2DB) UpdateUser(u *User) bool {
	return c.MockUpdateUserSuc
}

//GetUser GetUser
func (c *MockC2DB) GetUser(email string) *User {
	return c.MockUser
}

//DeleteUser DeleteUser
func (c *MockC2DB) DeleteUser(email string) bool {
	return c.MockDeleteUserSuc
}

//AddNote AddNote
func (c *MockC2DB) AddNote(n *Note) (bool, int64) {
	return c.MockAddNoteSuc, c.MockAddNoteID
}

//UpdateNote UpdateNote
func (c *MockC2DB) UpdateNote(n *Note) bool {
	return c.MockUpdateNoteSuc
}

//GetNote GetNote
func (c *MockC2DB) GetNote(id int64) *Note {
	return c.MockNote
}

//GetUsersNotes GetUsersNotes
func (c *MockC2DB) GetUsersNotes(email string) *[]Note {
	return c.MockNoteList
}

//DeleteNote DeleteNote
func (c *MockC2DB) DeleteNote(id int64) bool {
	return c.MockDeleteNoteSuc
}

//AddNoteItem AddNoteItem
func (c *MockC2DB) AddNoteItem(ni *NoteItem) (bool, int64) {
	return c.MockAddNoteItemSuc, c.MockNoteItemID
}

//UpdateNoteItem UpdateNoteItem
func (c *MockC2DB) UpdateNoteItem(ni *NoteItem) bool {
	return c.MockUpdateNoteItemSuc
}

//GetNoteItemList GetNoteItemList
func (c *MockC2DB) GetNoteItemList(noteID int64) *[]NoteItem {
	return c.MockNoteItemList
}

//DeleteNoteItem DeleteNoteItem
func (c *MockC2DB) DeleteNoteItem(id int64) bool {
	return c.MockDeleteNoteItemSuc
}

//AddCheckboxItem AddCheckboxItem
func (c *MockC2DB) AddCheckboxItem(ni *CheckboxNoteItem) (bool, int64) {
	return c.MockAddCheckboxNoteItemSuc, c.MockCheckboxNoteItemID
}

//UpdateCheckboxItem UpdateCheckboxItem
func (c *MockC2DB) UpdateCheckboxItem(ni *CheckboxNoteItem) bool {
	return c.MockUpdateCheckboxNoteItemSuc
}

//GetCheckboxItemList GetCheckboxItemList
func (c *MockC2DB) GetCheckboxItemList(noteID int64) *[]CheckboxNoteItem {
	return c.MockCheckboxNoteItemList
}

//DeleteCheckboxItem DeleteCheckboxItem
func (c *MockC2DB) DeleteCheckboxItem(id int64) bool {
	return c.MockDeleteCheckboxNoteItemSuc
}

//AddNoteUser AddNoteUser
func (c *MockC2DB) AddNoteUser(nu *NoteUsers) bool {
	return c.MockAddNoteUserSuc
}

//GetNoteUserList GetNoteUserList
func (c *MockC2DB) GetNoteUserList(noteID int64, ownerEmail string) *[]string {
	return c.MockNoteUserList
}

//DeleteNoteUser DeleteNoteUser
func (c *MockC2DB) DeleteNoteUser(nu *NoteUsers) bool {
	return c.MockDeleteNoteUserSuc
}
