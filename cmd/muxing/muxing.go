package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", GetName).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadRequest).Methods(http.MethodGet)
	router.HandleFunc("/data", PostParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", SetNewHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := "localhost"
	port := 8080
	//port, err := strconv.Atoi(os.Getenv("PORT"))
	//if err != nil {
	//	port = 8081
	//}
	Start(host, port)
}

func GetName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body := fmt.Sprintf("Hello, %s!", params["param"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
	return

}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func PostParam(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("I got message:\n%s", body)
	w.Write([]byte(response))
	w.WriteHeader(http.StatusOK)
}

func SetNewHeaders(w http.ResponseWriter, r *http.Request) {

	aStr := r.Header.Get("a")
	bStr := r.Header.Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := strconv.Itoa(a + b)
	arr := []string{res}
	header := w.Header()
	header["a+b"] = arr
	//w.Header().Set("a+b", res)
	w.WriteHeader(http.StatusOK)

}
