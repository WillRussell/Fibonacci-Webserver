package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "strings"
    "encoding/json"
)



func parseInput(s string) int {
	string := strings.Replace(s, "/", "", -1) // parse non-numeric characters 
	number, _ := strconv.Atoi(string) // convert it to an integer
	return int(number)
}



func fibonacci() func() int { 
    x, y := 0, 1
    return func() (r int) {
        r = x
        x, y = y, x + y
        return 
    }
}


func generateFibonacci(w http.ResponseWriter, r * http.Request) {

    var collection[] int
    var placeholder int

    parsedNumber := parseInput(r.URL.Path)
    f := fibonacci()

    for i := 0; i < parsedNumber; i++ {
        placeholder = f()
        collection = append(collection, placeholder)
    }

    result, _ := json.Marshal(collection)

    w.Header().Set("Access-Control-Allow-Origin", "*") // allow cross domain AJAX requests
    w.Header().Set("Content-Type", "application/json") // set content-type so client knows to expect json
    w.WriteHeader(201)
    w.Write(result)
}


func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/{id:[0-9]+}", generateFibonacci).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}


func main() {
    fmt.Println("Starting server on 8080...")
    handleRequests()
}