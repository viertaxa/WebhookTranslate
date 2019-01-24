package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

type Configuration struct {
	WebhookIDs      []string
	Port            int
	ListeningServer string
}

var conf Configuration

func main() {
	if _, err := toml.DecodeFile("settings.toml", &conf); err != nil {
		log.Fatal(err)
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/{hook}", HookHandler)
	listener := ":" + strconv.Itoa(conf.Port)
	log.Fatal(http.ListenAndServe(listener, r))
}

func HookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if stringInSlice(vars["hook"], conf.WebhookIDs) {
		log.Println("Processing Hook: " + vars["hook"][0:3])
		req, err := http.NewRequest("POST", conf.ListeningServer+"/api/webhook/"+vars["hook"], nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			log.Println("response Status:", resp.Status)
			log.Println("response Headers:", resp.Header)
			return
		}

		w.WriteHeader(resp.StatusCode)
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
	} else {
		log.Println("Ignored req:", vars["hook"])
		return
	}

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
