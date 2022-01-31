package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Groupe struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {

	GetData("https://groupietrackers.herokuapp.com/api/artists")
	GetData("https://groupietrackers.herokuapp.com/api/artists")
	GetData("https://groupietrackers.herokuapp.com/api/artists")
	GetData("https://groupietrackers.herokuapp.com/api/artists")

}

func GetData(url string) {

	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)

	ToWrite(string(html))
}

func ToWrite(url string) {
	file, err := os.Create("save.json") // Create a json file

	if err != nil {
		return
	}

	FileData, _ := json.MarshalIndent(url, "", " ")   // Encode the json file
	_ = ioutil.WriteFile("save.json", FileData, 0644) // Write the variable in the json file

	defer file.Close()
}
