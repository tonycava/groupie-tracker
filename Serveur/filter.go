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

func FilterCreation(allband []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	var output []Struct.ArtistPage
	for i := 0; i < len(allband); i++ {
		date, _ := strconv.Atoi(search[1])
		if allband[i].CreationDate > date {
			output = append(output, allband[i])
		}
	}
	if output == nil {
		return allband
	}
	return output
}

func FilterFirstAlbum(allband []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	if search[0] == "" {
		return allband
	}
	var output []Struct.ArtistPage
	for i := 0; i < len(allband); i++ {
		if allband[i].FirstAlbum == search[0] {
			output = append(output, allband[i])
		}
	}
	if output == nil {
		return allband
	}
	return output
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

func FilterByName(allBand []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	if search[0] == "" {
		return allBand
	}
	var allBandFilteredWithLocation []Struct.ArtistPage
	for i := 0; i < len(allBand); i++ {
		if strings.Contains(allBand[i].Name, search[0]) {
			allBandFilteredWithLocation = append(allBandFilteredWithLocation, allBand[i])
		}
	}
	if allBandFilteredWithLocation == nil {
		return allBand
	}
	return allBandFilteredWithLocation
}

func FilterNameOfMember(allband []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	if search[0] == "" {
		return allband
	}
	var output []Struct.ArtistPage
	for i := 0; i < len(allband); i++ {
		for j := 0; j < len(allband[i].Members); j++ {
			if strings.Contains(allband[i].Members[j], search[0]) {
				output = append(output, allband[i])
			}
		}
	}
	if output == nil {
		return allband
	}
	return output
}

func FilterByDate(allband []Struct.ArtistPage, search []string) []Struct.ArtistPage {
	if search[0] == "" {
		return allband
	}
	var output []Struct.ArtistPage
	for i := 0; i < len(allband); i++ {
		for j := 0; j < len(allband[i].Dates); j++ {
			if allband[i].Dates[j] == search[0] {
				output = append(output, allband[i])
			}
		}
	}
	if output == nil {
		return allband
	}
	return output
}

func Filter(search []string) []Struct.ArtistPage {
	allband := GetAllArtistPageData()

	return FilterByName(FilterByDate(FilterLocation(FilterFirstAlbum(FilterNumberOfMember(FilterCreation(FilterNameOfMember(allband, search), search), search), search), search), search), search)
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
