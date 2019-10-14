package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.optum.com/ssistla1/go-rest-api/person"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	router.HandleFunc("/emps", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(person.GetPeople())
	}).Methods("GET")

	router.HandleFunc("/emps/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			w.WriteHeader(400)
		} else {
			var p, e = person.GetPersonByID(id)
			if e == nil {
				json.NewEncoder(w).Encode(p)
			} else {
				w.WriteHeader(422)
			}
		}
	}).Methods("GET")

	router.HandleFunc("/emps/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			w.WriteHeader(400)
		} else {
			var p, e = person.DeletePersonByID(id)
			if e == nil {
				json.NewEncoder(w).Encode(p)
			} else {
				w.WriteHeader(422)
			}
		}
	}).Methods("DELETE")

	router.HandleFunc("/emp", func(w http.ResponseWriter, r *http.Request) {
		var reqBody, _ = ioutil.ReadAll(r.Body)
		var p person.Person
		json.Unmarshal(reqBody, &p)
		if p.Name == "" {
			w.WriteHeader(400)
		} else {
			var p = person.CreatePerson(p)
			json.NewEncoder(w).Encode(p)
		}
	}).Methods("POST")

	router.HandleFunc("/emp", func(w http.ResponseWriter, r *http.Request) {
		var reqBody, _ = ioutil.ReadAll(r.Body)
		var p person.Person
		json.Unmarshal(reqBody, &p)
		if p.Name == "" {
			w.WriteHeader(400)
		} else {
			var p, _ = person.UpdatePerson(p)
			json.NewEncoder(w).Encode(p)
		}
	}).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
