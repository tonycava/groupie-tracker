package Serveur

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/Struct"
	"io/ioutil"
	"net/http"
	"os"
)

func Main() {
	apiUrl := GetAPI()
	fmt.Print(apiUrl)

	GetData(apiUrl.Artists, "../json/artist.json", "Artist")
	GetData(apiUrl.Locations, "../json/locations.json", "Locations")
	GetData(apiUrl.Dates, "../json/dates.json", "Date")
	GetData(apiUrl.Relation, "../json/relation.json", "Relation")

	GroupieServer()
}

func GetAPI() Struct.API {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	var apiUrl Struct.API
	err = json.Unmarshal(html, &apiUrl)
	if err != nil {
		panic(err)
	}

	return apiUrl
}

func GetData(url, path, usedStruct string) ([]Struct.Artist, []byte) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)

	switch usedStruct {
	case "Artist":
		var apiStruct []Struct.Artist
		err = json.Unmarshal(html, &apiStruct)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(path) // Create a json file
		if err != nil {

		}

		FileData, _ := json.MarshalIndent(apiStruct, "", " ") // Encode the json file
		_ = ioutil.WriteFile(path, FileData, 0644)            // Write the variable in the json file

		return apiStruct, html

		defer file.Close()

	case "Locations":
		var apiStruct Struct.Locations
		err = json.Unmarshal(html, &apiStruct)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path) // Create a json file
		if err != nil {
			println(err)
		}
		FileData, _ := json.MarshalIndent(apiStruct, "", " ") // Encode the json file
		_ = ioutil.WriteFile(path, FileData, 0644)            // Write the variable in the json file
		defer file.Close()

	case "Date":
		var apiStruct Struct.Date
		err = json.Unmarshal(html, &apiStruct)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(path) // Create a json file
		if err != nil {
			println(err)
		}
		FileData, _ := json.MarshalIndent(apiStruct, "", " ") // Encode the json file
		_ = ioutil.WriteFile(path, FileData, 0644)            // Write the variable in the json file
		defer file.Close()

	case "Relation":
		var apiStruct Struct.Relation
		err = json.Unmarshal(html, &apiStruct)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path) // Create a json file
		if err != nil {
			println(err)
		}
		FileData, _ := json.MarshalIndent(apiStruct, "", " ") // Encode the json file
		_ = ioutil.WriteFile(path, FileData, 0644)            // Write the variable in the json file
		defer file.Close()
	}
	return []Struct.Artist{}, []byte{}
}
