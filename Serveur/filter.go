package Serveur

import (
	"groupie-tracker/Struct"
	"strconv"
	"strings"
)

type SearchInfo struct {
	Search   string `json:"search"`
	RangeMin int    `json:"int_min"`
	RangeMax int    `json:"int_max"`
	Number   int    `json:"flex"`
}

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

func FilterCreation(allBand []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	var output []Struct.ArtistPage
	for i := 0; i < len(allBand); i++ {
		date, _ := strconv.Atoi(search[1])
		if allBand[i].CreationDate > date {
			output = append(output, allBand[i])
		}
	}
	return output
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

func FilterNumberOfMember(allBand []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	var N1, N2, N3, N4, N5, N6 []Struct.ArtistPage
	if search[2] != "" {
		for i := 0; i < len(allBand); i++ {
			if len(allBand[i].Members) == 1 {
				N1 = append(N1, allBand[i])
			}
		}
	}
	if search[3] != "" {
		for i := 0; i < len(allBand); i++ {
			if len(allBand[i].Members) == 2 {
				N2 = append(N2, allBand[i])
			}
		}
	}
	if search[4] != "" {
		for i := 0; i < len(allBand); i++ {
			if len(allBand[i].Members) == 3 {
				N3 = append(N3, allBand[i])
			}
		}
	}
	if search[5] != "" {
		for i := 0; i < len(allBand); i++ {
			if len(allBand[i].Members) == 4 {
				N4 = append(N4, allBand[i])
			}
		}
	}
	if search[6] != "" {
		for i := 0; i < len(allBand); i++ {
			if len(allBand[i].Members) == 5 {
				N5 = append(N5, allBand[i])
			}
		}
	}
	if search[7] != "" {
		for i := 0; i < len(allBand); i++ {
			if len(allBand[i].Members) == 6 {
				N6 = append(N6, allBand[i])
			}
		}
	}
	N1 = append(N1, N2...)
	N1 = append(N1, N3...)
	N1 = append(N1, N4...)
	N1 = append(N1, N5...)
	N1 = append(N1, N6...)
	if N1 == nil {
		return allBand
	}
	return N1
}

func FilterLocation(allBand []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	if search[0] == "" {
		return allBand
	}
	var allBandFilteredWithLocation []Struct.ArtistPage
	for i := 0; i < len(allBand); i++ {
		for j := 0; j < len(allBand[i].Locations); j++ {
			if strings.Contains(allBand[i].Locations[j], search[0]) {
				allBandFilteredWithLocation = append(allBandFilteredWithLocation, allBand[i])
			}
		}
	}
	if allBandFilteredWithLocation == nil {
		return allBand
	}
	return allBandFilteredWithLocation
}

func Filter(search []string) []Struct.ArtistPage {
	allband := GetAllArtistPageData()

	//creation date
	allband1 := FilterCreation(allband, search)
	//N member
	allband2 := FilterNumberOfMember(allband1, search)
	//first album
	allband3 := FilterFirstAlbum(allband2)
	//location
	allband4 := FilterLocation(allband3, search)

	return allband4
}

func inArray(inputSearchbar string, inputRange string, inputMembers1 string, inputMembers2 string, inputMembers3 string, inputMembers4 string, inputMembers5 string, inputMembers6 string) []string {

	return []string{
		inputSearchbar,
		inputRange,
		inputMembers1,
		inputMembers2,
		inputMembers3,
		inputMembers4,
		inputMembers5,
		inputMembers6,
	}
}
