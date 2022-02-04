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
	server.HandleFunc("/artist", Artists)
	server.HandleFunc("/artist/", ArtistsId)

	// or use strings.Split, or use regexp, etc.

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8080/artist")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

}

func Artists(w http.ResponseWriter, r *http.Request) {
	apiUrl := GetAPI()
	result, _, _, _, _ := GetData(apiUrl)

	tmpl := template.Must(template.ParseFiles("Client/Mainindex.gohtml"))
	_ = tmpl.Execute(w, result)

}

func ArtistsId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	fmt.Println(id)
	request := id[len("/artist/"):]
	fmt.Println(request)

	if strings.Contains(request, "/") {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("ID Out of Range"))
		if err != nil {
			return
		}
		return
	}

	ResultId := Search(request)
	fmt.Println(ResultId)
	tmpl := template.Must(template.ParseFiles("Client/ArtistsId.gohtml"))
	_ = tmpl.Execute(w, struct {
		Data []Struct.ArtistPage
	}{Data: ResultId})

}
