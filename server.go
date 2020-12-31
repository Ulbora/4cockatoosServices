package main

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

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	cacka2db "github.com/Ulbora/cocka2notesServices/cocka2db"
	hand "github.com/Ulbora/cocka2notesServices/handlers"
	man "github.com/Ulbora/cocka2notesServices/managers"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var mydb mdb.MyDB

	var dbHost string
	var dbUser string
	var dbPassword string
	var dbName string

	var apiKey string

	var l lg.Logger
	l.LogLevel = lg.AllLevel

	if os.Getenv("COCKA2_DB_HOST") != "" {
		dbHost = os.Getenv("COCKA2_DB_HOST")
	} else {
		dbHost = "localhost:3306"
	}

	if os.Getenv("COCKA2_DB_USER") != "" {
		dbUser = os.Getenv("COCKA2_DB_USER")
	} else {
		dbUser = "admin"
	}

	if os.Getenv("COCKA2_DB_PASSWORD") != "" {
		dbPassword = os.Getenv("COCKA2_DB_PASSWORD")
	} else {
		dbPassword = "admin"
	}

	if os.Getenv("COCKA2_DB_DATABASE") != "" {
		dbName = os.Getenv("COCKA2_DB_DATABASE")
	} else {
		dbName = "cocka2_notes"
	}

	if os.Getenv("COCKA2_API_KEY") != "" {
		apiKey = os.Getenv("COCKA2_API_KEY")
	} else {
		apiKey = "GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5"
	}

	mydb.Host = dbHost         // "localhost:3306"
	mydb.User = dbUser         // "admin"
	mydb.Password = dbPassword // "admin"
	mydb.Database = dbName     // "six910"
	var dbi db.Database = &mydb

	var cdb cacka2db.C2DB

	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	var c2m man.C2Manager
	c2m.Db = cdb.GetNew()
	c2m.Log = &l

	var sh hand.C2Handler
	sh.Manager = c2m.GetNew()
	sh.APIKey = apiKey
	sh.Log = &l

	router := mux.NewRouter()
	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	h := sh.GetNew()

	router.HandleFunc("/rs/user/add", h.AddUser).Methods("POST")
	router.HandleFunc("/rs/user/update", h.UpdateUser).Methods("PUT")

	router.HandleFunc("/rs/note/user/add", h.AddUserToNote).Methods("POST")

	router.HandleFunc("/rs/note/add", h.AddNote).Methods("POST")
	router.HandleFunc("/rs/note/update", h.UpdateNote).Methods("PUT")
	router.HandleFunc("/rs/note/get/{id}", h.GetNote).Methods("GET")
	router.HandleFunc("/rs/note/get/all/{email}", h.GetUsersNotes).Methods("GET")
	router.HandleFunc("/rs/note/delete/{id}/{email}", h.DeleteNote).Methods("DELETE")

	router.HandleFunc("/rs/checkbox/add", h.AddCheckboxItem).Methods("POST")
	router.HandleFunc("/rs/checkbox/update", h.UpdateCheckboxItem).Methods("PUT")
	router.HandleFunc("/rs/checkbox/delete/{id}", h.DeleteCheckboxItem).Methods("DELETE")

	router.HandleFunc("/rs/item/add", h.AddNoteItem).Methods("POST")
	router.HandleFunc("/rs/item/update", h.UpdateNoteItem).Methods("PUT")
	router.HandleFunc("/rs/item/delete/{id}", h.DeleteNoteItem).Methods("DELETE")

	fmt.Println("Cocka2Services server is running on port " + port + "!")

	http.ListenAndServe(":"+port, router)

}
