package Serveur

import (
	"groupie-tracker/Struct"
	"strconv"
)

func GetAllArtistPageData() []Struct.ArtistPage {
	a, l, d, r, _ := GetData(GetAPI())
	allArtistPage := []Struct.ArtistPage{}
	for i := range a {
		rawArtistData := []Struct.Artist{a[i]}
		locations := l.Index[i].Locations
		dates := d.Index[i].Date
		relations := r.Index[i].DatesLocations
		allArtistPage = append(allArtistPage, createPage(rawArtistData, locations, dates, relations)[0])
	}
	return allArtistPage
}

func FilterCreation(allBand []Struct.ArtistPage) []Struct.ArtistPage {
	loop := true
	for loop {
		loop = false
		for i := 1; i < len(allBand); i++ {
			if allBand[i-1].CreationDate > allBand[i].CreationDate {
				allBand[i-1], allBand[i] = allBand[i], allBand[i-1]
				loop = true
			}
		}
	}
	return allBand
}

func FilterFirstAlbum(allBand []Struct.ArtistPage) []Struct.ArtistPage {
	var arrayDateInStandardOrder []string
	for i := 1; i < len(allBand); i++ {
		dateInStandardOrder := allBand[i].FirstAlbum[6:] + allBand[i].FirstAlbum[3:5] + allBand[i].FirstAlbum[0:2]
		arrayDateInStandardOrder = append(arrayDateInStandardOrder, dateInStandardOrder)
	}
	loop := true
	for loop {
		loop = false
		for i := 1; i < len(arrayDateInStandardOrder); i++ {
			a, _ := strconv.Atoi(arrayDateInStandardOrder[i-1])
			b, _ := strconv.Atoi(arrayDateInStandardOrder[i])
			if a > b {
				allBand[i-1], allBand[i] = allBand[i], allBand[i-1]
				arrayDateInStandardOrder[i-1], arrayDateInStandardOrder[i] = arrayDateInStandardOrder[i], arrayDateInStandardOrder[i-1]
				loop = true
			}
		}
	}
	return allBand
}

func FilterNumberOfMember(allBand []Struct.ArtistPage) []Struct.ArtistPage {
	loop := true
	for loop {
		loop = false
		for i := 1; i < len(allBand); i++ {
			if len(allBand[i-1].Members) > len(allBand[i].Members) {
				allBand[i-1], allBand[i] = allBand[i], allBand[i-1]
				loop = true
			}
		}
	}
	return allBand
}

func FilterLocation(allBand []Struct.ArtistPage, location string) []Struct.ArtistPage {
	var allBandFilteredWithLocation []Struct.ArtistPage
	for i := 0; i < len(allBand); i++ {
		for j := 0; j < len(allBand[i].Locations); j++ {
			if allBand[i].Locations[j] == location {
				allBandFilteredWithLocation = append(allBandFilteredWithLocation, allBand[i])
			}
		}
	}
	return allBandFilteredWithLocation
}
