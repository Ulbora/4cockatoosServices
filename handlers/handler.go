package handlers

import (
	"net/http"
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

//Handler Handler
type Handler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)

	AddUserToNote(w http.ResponseWriter, r *http.Request)
	GetNoteUserList(w http.ResponseWriter, r *http.Request)

	AddNote(w http.ResponseWriter, r *http.Request)
	UpdateNote(w http.ResponseWriter, r *http.Request)
	GetNote(w http.ResponseWriter, r *http.Request)
	GetUsersNotes(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)

	AddCheckboxItem(w http.ResponseWriter, r *http.Request)
	UpdateCheckboxItem(w http.ResponseWriter, r *http.Request)
	DeleteCheckboxItem(w http.ResponseWriter, r *http.Request)

	AddNoteItem(w http.ResponseWriter, r *http.Request)
	UpdateNoteItem(w http.ResponseWriter, r *http.Request)
	DeleteNoteItem(w http.ResponseWriter, r *http.Request)

	SetLogLevel(w http.ResponseWriter, r *http.Request)
}
