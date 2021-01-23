package managers

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
	b64 "encoding/base64"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

//GetMailServer GetMailServer
func (m *C2Manager) GetMailServer() *db.MailServer {
	m.Log.Debug("GetMailServer ")
	rtn := m.Db.GetMailServerInfo()
	if rtn != nil {
		sDec, _ := b64.StdEncoding.DecodeString(rtn.Password)
		rtn.Password = string(sDec)
	}
	return rtn
}
