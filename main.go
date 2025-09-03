package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/register", Register)

	port := ":8080"
	log.Println("Server is listening at port", port[1:]+"...")
	http.ListenAndServe(port, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "invalid method")
	} else {
		/*		data, _ := io.ReadAll(r.Body)
				RegisterReq := userservice.RegisterRequest{}
				err := json.Unmarshal(data, &RegisterReq)
				if err != nil {
					log.Println(err)
				}
				db := mysql.New()
				RegisteredUser, err := db.Register(RegisterReq)
				fmt.Print(string(data))*/
	}
}
