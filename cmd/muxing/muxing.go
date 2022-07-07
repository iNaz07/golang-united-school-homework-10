package main

import (
	"fmt"
	"github.com/gorilla/mux"
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

	router.HandleFunc("/name", GetName)
	router.HandleFunc("/bad", BadRequest)
	router.HandleFunc("/data", PostParam)
	router.HandleFunc("/headers", SetNewHeaders)

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
	if r.Method == http.MethodGet {

		body := fmt.Sprintf("Hello, %s!", r.FormValue("name"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func PostParam(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body []byte

		_, err := r.Body.Read(body)
		if err != nil {
			w.Write([]byte("unable to read body"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response := fmt.Sprintf("I got message:\n%s", body)
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func SetNewHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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
