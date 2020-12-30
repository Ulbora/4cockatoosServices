package cocka2db

import (
	"strconv"
)

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

//AddNoteItem AddNoteItem
func (c *C2DB) AddNoteItem(ni *NoteItem) (bool, int64) {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, ni.Text, ni.NoteID)
	suc, id := c.DB.Insert(insertNoteItem, a...)
	c.Log.Debug("suc in add NoteItem", suc)
	c.Log.Debug("id in add NoteItem", id)
	return suc, id
}

//UpdateNoteItem UpdateNoteItem
func (c *C2DB) UpdateNoteItem(ni *NoteItem) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, ni.Text, ni.ID)
	suc := c.DB.Update(updateNoteItem, a...)
	return suc
}

//GetNoteItemList GetNoteItemList
func (c *C2DB) GetNoteItemList(noteID int64) *[]NoteItem {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var rtn = []NoteItem{}
	var a []interface{}
	a = append(a, noteID)
	rows := c.DB.GetList(getNoteItemList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := c.parseNoteItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteNoteItem DeleteNoteItem
func (c *C2DB) DeleteNoteItem(id int64) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return c.DB.Delete(deleteNoteItem, a...)
}

func (c *C2DB) parseNoteItemRow(foundRow *[]string) *NoteItem {
	c.Log.Debug("foundRow in get NoteItem", *foundRow)
	var rtn NoteItem
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		c.Log.Debug("id err in get NoteItem", err)
		if err == nil {
			nid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
			c.Log.Debug("id err in get NoteItem", err)
			if err == nil {
				rtn.ID = id
				rtn.Text = (*foundRow)[1]
				rtn.NoteID = nid
			}
		}
	}
	return &rtn
}
