package Serveur

import (
	"groupie-tracker/Struct"
	"strconv"
)

func Search(id string) []Struct.Artist {

	apiUrl := GetAPI()
	result, _ := GetData(apiUrl.Artists, "../json/artist.json", "Artist")

	idInt, _ := strconv.Atoi(id)

	for i := range result {
		if result[i].Id == idInt {
			return []Struct.Artist{result[i]}
		}
	}
	return []Struct.Artist{}
}
