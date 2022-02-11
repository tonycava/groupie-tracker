package Serveur

import (
	"groupie-tracker/Struct"
	"strconv"
)

func createPage(rawArtistData []Struct.Artist, location, dates []string, relation map[string][]string) []Struct.ArtistPage {
	return []Struct.ArtistPage{
		{
			Id:            rawArtistData[0].Id,
			Image:         rawArtistData[0].Image,
			Name:          rawArtistData[0].Name,
			Members:       rawArtistData[0].Members,
			CreationDate:  rawArtistData[0].CreationDate,
			FirstAlbum:    rawArtistData[0].FirstAlbum,
			Locations:     location,
			Dates:         dates,
			DateLocations: relation,
		},
	}
}

func Search(id string) []Struct.ArtistPage {

	artistList, artistLocation, artistDates, artistRelations, _ := GetData(GetAPI())
	idInt, _ := strconv.Atoi(id)

	for i := range artistList {
		if artistList[i].Id == idInt {
			rawArtistData := []Struct.Artist{artistList[i]}
			locations := artistLocation.Index[i].Locations
			dates := artistDates.Index[i].Date
			relations := artistRelations.Index[i].DatesLocations
			return createPage(rawArtistData, locations, dates, relations)
		}
	}
	return []Struct.ArtistPage{}
}
