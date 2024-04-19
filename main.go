package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

var addressBookData []addressBook 

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(addressBookData)
}

func addAddressBook(w http.ResponseWriter, r *http.Request) {
	var newEntry addressBook
	err := json.NewDecoder(r.Body).Decode(&newEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addressBookData = append(addressBookData, newEntry)
	fmt.Fprintf(w, "Address book entry added successfully!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

//---handle---//
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.HandleFunc("/addAddress", addAddressBook) 
	http.ListenAndServe(":5050", nil)
}

func loadAddressBookData() {
	file, err := os.Open("address_book_data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&addressBookData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func main() {
	loadAddressBookData()
	handleRequest()
}
