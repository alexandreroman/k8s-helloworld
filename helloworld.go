/*
 * Hello World app
 * Copyright (c) 2018 Pivotal Software, Inc.
 */

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Greetings struct {
	Hostname string
	Message  string
}

func _GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Cannot get hostname", err)
	}
	return hostname
}

func _GetServerPort() int {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		intPort, err := strconv.Atoi(port)
		if err != nil {
			log.Fatal("Cannot parse environment variable: SERVER_PORT")
		}
		return intPort
	}
	return 9000
}

func _GetMessage() string {
	msg := os.Getenv("MESSAGE")
	return msg
}

func _HandleRequests(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" || r.URL.Path == "/" {
		_SayHello(w, r)
	} else if r.URL.Path == "/logo.png" {
		http.ServeFile(w, r, "logo.png")
	}
}

func _SayHello(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("helloworld.html"))
	data := Greetings{
		Hostname: _GetHostname(),
		Message:  _GetMessage(),
	}
	log.Printf("Sending page to client")
	tpl.Execute(w, data)
}

func main() {
	log.Printf("Process PID: %d", os.Getpid())

	port := _GetServerPort()
	log.Printf("Listening on port: %d", port)

	http.HandleFunc("/", _HandleRequests)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("Cannot start web server", err)
	}
}
