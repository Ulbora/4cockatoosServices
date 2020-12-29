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

//AddNoteUser AddNoteUser
func (c *C2DB) AddNoteUser(nu *NoteUsers) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, nu.NoteID, nu.UserEmail)
	suc, id := c.DB.Insert(insertNoteUser, a...)
	c.Log.Debug("suc in add NoteUsers", suc)
	c.Log.Debug("id in add NoteUsers", id)
	return suc
}

//DeleteNoteUser DeleteNoteUser
func (c *C2DB) DeleteNoteUser(nu *NoteUsers) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, nu.NoteID, nu.UserEmail)
	return c.DB.Delete(deleteNoteUser, a...)
}
