package Serveur

import (
	"fmt"
	"groupie-tracker/Struct"
	"html/template"
	"net/http"
	"strconv"
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
	if r.Method == "POST" {
		inputSearchBar := r.FormValue("wizards")
		toSearch := searchBarForId(inputSearchBar)

		s2 := strconv.Itoa(toSearch)
		http.Redirect(w, r, "/artist/"+s2, 303)

	}

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)
	ToDisplaySearchBar()

	tmpl := template.Must(template.ParseFiles("Client/Home.gohtml"))
	_ = tmpl.Execute(w, struct {
		ToDisplay []Struct.Artist
		Location  Struct.Locations
	}{ToDisplay: result, Location: Locations})
}

func Artists(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		inputSearchBar := r.FormValue("wizards")
		inputRange := r.FormValue("range")

		inputMembers1 := r.FormValue("1")
		inputMembers2 := r.FormValue("2")
		inputMembers3 := r.FormValue("3")
		inputMembers4 := r.FormValue("4")
		inputMembers5 := r.FormValue("5")
		inputMembers6 := r.FormValue("6")

		fmt.Println(inputSearchBar)
		fmt.Println(inputRange)

		fmt.Println(inputMembers1)
		fmt.Println(inputMembers2)
		fmt.Println(inputMembers3)
		fmt.Println(inputMembers4)
		fmt.Println(inputMembers5)
		fmt.Println(inputMembers6)

		useformatis := inArray(inputSearchBar, inputRange, inputMembers1, inputMembers2, inputMembers3, inputMembers4, inputMembers5, inputMembers6)
		fmt.Println(useformatis)

	}

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)
	ArrayLocation := ToDisplaySearchBar()

	tmpl := template.Must(template.ParseFiles("Client/Artists.gohtml"))
	_ = tmpl.Execute(w, struct {
		Arralocation []string
		ToDisplay    []Struct.Artist
		Location     Struct.Locations
	}{ToDisplay: result, Location: Locations, Arralocation: ArrayLocation})
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

	if r.Method == "POST" {
		inputSearchBar := r.FormValue("wizards")
		toSearch := searchBarForId(inputSearchBar)

		s2 := strconv.Itoa(toSearch)
		http.Redirect(w, r, "/artist/"+s2, 303)
		fmt.Println(toSearch)
	}

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)

	ResultId := Search(request)
	//fmt.Println(ResultId)
	tmpl := template.Must(template.ParseFiles("Client/ArtistsId.gohtml"))
	_ = tmpl.Execute(w, struct {
		Data      []Struct.ArtistPage
		ToDisplay []Struct.Artist
		Location  Struct.Locations
	}{Data: ResultId, ToDisplay: result, Location: Locations})

}

func Credit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		inputSearchBar := r.FormValue("wizards")
		toSearch := searchBarForId(inputSearchBar)

		s2 := strconv.Itoa(toSearch)
		http.Redirect(w, r, "/artist/"+s2, 303)
		fmt.Println(toSearch)
	}

	apiUrl := GetAPI()
	result, Locations, _, _, _ := GetData(apiUrl)

	tmpl := template.Must(template.ParseFiles("Client/Credit.gohtml"))
	_ = tmpl.Execute(w, struct {
		ToDisplay []Struct.Artist
		Location  Struct.Locations
	}{ToDisplay: result, Location: Locations})

}

func FavIcon(w http.ResponseWriter, r *http.Request) { //get icon for the site
	http.ServeFile(w, r, "favicon/favicon.ico")
}
