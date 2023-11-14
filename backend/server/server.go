package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleInternalServerError(
	w http.ResponseWriter,
	err error,
) {
	w.WriteHeader(http.StatusInternalServerError)
	data := map[string]string{
		"errorMessage": err.Error(),
	}
	jsonErr := json.NewEncoder(w).Encode(data)
	if jsonErr != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "server error: %s\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := map[string]string{"response": "hello"}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			handleInternalServerError(w, err)
			return
		}
	default:
		http.Error(w, "disallowed http method", http.StatusMethodNotAllowed)
	}
}

func Start() {
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}
