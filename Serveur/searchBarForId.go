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
	var Container []string

	Myreturn := GetAllArtistPageData()

	for i := range Myreturn {
		for j := range Myreturn[i].Locations {
			fmt.Println(Myreturn[i].Locations[j])

			for l := range Container {
				if Container[l] != Myreturn[i].Locations[j] {
					Container = append(Container, Myreturn[i].Locations[j])
				}
			}
		}
	}
	fmt.Println(Container)
}
