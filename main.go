package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
	"teemo-trap/mushroom"
)

func main() {
	configData, err := ioutil.ReadFile("./application.yml")
	config := mushroom.Config{}
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalln(err)
	}

	mushroom.InitDatabase(&config.Mongo)
	if err != nil {
		log.Fatalln(err)
	}

	var router = mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!"))
	})
	log.Fatalln(http.ListenAndServe(":12345", router))
}
