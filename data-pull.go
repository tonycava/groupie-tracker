package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

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

	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)

	var APITest []Artist
	err = json.Unmarshal(html, &APITest)
	fmt.Println(APITest)

	if err != nil {
		panic(err)
	}

	file, err := os.Create("save.json") // Create a json file

	if err != nil {
		return
	}

	FileData, _ := json.MarshalIndent(APITest, "", " ") // Encode the json file
	_ = ioutil.WriteFile("save.json", FileData, 0644)   // Write the variable in the json file

	defer file.Close()
}
