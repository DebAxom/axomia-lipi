package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func TranslationHanler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	if word, WordExists := vars["word"]; WordExists {
		word = TranslateWord(word)
		json.NewEncoder(w).Encode(map[string]any{"word": word, "error": false})
	} else {
		json.NewEncoder(w).Encode(map[string]any{"word": nil, "error": true})
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/en-as/{word}", TranslationHanler)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running on localhost:8000")
	srv.ListenAndServe()
}
