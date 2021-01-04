package cocka2db

import (
	"strconv"
	"time"
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

//AddNote AddNote
func (c *C2DB) AddNote(n *Note) (bool, int64) {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, n.Title, n.Type, n.OwnerEmail, time.Now())
	suc, id := c.DB.Insert(insertNote, a...)
	c.Log.Debug("suc in add Note", suc)
	c.Log.Debug("id in add Note", id)
	return suc, id
}

//UpdateNote UpdateNote
func (c *C2DB) UpdateNote(n *Note) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, n.Title, n.Type, time.Now(), n.ID)
	suc := c.DB.Update(updateNote, a...)
	return suc
}

//GetNote GetNote
func (c *C2DB) GetNote(id int64) *Note {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	row := c.DB.Get(getNote, a...)
	rtn := c.parseNoteRow(&row.Row)
	return rtn
}

//GetUsersNotes GetUsersNotes
func (c *C2DB) GetUsersNotes(email string) *[]Note {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var rtn = []Note{}
	var a []interface{}
	a = append(a, email)
	rows := c.DB.GetList(getNoteList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := c.parseNoteRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteNote DeleteNote
func (c *C2DB) DeleteNote(id int64) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return c.DB.Delete(deleteNote, a...)
}

func (c *C2DB) parseNoteRow(foundRow *[]string) *Note {
	c.Log.Debug("foundRow in get Note", *foundRow)
	var rtn Note
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		c.Log.Debug("id err in get note", err)
		if err == nil {
			uTime, _ := time.Parse(timeFormat, (*foundRow)[4])
			rtn.ID = id
			rtn.Title = (*foundRow)[1]
			rtn.Type = (*foundRow)[2]
			rtn.OwnerEmail = (*foundRow)[3]
			rtn.LastUsed = uTime
		}
	}
	return &rtn
}
