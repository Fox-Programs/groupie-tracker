package main

import (
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
)

const port = ":8080" // port mis dans une constante pour le réemployer par la suite

func main() {
	// URL de l'API
	apiURL_index := "https://groupietrackers.herokuapp.com/api/artists"
	apiURL_locations := "https://groupietrackers.herokuapp.com/api/locations"
	apiURL_dates := "https://groupietrackers.herokuapp.com/api/dates"

	// Appeler la fonction FetchData_index
	data_index, err := models.FetchData_index(apiURL_index)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	// Appeler la fonction FetchData_locations_dates
	data_locations, err := models.FetchData_index(apiURL_locations)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
	data_dates, err := models.FetchData_index(apiURL_dates)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
	data_locations_dates := append(data_locations, data_dates...)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("templates/index.html"))
		tpl.Execute(w, data_index)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) { // Page secondaire acessible depuis la première page
		tpl := template.Must(template.ParseFiles("templates/artiste.html"))
		tpl.Execute(w, nil)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) { // Page secondaire acessible depuis la première page
		tpl := template.Must(template.ParseFiles("templates/location.html"))
		tpl.Execute(w, data_locations_dates)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	fmt.Print("Pour accéder au serveur lancé, aller sur la page http://localhost:8080") // Lancement du serveur sur le port 8080
	http.ListenAndServe(port, nil)                                                      // Utilisée pour démarrer un serveur HTTP (nil car on va utiliser nos propres fonctions)
}
