package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	m "github.com/Ulbora/cocka2notesServices/managers"
	"github.com/gorilla/mux"
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
func (h *C2Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("add user authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq db.User
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("userReq: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.AddUser(&uReq)
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

//UpdateUser UpdateUser
func (h *C2Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("UpdateUser authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq db.User
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("userReq: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.UpdateUser(&uReq)
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

//GetUser GetUser
func (h *C2Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("GetUser authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var email = vars["email"]
			res := h.Manager.GetUser(email)
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

//Login Login
func (h *C2Handler) Login(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("Login authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq db.User
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("userReq: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.Login(uReq.Email, uReq.Password)
				h.Log.Debug("ures: ", *ures)
				if ures.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
				resJSON, _ := json.Marshal(ures)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uadfl m.LoginResponse
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uadfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//ResetPassword ResetPassword
func (h *C2Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("ResetPassword authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uReq db.User
			usuc, uerr := h.ProcessBody(r, &uReq)
			h.Log.Debug("usuc: ", usuc)
			h.Log.Debug("uerr: ", uerr)
			h.Log.Debug("userReq: ", uReq)
			if !usuc && uerr != nil {
				http.Error(w, uerr.Error(), http.StatusBadRequest)
			} else {
				ures := h.Manager.ResetPassword(uReq.Email)
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
