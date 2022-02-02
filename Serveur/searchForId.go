package Serveur

import (
	"fmt"
	"strconv"
)

func Search(id string) {

	apiUrl := GetAPI()
	result, _ := GetData(apiUrl.Artists, "../json/artist.json", "Artist")

	idInt, _ := strconv.Atoi(id)

	for i := range result {
		if result[i].Id == idInt {
			fmt.Println(result[i])
		}
	}
}
