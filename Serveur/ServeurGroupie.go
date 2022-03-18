package Serveur

import (
	"fmt"
	"groupie-tracker/Struct"
	"html/template"
	"net/http"
	"strings"
)

func GroupieServer() {

	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/credit", Credit)
	server.HandleFunc("/artist", Artists)
	server.HandleFunc("/artist/", ArtistsId)

	server.HandleFunc("/favicon.ico", FavIcon)

	// or use strings.Split, or use regexp, etc.

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	server.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("./script"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8080/artist")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)
	ArrayLocation := ToDisplaySearchBar()

	tmpl := template.Must(template.ParseFiles("Client/Home.gohtml"))
	_ = tmpl.Execute(w, struct {
		InSearch     []Struct.Artist
		ToDisplay    []Struct.Artist
		Arralocation []string
		Location     Struct.Locations
	}{ToDisplay: result, Location: Locations, InSearch: result, Arralocation: ArrayLocation})
}

func Artists(w http.ResponseWriter, r *http.Request) {
	display := true
	apiUrl := GetAPI()
	result, _, _, _, _ := GetData(apiUrl)

	if r.Method == "POST" {
		display = false
		inputSearchBar := r.FormValue("wizards")
		inputRange := r.FormValue("range")

		inputMembers1 := r.FormValue("1")
		inputMembers2 := r.FormValue("2")
		inputMembers3 := r.FormValue("3")
		inputMembers4 := r.FormValue("4")
		inputMembers5 := r.FormValue("5")
		inputMembers6 := r.FormValue("6")

		inArray(inputSearchBar, inputRange, inputMembers1, inputMembers2, inputMembers3, inputMembers4, inputMembers5, inputMembers6)

		allband := Filter(inArray(inputSearchBar, inputRange, inputMembers1, inputMembers2, inputMembers3, inputMembers4, inputMembers5, inputMembers6))
		ArrayLocation := ToDisplaySearchBar()

		tmpl := template.Must(template.ParseFiles("Client/Artists.gohtml"))
		_ = tmpl.Execute(w, struct {
			InSearch     []Struct.Artist
			Arralocation []string
			ToDisplay    []Struct.ArtistPage
		}{ToDisplay: allband, Arralocation: ArrayLocation, InSearch: result})
	}

	if display == true {

		ArrayLocation := ToDisplaySearchBar()

		result2 := GetAllArtistPageData()

		tmpl := template.Must(template.ParseFiles("Client/Artists.gohtml"))
		_ = tmpl.Execute(w, struct {
			InSearch     []Struct.Artist
			Arralocation []string
			ToDisplay    []Struct.ArtistPage
		}{ToDisplay: result2, Arralocation: ArrayLocation, InSearch: result})
	}
	display = true
}

func ArtistsId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	request := id[len("/artist/"):]

	if strings.Contains(request, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("ID Out of Range"))
		if err != nil {
			return
		}
		return
	}

	ArrayLocation := ToDisplaySearchBar()

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)

	ResultId := Search(request)

	tmpl := template.Must(template.ParseFiles("Client/ArtistsId.gohtml"))
	_ = tmpl.Execute(w, struct {
		InSearch     []Struct.Artist
		Arralocation []string
		Data         []Struct.ArtistPage
		ToDisplay    []Struct.Artist
		Location     Struct.Locations
	}{Data: ResultId, ToDisplay: result, Location: Locations, Arralocation: ArrayLocation, InSearch: result})
}

func Credit(w http.ResponseWriter, r *http.Request) {

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)
	ArrayLocation := ToDisplaySearchBar()

	tmpl := template.Must(template.ParseFiles("Client/Credit.gohtml"))
	_ = tmpl.Execute(w, struct {
		InSearch     []Struct.Artist
		Arralocation []string
		ToDisplay    []Struct.Artist
		Location     Struct.Locations
	}{ToDisplay: result, Location: Locations, Arralocation: ArrayLocation, InSearch: result})

}

func FavIcon(w http.ResponseWriter, r *http.Request) { //get icon for the site
	http.ServeFile(w, r, "favicon/favicon.ico")
}
