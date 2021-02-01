package managers

import (
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

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

//AddNote AddNote
func (m *C2Manager) AddNote(n *db.Note) *ResponseID {
	m.Log.Debug("AddNote: ", *n)
	var rtn ResponseID
	suc, id := m.Db.AddNote(n)
	if suc && id != 0 {
		var nu db.NoteUsers
		nu.UserEmail = n.OwnerEmail
		nu.NoteID = id
		m.Log.Debug("Add Note User: ", nu)
		nusuc := m.Db.AddNoteUser(&nu)
		rtn.Success = nusuc
		rtn.ID = id
	}
	return &rtn
}

//UpdateNote UpdateNote
func (m *C2Manager) UpdateNote(n *db.Note) *Response {
	m.Log.Debug("UpdateNote: ", *n)
	var rtn Response
	en := m.Db.GetNote(n.ID)
	if en.ID != 0 {
		en.Title = n.Title
		m.Log.Debug("Updated Note: ", *en)
		suc := m.Db.UpdateNote(en)
		rtn.Success = suc
	}
	return &rtn
}

//GetUsersNotes GetUsersNotes
func (m *C2Manager) GetUsersNotes(email string) *[]*Note { //*[]db.Note
	var rtn []*Note
	m.Log.Debug("GetUsersNotes: ", email)
	nlist := m.Db.GetUsersNotes(email)
	for _, n := range *nlist {
		var nn Note
		nn.ID = n.ID
		nn.OwnerEmail = n.OwnerEmail
		nn.Title = n.Title
		nn.Type = n.Type
		nn.LastUsed = n.LastUsed
		if n.Type == noteTypeCheckbox {
			ni := m.Db.GetCheckboxItemList(n.ID)
			nn.NoteItems = ni
		} else if n.Type == notetypeNote {
			ni := m.Db.GetNoteItemList(n.ID)
			nn.NoteItems = ni
		}
		rtn = append(rtn, &nn)
	}
	return &rtn
}

//GetNote GetNote
func (m *C2Manager) GetNote(id int64) *Note {
	var rtn Note
	n := m.Db.GetNote(id)
	rtn.ID = n.ID
	rtn.OwnerEmail = n.OwnerEmail
	rtn.Title = n.Title
	rtn.Type = n.Type

	if n.Type == noteTypeCheckbox {
		ni := m.Db.GetCheckboxItemList(id)
		rtn.NoteItems = ni
	} else if n.Type == notetypeNote {
		ni := m.Db.GetNoteItemList(id)
		rtn.NoteItems = ni
	}
	return &rtn
}

//DeleteNote DeleteNote
func (m *C2Manager) DeleteNote(id int64, ownerEmail string) *Response {
	m.Log.Debug("DeleteNote: ", id, ownerEmail)
	var rtn Response
	en := m.Db.GetNote(id)
	if en.OwnerEmail == ownerEmail {
		res := m.Db.DeleteNote(id)
		rtn.Success = res
	}
	return &rtn
}
