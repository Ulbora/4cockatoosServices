package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"

	m "github.com/Ulbora/cocka2notesServices/managers"
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

//AddUserToNote AddUserToNote
func (h *C2Handler) AddUserToNote(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("AddUserToNote authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq m.NoteUsers
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("userReq: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.AddUserToNote(&uReq)
				h.Log.Debug("ures: ", *ures)
				if ures.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ures)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uadfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uadfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetNoteUserList GetNoteUserList
func (h *C2Handler) GetNoteUserList(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("GetNoteUserList authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars len: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var idStr = vars["noteId"]
			noteID, iderr := strconv.ParseInt(idStr, 10, 64)
			var res *[]string
			if iderr == nil {
				var ownerEmail = vars["ownerEmail"]
				res = h.Manager.GetNoteUserList(noteID, ownerEmail)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nr []string
				res = &nr
			}
			h.Log.Debug("res: ", res)
			resJSON, _ := json.Marshal(res)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
