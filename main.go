package main

import (
	"GameApp/repository/mysql"
	"GameApp/service/userservice"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/register", Register)
	http.HandleFunc("/user/login", Login)

	port := ":8080"
	log.Println("Server is listening on port", port[1:]+"...")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	var req userservice.RegisterRequest
	if err := json.Unmarshal(data, &req); err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db := mysql.New()
	registeredUser, err := db.Register(req)
	if err != nil {
		http.Error(w, "failed to register user", http.StatusInternalServerError)
		log.Println("error in db.Register:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registeredUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	var req userservice.LoginRequest
	if err := json.Unmarshal(data, &req); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	db := mysql.New()
	ok, err := db.Login(req.PhoneNumber, req.Password)
	if ok == false && err == nil {
		w.Write([]byte(`{"message":"username or password wrong"}`))
	}

	if ok == false && err != nil {
		fmt.Fprintf(w, err.Error())
	}

	if ok == true && err == nil {
		w.Write([]byte(`{"message":"username and password are OK"}`))
	}
}
