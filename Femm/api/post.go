package api

import (
	"encoding/json"
	"net/http"

	"frontendmasters.com/go/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, "Uneupported Method", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	}else{
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}