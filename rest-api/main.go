package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

type Profile struct {
	Hobby string `json:"hobby"`
	UserID int `json:"userId"`
}

var profiles []Profile

func addProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProfile Profile
	err := json.NewDecoder(r.Body).Decode(&newProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func getAllProfiles(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		// w.WriteHeader(400)
		// w.Write([]byte("id could not be converted to int"))
		// 以下の方が可動性高そう
		http.Error(w, "id could not be converted to int", http.StatusBadRequest)
		return
	}

	for _, profile := range profiles {
		if profile.UserID == id {
			json.NewEncoder(w).Encode(profile)
			return
		}
	}

	http.Error(w, "profile not found", http.StatusNotFound)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "id could not be converted to int", http.StatusBadRequest)
		return
	}

	var updatedProfile Profile
	err = json.NewDecoder(r.Body).Decode(&updatedProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, profile := range profiles {
		if profile.UserID == id {
			profiles[i] = updatedProfile
			json.NewEncoder(w).Encode(profiles)
			return
		}
	}

	http.Error(w, "profile not found", http.StatusNotFound)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "id could not be converted to int", http.StatusBadRequest)
		return
	}

	for i, profile := range profiles {
		if profile.UserID == id {
			profiles = append(profiles[:i], profiles[i+1:]...)
			json.NewEncoder(w).Encode(profiles)
			return
		}
	}

	http.Error(w, "profile not found", http.StatusNotFound)
}

// cors処理をmiddlewareに切り出す(ディレクトリも分けたほうがbetterかも)
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 他ドメインからのアクセスされることを想定して、CORSを許可
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Preflight request https://qiita.com/popo62520908/items/24956284e40871497082
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func main() {
	router := mux.NewRouter()

	router.Handle("/profiles", enableCORS(http.HandlerFunc(addProfile))).Methods("POST")

	router.Handle("/profiles", enableCORS(http.HandlerFunc(getAllProfiles))).Methods("GET")

	router.Handle("/profiles/{id}", enableCORS(http.HandlerFunc(getProfile))).Methods("GET")

	router.Handle("/profiles/{id}", enableCORS(http.HandlerFunc(updateProfile))).Methods("PUT")

	router.Handle("/profiles/{id}", enableCORS(http.HandlerFunc(deleteProfile))).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

// api叩くコマンド
// - http http://localhost:8080/profiles hobby="tennis" userId:=1
// - http http://localhost:8080/profiles
// - http PUT http://localhost:8080/profiles/1 hobby="soccer" userId:=1
// - http DELETE http://localhost:8080/profiles/1
