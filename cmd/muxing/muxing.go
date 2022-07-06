package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", handleGetParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handlerGetBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handlerData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeader).Methods(http.MethodPost)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}
func handlerGetParam(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func handlerGetBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
func handlerData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "I got message:\n%v", string(b))
}
func handleHeader(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Set("a+b", fmt.Sprint(a+b))
	w.WriteHeader(http.StatusOK)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
