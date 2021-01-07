package managers

/*
 Six910 is a shopping cart and E-commerce system.

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
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

//AddUserToNote AddUserToNote
func (m *C2Manager) AddUserToNote(nu *NoteUsers) *Response {
	m.Log.Debug("UpdateNoteUser: ", *nu)
	var rtn Response
	fo := m.Db.GetUser(nu.OwnerEmail)
	fu := m.Db.GetUser(nu.UserEmail)
	n := m.Db.GetNote(nu.NoteID)
	if fo.Email != "" && fu.Email != "" && n.ID != 0 {
		var ntu db.NoteUsers
		ntu.NoteID = n.ID
		ntu.UserEmail = fu.Email
		m.Log.Debug("Updated Note User: ", *nu)
		suc := m.Db.AddNoteUser(&ntu)
		rtn.Success = suc
	}
	return &rtn
}

//GetNoteUserList GetNoteUserList
func (m *C2Manager) GetNoteUserList(noteID int64, ownerEmail string) *[]string {
	m.Log.Debug("GetNoteUserList: ", ownerEmail)
	return m.Db.GetNoteUserList(noteID, ownerEmail)
}
