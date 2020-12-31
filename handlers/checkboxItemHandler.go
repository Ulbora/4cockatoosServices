package handlers

import (
	"encoding/json"
	"fmt"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	m "github.com/Ulbora/cocka2notesServices/managers"
	"github.com/gorilla/mux"
	"strconv"

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

//AddCheckboxItem AddCheckboxItem
func (h *C2Handler) AddCheckboxItem(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("AddCheckboxItem authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var req db.CheckboxNoteItem
			suc, err := h.ProcessBody(r, &req)
			h.Log.Debug("suc: ", suc)
			h.Log.Debug("err: ", err)
			h.Log.Debug("req: ", req)
			if !suc && err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				res := h.Manager.AddCheckboxItem(&req)
				h.Log.Debug("res: ", *res)
				if res.Success && res.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(res)
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

//UpdateCheckboxItem UpdateCheckboxItem
func (h *C2Handler) UpdateCheckboxItem(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("UpdateCheckboxItem authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var req db.CheckboxNoteItem
			suc, err := h.ProcessBody(r, &req)
			h.Log.Debug("suc: ", suc)
			h.Log.Debug("err: ", err)
			h.Log.Debug("Req: ", req)
			if !suc && err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				res := h.Manager.UpdateCheckboxItem(&req)
				h.Log.Debug("res: ", *res)
				if res.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(res)
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

//DeleteCheckboxItem DeleteCheckboxItem
func (h *C2Handler) DeleteCheckboxItem(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("DeleteCheckboxItem authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var idStr = vars["id"]
			id, iderr := strconv.ParseInt(idStr, 10, 64)
			var res *m.Response
			if iderr == nil {
				res = h.Manager.DeleteCheckboxItem(id)
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
