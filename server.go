package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
)

var (
	listHolidayRe = regexp.MustCompile(`^\/holiday[\/]*$`)
	getHolidayRe  = regexp.MustCompile(`^\/holiday\/year\/(\d+)$`)
)

type Data []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

type datastore struct {
	m map[string]Data
	*sync.RWMutex
}

type dateHandler struct {
	store *datastore
}

func (h *dateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listHolidayRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getHolidayRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *dateHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	holiday := make([]Data, 0, len(h.store.m))
	for _, v := range h.store.m {
		holiday = append(holiday, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.MarshalIndent(holiday, " ", " ")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (h *dateHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getHolidayRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	y, ok := h.store.m[matches[1]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("year not found"))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	jsonBytes, err := json.MarshalIndent(y, " ", " ")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("internal server error"))
	if err != nil {
		log.Fatal(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("not found"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	h1 := holiday_2021()
	h2 := holiday_2022()
	h3 := holiday_2023()

	mux := http.NewServeMux()
	dHandler := &dateHandler{
		store: &datastore{
			m: map[string]Data{
				"2021": Data(h1),
				"2022": Data(h2),
				"2023": Data(h3),
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/holiday", dHandler)
	mux.Handle("/holiday/", dHandler)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
