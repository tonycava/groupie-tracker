package Serveur

func contain(tofind string, container []string) bool {
	for i := range container {
		if container[i] == tofind {
			return true
		}
	}
	return false
}

func ToDisplaySearchBar() []string {
	var Container []string

	Myreturn := GetAllArtistPageData()

	for i := range Myreturn {
		for j := range Myreturn[i].Locations {
			if !contain(Myreturn[i].Locations[j], Container) {
				Container = append(Container, Myreturn[i].Locations[j])
			}
		}
	}
	return Container
}
