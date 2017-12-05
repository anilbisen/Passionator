package main

import (
	"encoding/json"
	"net/http"
	"time"

	"goji.io"

	"goji.io/pat"
	"gopkg.in/mgo.v2/bson"
)

type (
	// ProfileController represents the controller for working with this app
	ProfileController struct {
		Session SessionWrapper
	}
)

// NewProfileController provides a reference to an EventController with
func NewProfileController(session SessionWrapper) *ProfileController {
	return &ProfileController{session}
}

// AddHandlers inserts new Profile
func (pc *ProfileController) AddHandlers(mux *goji.Mux) {
	mux.HandleFunc(pat.Post("/v1/profiles"), Logger(pc.CreateProfile))
	mux.HandleFunc(pat.Get("/v1/profiles"), Logger(pc.GetProfiles))
	mux.HandleFunc(pat.Get("/v1/profiles/:id"), Logger(pc.GetProfile))
	mux.HandleFunc(pat.Delete("/v1/profiles/:id"), Logger(pc.DeleteProfile))
	mux.HandleFunc(pat.Put("/v1/profiles/:id"), Logger(pc.UpdateProfile))
}

// CreateProfile inserts new Profile
func (pc *ProfileController) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var pm ProfileModel
	err := json.NewDecoder(r.Body).Decode(&pm)
	if err != nil {
		MainLogger.Printf("Error decoding  body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pm.ID = bson.NewObjectId()
	pm.CreatedAt = time.Now()
	pm.UpdatedAt = pm.CreatedAt
	pc.Session.DB("passionatordb").C("Profiles").Insert(pm)
	pmj, err := json.Marshal(pm)
	if err != nil {
		MainLogger.Println("Error marshaling into JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(pmj)
}

// GetProfiles retrieves all Profiles
func (pc *ProfileController) GetProfiles(w http.ResponseWriter, r *http.Request) {
	var pms []ProfileModel
	err := pc.Session.DB("passionatordb").C("Profiles").Find(ExtractQuery(r)).All(&pms)
	if err != nil {
		MainLogger.Println("Error reading from db")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	pmsj, err := json.Marshal(pms)
	if err != nil {
		MainLogger.Println("Error marshaling into JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pmsj)
}

// GetProfile retrieves specific Profile
func (pc *ProfileController) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r, "id")
	if !bson.IsObjectIdHex(id) {
		MainLogger.Println("Invalid id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idh := bson.ObjectIdHex(id)
	pm := ProfileModel{}
	err := pc.Session.DB("passionatordb").C("Profiles").FindId(idh).One(&pm)
	if err != nil {
		MainLogger.Println("Unknown id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	pmj, err := json.Marshal(pm)
	if err != nil {
		MainLogger.Println("Error marshaling into JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pmj)
}

// UpdateProfile upwrite new data to existing Profile
func (pc *ProfileController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r, "id")
	if !bson.IsObjectIdHex(id) {
		MainLogger.Println("Invalid id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idh := bson.ObjectIdHex(id)
	pmdb := ProfileModel{}
	err := pc.Session.DB("passionatordb").C("Profiles").FindId(idh).One(&pmdb)
	if err != nil {
		MainLogger.Println("Unknown id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	pmreq := ProfileModel{}
	err = json.NewDecoder(r.Body).Decode(&pmreq)
	if err != nil {
		MainLogger.Println("Error decoding body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pmdb.UpdatedAt = time.Now()
	pmdb.Clone(pmreq)
	err = pc.Session.DB("passionatordb").C("Profiles").UpdateId(idh, pmdb)
	if err != nil {
		MainLogger.Printf("Error updating database: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProfile removes existing Profile
func (pc *ProfileController) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r, "id")
	if !bson.IsObjectIdHex(id) {
		MainLogger.Println("Invalid id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idh := bson.ObjectIdHex(id)
	err := pc.Session.DB("passionatordb").C("Profiles").RemoveId(idh)
	if err != nil {
		MainLogger.Println("Unknown id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
