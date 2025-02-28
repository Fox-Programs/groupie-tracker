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
	// URL des APIs
	apiURLIndex := "https://groupietrackers.herokuapp.com/api/artists"
	apiURLLocations := "https://groupietrackers.herokuapp.com/api/locations"

	dataIndex, err := models.FetchDataIndex(apiURLIndex)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	// Récupération des données
	dataLocationsDates, err := models.GetLocationsWithDates(apiURLLocations)
	if err != nil {
		log.Fatal("Erreur récupération des données:", err)
	}

	// Fonctions qui permttent d'envoyer les données de l'API récupéré vers les différentes pages html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("templates/index.html"))
		tpl.Execute(w, dataIndex)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) { // Page secondaire acessible depuis la première page
		tpl := template.Must(template.ParseFiles("templates/artiste.html"))
		tpl.Execute(w, dataIndex)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) { // Page qui devait être faite par Lucas
		tpl := template.Must(template.ParseFiles("templates/location.html"))
		tpl.Execute(w, dataLocationsDates)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	fmt.Print("Pour accéder au serveur lancé, aller sur la page http://localhost:8080") // Lancement du serveur sur le port 8080
	http.ListenAndServe(port, nil)                                                      // Utilisée pour démarrer un serveur HTTP (nil car on va utiliser nos propres fonctions)
}
