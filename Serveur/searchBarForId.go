package Serveur

import (
	"fmt"
)

func searchBarForId(ToSearch string) int {
	var toRedirect int
	apiUrl := GetAPI()
	artistList, _, _, _, _ := GetData(apiUrl)
	for i := range artistList {
		for j := range artistList[i].Members {
			if artistList[i].Members[j] == ToSearch || artistList[i].FirstAlbum == ToSearch {
				toRedirect = artistList[i].Id
			}
		}
	}
	return toRedirect
}

func ToDisplaySearchBar() {
	//var Container []string

	Myreturn := GetAllArtistPageData()

	for i := range Myreturn {
		fmt.Println(Myreturn[i].Locations)

	}
}
