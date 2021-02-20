package main

import (
	"encoding/json"
	"net/http"
)

func databaseErrorRequest(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return true
	}
	return false
}

func databaseError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func checkEmpty(w http.ResponseWriter, length int) {
	if length == 0 {
		w.WriteHeader(204)
	}
}

// func catchErrors(fn MyFancyFunc) http.HandlerFunc {
// 	return funhttp.HandlerFunc(writer http.HandlerFunc, request *http.Request) {
// 		if err := fn(writer, request); err != nil {
// 			// do something with the error, eg encode it or log it or whatever
// 		}
// 	}
// }

// type (
// 	HttpFunc    func(http.HandlerFunc, *http.Request)
// 	MyFancyFunc func(http.HandlerFunc, *http.Request) error
// )
