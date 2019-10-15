package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//what is json:"ID"
type event struct {
	ID          string `json:"ID"` //still don't know what those line are for
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

//http.ResponseWriter interface is used by an HTTP handler to construct an HTTP response
//http.Request represents an HTTP request received by the server or to be sent by a client
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//ioutil.ReadAll(r.Body) ReadAll read from r until an error or EOF and returns the data it read. 
//A successfull call return err == nil and not err == EOF
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newEvent) //(src, dest)
	events = append(events, newEvent) 
	w.WriteHeader(http.StatusCreated) //still not sure about what WriteHeader is for
	json.NewEncoder(w).Encode(newEvent) //newEncoder returns a new encoder that writes to w.
	//Encode writes the JSON encoding of v to the stream, followed by a newline character.
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"] //return the route variable for the current request and id == {id}
	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent) //newEncoder returns a new encoder taht writes to w.
			//Encode writes the JSON encoding of v to the stream, followed by a newline character.
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events) //newEncoder returns a new encoder that writes to w.
	//Encode writes the JSON encoding of v to the stream, followed by a newline
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"] //return the route variable for the current request
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body) //ioutil.ReadAll(r.Body) read from r until an error or EOF and returns the data it read.
	//A successfull call return err == nil and not err == EOF
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent) //(src, dest)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent) // (src, dest) i include
			json.NewEncoder(w).Encode(singleEvent) //newEncoder returns a new encoder that writes to w.
			//Encode writes the JSON encoding of v to the stream, followed by a newline character.
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"] //return the route variable for the current request

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			partOne := events[:i]
			partTwo := events[i+1:]
			events = append(partOne, partTwo...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true) //they don't care about the last slash
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router)) //Fatal is equivalent to Printf() followed by a call to os.Exit(1)
}