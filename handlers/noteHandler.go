package handlers

import (
	"encoding/json"
	"fmt"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	m "github.com/Ulbora/cocka2notesServices/managers"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

//AddNote AddNote
func (h *C2Handler) AddNote(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("AddNote authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq db.Note
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("userReq: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.AddNote(&uReq)
				h.Log.Debug("ures: ", *ures)
				if ures.Success && ures.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ures)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uadfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uadfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateNote UpdateNote
func (h *C2Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("UpdateNote authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq db.Note
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("Req: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.UpdateNote(&uReq)
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

//GetNote GetNote
func (h *C2Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("GetNote authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var idStr = vars["id"]
			id, iderr := strconv.ParseInt(idStr, 10, 64)
			var res *m.Note
			if iderr == nil {
				res = h.Manager.GetNote(id)
				// h.Log.Debug("res: ", *res)
				// ni := res.NoteItems
				// h.Log.Debug("ni: ", *ni.(*[]db.CheckboxNoteItem))
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Note
				res = &nc
			}
			resJSON, _ := json.Marshal(res)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetUsersNotes GetUsersNotes
func (h *C2Handler) GetUsersNotes(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("GetUsersNotes authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars len: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var email = vars["email"]
			var res *[]db.Note
			res = h.Manager.GetUsersNotes(email)
			h.Log.Debug("res: ", res)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(res)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteNote DeleteNote
func (h *C2Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("DeleteNote authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var idStr = vars["id"]
			var email = vars["email"]
			id, iderr := strconv.ParseInt(idStr, 10, 64)
			var res *m.Response
			if iderr == nil {
				res = h.Manager.DeleteNote(id, email)
				h.Log.Debug("res: ", res)
				if res.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				res = &nc
			}
			resJSON, _ := json.Marshal(res)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
