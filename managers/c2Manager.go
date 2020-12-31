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
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	"golang.org/x/crypto/bcrypt"
)

//C2Manager C2Manager
type C2Manager struct {
	Db  db.NotesDB
	Log *lg.Logger
}

//GetNew GetNew
func (m *C2Manager) GetNew() Manager {
	return m
}

func (m *C2Manager) hashPassword(pw string) (bool, string) {
	var suc bool
	var rtn string
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err == nil {
		rtn = string(hashedPw)
		suc = true
	}
	return suc, rtn
}

//ValidatePassword ValidatePassword
func (m *C2Manager) ValidatePassword(pw string, hpw string) bool {
	var suc bool
	err := bcrypt.CompareHashAndPassword([]byte(hpw), []byte(pw))
	if err == nil {
		suc = true
	}
	return suc
}
