package Serveur

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func GroupieServer() {

	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	apiUrl := GetAPI()
	result, html := GetData(apiUrl.Artists, "../json/artist.json", "Artist")
	var data []string
	var dataforint []int

	var apiStruct []Artist
	err := json.Unmarshal(html, &apiStruct)
	if err != nil {
		panic(err)
	}

	for i := range apiStruct {
		data = append(data, result[i].Image)
	}

	for i := range apiStruct {
		dataforint = append(dataforint, result[i].Id)
	}

	tmpl := template.Must(template.ParseFiles("Client/Mainindex.gohtml"))
	_ = tmpl.Execute(w, struct {
		Data   []string
		DataId []int
	}{Data: data, DataId: dataforint})

	fmt.Println(result)
	fmt.Println(result[1].Image)

}
