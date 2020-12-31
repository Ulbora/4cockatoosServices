package handlers

import (
	"encoding/json"
	"errors"
	lg "github.com/Ulbora/Level_Logger"
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

//C2Handler C2Handler
type C2Handler struct {
	Manager m.Manager
	Log     *lg.Logger
	APIKey  string
}

//GetNew GetNew
func (h *C2Handler) GetNew() Handler {
	return h
}

//CheckContent CheckContent
func (h *C2Handler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		// http.Error(w, "json required", http.StatusUnsupportedMediaType)
		rtn = true
	}
	return rtn
}

//SetContentType SetContentType
func (h *C2Handler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//ProcessBody ProcessBody
func (h *C2Handler) ProcessBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	h.Log.Debug("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}

func (h *C2Handler) processAPIKeySecurity(r *http.Request) bool {
	var rtn bool
	apiKey := r.Header.Get("apiKey")
	h.Log.Debug("apiKey: ", apiKey)
	h.Log.Debug("h.APIKey: ", h.APIKey)
	if apiKey == h.APIKey {
		rtn = true
	}
	return rtn
}
