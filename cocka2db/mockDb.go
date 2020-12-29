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

import (
	lg "github.com/Ulbora/Level_Logger"
	dbi "github.com/Ulbora/dbinterface"
)

//MockC2DB MockC2DB
type MockC2DB struct {
	DB  dbi.Database
	Log *lg.Logger

	MockAddUserSuc    bool
	MockUpdateUserSuc bool
	MockUser          *User
	MockDeleteUserSuc bool
}

//AddUser AddUser
func (c *MockC2DB) AddUser(u *User) bool {
	return c.MockAddUserSuc
}

//UpdateUser UpdateUser
func (c *MockC2DB) UpdateUser(u *User) bool {
	return c.MockUpdateUserSuc
}

//GetUser GetUser
func (c *MockC2DB) GetUser(email string) *User {
	return c.MockUser
}

//DeleteUser DeleteUser
func (c *MockC2DB) DeleteUser(email string) bool {
	return c.MockDeleteUserSuc
}
