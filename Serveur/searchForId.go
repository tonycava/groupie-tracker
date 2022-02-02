package Serveur

import (
	"groupie-tracker/Struct"
	"strconv"
)

func createPage(rawArtistData []Struct.Artist, location, dates []string, relation map[string][]string) Struct.ArtistPage {
	var artistPageToSite Struct.ArtistPage
	artistPageToSite.Id = rawArtistData[0].Id
	artistPageToSite.Image = rawArtistData[0].Image
	artistPageToSite.Name = rawArtistData[0].Name
	artistPageToSite.Members = rawArtistData[0].Members
	artistPageToSite.CreationDate = rawArtistData[0].CreationDate
	artistPageToSite.FirstAlbum = rawArtistData[0].FirstAlbum
	artistPageToSite.Locations = location
	artistPageToSite.Dates = dates
	artistPageToSite.DateLocations = relation
	return artistPageToSite
}

func Search(id string) Struct.ArtistPage {

	apiUrl := GetAPI()

	artistList, _, _, _, _ := GetData(apiUrl.Artists, "", "Artist")
	_, artistLocation, _, _, _ := GetData(apiUrl.Locations, "", "Locations")
	_, _, artistDates, _, _ := GetData(apiUrl.Dates, "", "Date")
	_, _, _, artistRelations, _ := GetData(apiUrl.Relation, "", "Relation")

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

	return Struct.ArtistPage{}
}
