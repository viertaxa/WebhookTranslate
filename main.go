package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

}

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
		w.WriteHeader(http.StatusOK)
		fmt.Println("Processing Hook: " + vars["hook"][0:3])
		req, err := http.NewRequest("POST", conf.ListeningServer+"/api/webhook/"+vars["hook"], nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)

		println(resp.Status)
	} else {
		fmt.Println("Ignored req:", vars["hook"])
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
