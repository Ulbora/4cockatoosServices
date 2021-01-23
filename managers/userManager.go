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

//AddUser AddUser
func (m *C2Manager) AddUser(u *db.User) *Response {
	m.Log.Debug("add user: ", *u)
	var rtn Response
	var pwsuc = true

	if u.Password != "" {
		pwsuc, u.Password = m.hashPassword(u.Password)
	}
	if pwsuc {
		m.Log.Debug("adding user: ", *u)
		suc := m.Db.AddUser(u)
		rtn.Success = suc
	}
	return &rtn
}

//UpdateUser UpdateUser
func (m *C2Manager) UpdateUser(u *db.User) *Response {
	m.Log.Debug("UpdateUser: ", *u)
	var rtn Response
	var pwsuc = true

	if u.Password != "" {
		pwsuc, u.Password = m.hashPassword(u.Password)
	}
	if pwsuc {
		fu := m.Db.GetUser(u.Email)
		if fu.Email != "" {
			fu.Password = u.Password
			fu.WebEnabled = u.WebEnabled
			m.Log.Debug("Updated User: ", *fu)
			suc := m.Db.UpdateUser(fu)
			rtn.Success = suc
		}
	}
	return &rtn
}

//GetUser GetUser
func (m *C2Manager) GetUser(email string) *db.User {
	m.Log.Debug("GetUser: ", email)
	rtn := m.Db.GetUser(email)
	rtn.Password = ""
	return rtn
}

//Login Login
func (m *C2Manager) Login(email string, pw string) *LoginResponse {
	var rtn LoginResponse
	m.Log.Debug("GetUser: ", email)
	usr := m.Db.GetUser(email)
	if usr.Email != "" {
		suc := m.ValidatePassword(pw, usr.Password)
		rtn.Email = usr.Email
		rtn.Success = suc
	}
	return &rtn
}
