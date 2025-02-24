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
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	// Appeler la fonction FetchData
	data, err := models.FetchData(apiURL)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("templates/index.html"))
		tpl.Execute(w, data)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) { // Page secondaire acessible depuis la première page
		tpl := template.Must(template.ParseFiles("templates/artiste.html"))
		tpl.Execute(w, nil)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	fmt.Print("Pour accéder au serveur lancé, aller sur la page http://localhost:8080") // Lancement du serveur sur le port 8080
	http.ListenAndServe(port, nil)                                                      // Utilisée pour démarrer un serveur HTTP (nil car on va utiliser nos propres fonctions)
}
