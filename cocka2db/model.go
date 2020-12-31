package cocka2db

import (
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

//User User
type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	WebEnabled bool   `json:"webEnabled"`
}

//Note Note
type Note struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Type       string    `json:"type"`
	OwnerEmail string    `json:"ownerEmail"`
	LastUsed   time.Time `json:"lastUsed"`
}

//NoteUsers NoteUsers
type NoteUsers struct {
	UserEmail string `json:"userEmail"`
	NoteID    int64  `json:"noteId"`
}

//CheckboxNoteItem CheckboxNoteItem
type CheckboxNoteItem struct {
	ID      int64  `json:"id"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
	NoteID  int64  `json:"noteId"`
}

//NoteItem NoteItem
type NoteItem struct {
	ID     int64  `json:"id"`
	Text   string `json:"text"`
	NoteID int64  `json:"noteId"`
}
