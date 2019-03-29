package main

import (
	"encoding/json"
	"net/http"

	"github.com/skillitzimberg/surfapi/datautil"
)

func main() {
	http.HandleFunc("/", dataToJson)
	http.ListenAndServe(":3000", nil)
}

func dataToJson(w http.ResponseWriter, r *http.Request) {
	rawBouyData := datautil.GetBouyData()
	bouyData := datautil.HandleRawData(rawBouyData)
	packagedBouyData := datautil.DataToStructs(bouyData)

	js, err := json.Marshal(packagedBouyData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
