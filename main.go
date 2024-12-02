package main

import (
	"encoding/json"
	"fmt"
	"game/repository/mysql"
	"game/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)

	log.Println("server is listening on port 8080...")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
	}

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
		return
	}

	writer.Write([]byte(`{"message": "user created"}`))
}
func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message": "everything is good!"}`)
}