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

//AddCheckboxItem AddCheckboxItem
func (c *C2DB) AddCheckboxItem(ni *CheckboxNoteItem) (bool, int64) {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, ni.Text, ni.Checked, ni.NoteID)
	suc, id := c.DB.Insert(insertCheckboxNoteItem, a...)
	c.Log.Debug("suc in add CheckboxNoteItem", suc)
	c.Log.Debug("id in add CheckboxNoteItem", id)
	return suc, id
}

//UpdateCheckboxItem UpdateCheckboxItem
func (c *C2DB) UpdateCheckboxItem(ni *CheckboxNoteItem) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, ni.Text, ni.Checked, ni.ID)
	suc := c.DB.Update(updateCheckboxNoteItem, a...)
	return suc
}

//GetCheckboxItemList GetCheckboxItemList
func (c *C2DB) GetCheckboxItemList(noteID int64) *[]CheckboxNoteItem {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var rtn = []CheckboxNoteItem{}
	var a []interface{}
	a = append(a, noteID)
	rows := c.DB.GetList(getCheckboxNoteItemList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := c.parseCheckboxItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteCheckboxItem DeleteCheckboxItem
func (c *C2DB) DeleteCheckboxItem(id int64) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	return c.DB.Delete(deleteCheckboxNoteItem, a...)
}

func (c *C2DB) parseCheckboxItemRow(foundRow *[]string) *CheckboxNoteItem {
	c.Log.Debug("foundRow in get CheckboxItem", *foundRow)
	var rtn CheckboxNoteItem
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		c.Log.Debug("id err in get CheckboxItem", err)
		if err == nil {
			nid, err := strconv.ParseInt((*foundRow)[3], 10, 64)
			c.Log.Debug("id err in get CheckboxItem", err)
			if err == nil {
				ck, err := strconv.ParseBool((*foundRow)[2])
				if err == nil {
					rtn.ID = id
					rtn.Text = (*foundRow)[1]
					rtn.Checked = ck
					rtn.NoteID = nid
				}
			}
		}
	}
	return &rtn
}
