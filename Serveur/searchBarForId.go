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

func contain(tofind string, container []string) bool {
	for i := range container {
		if container[i] == tofind {
			return true
		}
	}
	return false
}

func ToDisplaySearchBar() {
	var Container []string

	Myreturn := GetAllArtistPageData()

	for i := range Myreturn {
		for j := range Myreturn[i].Locations {
			if !contain(Myreturn[i].Locations[j], Container) {
				Container = append(Container, Myreturn[i].Locations[j])
			}
		}
	}
	fmt.Print(Container)
}
