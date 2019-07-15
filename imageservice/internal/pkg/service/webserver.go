package service

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func StartWebServer(serviceName, port string) {

	r := NewRouter(serviceName)
	http.Handle("/", r)
	logrus.Infof("Starting HTTP service at %v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
