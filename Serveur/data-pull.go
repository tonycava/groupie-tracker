package Serveur

import (
	"encoding/json"
	"groupie-tracker/Struct"
	"io/ioutil"
	"net/http"
)

// toto

func Main() {
	apiUrl := GetAPI()
	GetData(apiUrl)
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

func getUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	return html
}

func GetData(url Struct.API) ([]Struct.Artist, Struct.Locations, Struct.Date, Struct.Relation, []byte) {
	var html []byte
	var err error

	html = getUrl(url.Artists)
	var apiStruct1 []Struct.Artist
	err = json.Unmarshal(html, &apiStruct1)
	if err != nil {
		panic(err)
	}

	html = getUrl(url.Locations)
	var apiStruct2 Struct.Locations
	err = json.Unmarshal(html, &apiStruct2)
	if err != nil {
		panic(err)
	}

	html = getUrl(url.Dates)
	var apiStruct3 Struct.Date
	err = json.Unmarshal(html, &apiStruct3)
	if err != nil {
		panic(err)
	}

	html = getUrl(url.Relation)
	var apiStruct4 Struct.Relation
	err = json.Unmarshal(html, &apiStruct4)
	if err != nil {
		panic(err)
	}

	return apiStruct1, apiStruct2, apiStruct3, apiStruct4, html
}
