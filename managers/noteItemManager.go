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

//AddNoteItem AddNoteItem
func (m *C2Manager) AddNoteItem(ni *db.NoteItem) *ResponseID {
	m.Log.Debug("AddNoteItem: ", *ni)
	var rtn ResponseID
	suc, id := m.Db.AddNoteItem(ni)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
	}
	return &rtn
}

//UpdateNoteItem UpdateNoteItem
func (m *C2Manager) UpdateNoteItem(ni *db.NoteItem) *Response {
	m.Log.Debug("UpdateNoteItem: ", *ni)
	var rtn Response
	suc := m.Db.UpdateNoteItem(ni)
	rtn.Success = suc
	return &rtn
}

//DeleteNoteItem DeleteNoteItem
func (m *C2Manager) DeleteNoteItem(id int64) *Response {
	m.Log.Debug("DeleteNoteItem: ", id)
	var rtn Response
	res := m.Db.DeleteNoteItem(id)
	rtn.Success = res
	return &rtn
}
