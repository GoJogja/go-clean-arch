package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GoJogja/go-clean-arch/context/addperson"
	"github.com/GoJogja/go-clean-arch/context/greeting"
	"github.com/GoJogja/go-clean-arch/context/listperson"
	"github.com/GoJogja/go-clean-arch/store"
	"github.com/GoJogja/go-clean-arch/transport/httphandler"
	addpersonhttp "github.com/GoJogja/go-clean-arch/transport/httphandler/addperson"
	greetinghttp "github.com/GoJogja/go-clean-arch/transport/httphandler/greeting"
	listpersonhttp "github.com/GoJogja/go-clean-arch/transport/httphandler/listperson"
	"github.com/gorilla/mux"
)

func main() {

	muxrr := httphandler.NewMuxRequestReader()

	guc := greeting.NewGreetingUseCase()
	gh := greetinghttp.NewGreetingHandler(guc, muxrr)

	par := store.NewPersonArrayStore()

	apuc := addperson.NewAddPersonUseCase(par)
	aph := addpersonhttp.NewHandler(apuc, muxrr)

	lpuc := listperson.NewListPersonUseCase(par)
	lph := listpersonhttp.NewHandler(lpuc, muxrr)

	router := mux.NewRouter()
	router.HandleFunc("/greet/{name}", gh.Greet).Methods("GET")
	router.HandleFunc("/person/add", aph.Handle).Methods("POST")
	router.HandleFunc("/person/list", lph.Handle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", logger(router)))
}

func logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
