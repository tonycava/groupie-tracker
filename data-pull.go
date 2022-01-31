package main

import (
	"encoding/json"
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

type Locations struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Date struct {
	Index []struct {
		Id   int      `json:"id"`
		Date []string `json:"dates"`
	}
}

type Relation struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func main() {
	apiUrl := GetAPI()
	GetData(apiUrl.Artists, "json/artist.json", "Artist")
	GetData(apiUrl.Locations, "json/locations.json", "Locations")
	GetData(apiUrl.Dates, "json/dates.json", "Date")
	GetData(apiUrl.Relation, "json/relation.json", "Relation")
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

func GetData(url, path, usedStruct string) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)

	switch usedStruct {
	case "Artist":
		var APITest []Artist
		err = json.Unmarshal(html, &APITest)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path)
		if err != nil {
			return
		}
		FileData, _ := json.MarshalIndent(APITest, "", " ")
		_ = ioutil.WriteFile(path, FileData, 0644)
		defer file.Close()

	case "Locations":
		var APITest Locations
		err = json.Unmarshal(html, &APITest)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path) // Create a json file
		if err != nil {
			return
		}
		FileData, _ := json.MarshalIndent(APITest, "", " ") // Encode the json file
		_ = ioutil.WriteFile(path, FileData, 0644)          // Write the variable in the json file
		defer file.Close()

	case "Date":
		var APITest Date
		err = json.Unmarshal(html, &APITest)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path)
		if err != nil {
			return
		}
		FileData, _ := json.MarshalIndent(APITest, "", " ")
		_ = ioutil.WriteFile(path, FileData, 0644)
		defer file.Close()

	case "Relation":
		var APITest Relation
		err = json.Unmarshal(html, &APITest)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path)
		if err != nil {
			return
		}
		FileData, _ := json.MarshalIndent(APITest, "", " ")
		_ = ioutil.WriteFile(path, FileData, 0644)
		defer file.Close()
	}
}
