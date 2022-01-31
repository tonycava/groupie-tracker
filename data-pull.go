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
	var APIconfig API

	GetAPI()

	jsonFile, err := os.Open("API.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	JsonReader := json.NewDecoder(jsonFile) // Decode the json file
	err = JsonReader.Decode(&APIconfig)

	/*GetAPI(APIconfig.Artists)
	GetAPI(APIconfig.Locations)
	GetAPI(APIconfig.Dates)
	GetAPI(APIconfig.Locations)*/
}

func GetAPI() API {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	var apiUrl API
	err = json.Unmarshal(html, &apiUrl)
	if err != nil {
		panic(err)
	}
	return apiUrl
}
