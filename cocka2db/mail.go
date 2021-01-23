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

//GetMailServerInfo GetMailServerInfo
func (c *C2DB) GetMailServerInfo() *MailServer {
	if !c.testConnection() {
		c.DB.Connect()
	}
	var rtn = MailServer{}
	var a []interface{}
	a = append(a)
	rows := c.DB.GetList(getMailServer, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rtn = c.parseMailServerRow(&foundRow)
			break
		}
	}
	return &rtn
}

func (c *C2DB) parseMailServerRow(foundRow *[]string) MailServer {
	c.Log.Debug("foundRow in get MailServer", *foundRow)
	var rtn MailServer
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		c.Log.Debug("id err in get note", err)
		if err == nil {
			rtn.ID = id
			rtn.Host = (*foundRow)[1]
			rtn.Port = (*foundRow)[3]
			rtn.Username = (*foundRow)[5]
			rtn.Password = (*foundRow)[6]
			rtn.SenderEmail = (*foundRow)[7]
		}
	}
	return rtn
}
