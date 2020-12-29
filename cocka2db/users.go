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

//AddUser AddUser
func (c *C2DB) AddUser(u *User) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, u.Email, u.Password, u.WebEnabled)
	suc, id := c.DB.Insert(insertUser, a...)
	c.Log.Debug("suc in add User", suc)
	c.Log.Debug("id in add User", id)
	return suc
}

//UpdateUser UpdateUser
func (c *C2DB) UpdateUser(u *User) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, u.Password, u.WebEnabled, u.Email)
	suc := c.DB.Update(updateUser, a...)
	return suc
}

//GetUser GetUser
func (c *C2DB) GetUser(email string) *User {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, email)
	row := c.DB.Get(getUser, a...)
	c.Log.Debug("foundRow in get user", row.Row)
	var rtn User
	if len(row.Row) > 0 {
		we, err := strconv.ParseBool(row.Row[2])
		if err == nil {
			rtn.WebEnabled = we
			rtn.Email = row.Row[0]
			rtn.Password = row.Row[1]
		}
	}
	return &rtn
}

//DeleteUser DeleteUser
func (c *C2DB) DeleteUser(email string) bool {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var a []interface{}
	a = append(a, email)
	return c.DB.Delete(deleteUser, a...)
}
