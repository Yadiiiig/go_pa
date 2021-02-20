package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func authenticationCheck(request http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "" {
			if auth == authKey {
				request(w, r)
			} else {
				nonAuthRequest(r.RemoteAddr)
				w.WriteHeader(403)
				json.NewEncoder(w).Encode("What are you trying to accomplish?")
			}
		} else {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode("What are you trying to accomplish?")
		}
	})
}

func nonAuthRequest(ip string) {
	returnedIP := selectedIP{}

	rows, err := db.Queryx("SELECT ip, tries, blocked FROM denylist WHERE ip = ?", ip)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		err := rows.StructScan(&returnedIP)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	rows.Close()

	if returnedIP.IP == "" {
		_, err := db.Query("INSERT INTO denylist (ip, tries) VALUES (?, 1)", ip)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if returnedIP.Tries != 4 {
		_, err := db.Query("UPDATE denylist SET tries = ? WHERE ip = ?", returnedIP.Tries+1, ip)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		_, err := db.Query("UPDATE denylist SET tries = ?, blocked = 1 WHERE ip = ?", returnedIP.Tries+1, ip)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Add code to add IP to map
	}

}

type selectedIP struct {
	IP      string `db:"ip"`
	Tries   int    `db:"tries"`
	Blocked int    `db:"blocked"`
}
