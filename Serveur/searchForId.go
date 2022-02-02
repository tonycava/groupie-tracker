package Serveur

import (
	"groupie-tracker/Struct"
	"strconv"
)

func createPage(rawArtistData []Struct.Artist, location, dates []string, relation map[string][]string) []Struct.ArtistPage {
	artistPageToSite := make([]Struct.ArtistPage, 1)
	artistPageToSite[0].Id = rawArtistData[0].Id
	artistPageToSite[0].Image = rawArtistData[0].Image
	artistPageToSite[0].Name = rawArtistData[0].Name
	artistPageToSite[0].Members = rawArtistData[0].Members
	artistPageToSite[0].CreationDate = rawArtistData[0].CreationDate
	artistPageToSite[0].FirstAlbum = rawArtistData[0].FirstAlbum
	artistPageToSite[0].Locations = location
	artistPageToSite[0].Dates = dates
	artistPageToSite[0].DateLocations = relation
	return artistPageToSite
}

func Search(id string) []Struct.ArtistPage {

	apiUrl := GetAPI()

	artistList, artistLocation, artistDates, artistRelations, _ := GetData(apiUrl)

	idInt, _ := strconv.Atoi(id)
	var rawArtistData []Struct.Artist
	var locations, dates []string
	var relations map[string][]string

	for i := range artistList {
		if artistList[i].Id == idInt {
			rawArtistData = []Struct.Artist{artistList[i]}
			locations = artistLocation.Index[i].Locations
			dates = artistDates.Index[i].Date
			relations = artistRelations.Index[i].DatesLocations
			return createPage(rawArtistData, locations, dates, relations)
		}
	}

	return []Struct.ArtistPage{}
}
