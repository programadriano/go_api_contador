package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
)

func main() {
	rotas := mux.NewRouter()

	rotas.HandleFunc("/api/contador", getAll).Methods("GET")

	var port = ":3500"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}

type Machine struct {
	ValorAtual  int
	MachineName string
	Sistema     string
	Plataforma  string
}

var count = 0

func getAll(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "error"
	}

	count++
	var machine Machine
	machine.ValorAtual = count
	machine.MachineName = hostname
	machine.Sistema = runtime.GOOS
	machine.Plataforma = "Go"

	json.NewEncoder(w).Encode(machine)
}
