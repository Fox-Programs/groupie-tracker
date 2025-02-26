package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Données qui se retrouveront dans la page principale
type Data_index struct {
	ID           int      `json:"id"`           // Rajout d'un id par artist pour ne pas avoir d'erreur dans la console de la page html
	Image        string   `json:"image"`        // Lien de l'image
	Name         string   `json:"name"`         // Nom de l'artiste
	CreationDate int      `json:"creationDate"` // Année de création
	FirstAlbum   string   `json:"firstAlbum"`   // Premier album
	Members      []string `json:"members"`      // Liste des membres
}

// Données qui se retrouveront dans la page principale
type Data_locations_dates struct {
	ID        int      `json:"id"`        // Rajout d'un id par artist pour ne pas avoir d'erreur dans la console de la page html
	Locations []string `json:"locations"` // Liste des loaclisations
	Dates     []string `json:"dates"`     // Liste des dates
}

func FetchData_index(apiURL string) ([]Data_index, error) {
	// Faire une requête HTTP GET
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'appel à l'API : %v", err)
	}
	defer response.Body.Close()

	// Vérifier le code HTTP
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur HTTP : %v", response.StatusCode)
	}

	// Désérialiser la réponse JSON
	var data_index []Data_index
	err = json.NewDecoder(response.Body).Decode(&data_index)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON : %v", err)
	}

	return data_index, nil
}

func FetchData_locations_dates(apiURL string) ([]Data_locations_dates, error) {
	// Faire une requête HTTP GET
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'appel à l'API : %v", err)
	}
	defer response.Body.Close()

	// Vérifier le code HTTP
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur HTTP : %v", response.StatusCode)
	}

	// Désérialiser la réponse JSON
	var data_locations_dates []Data_locations_dates
	err = json.NewDecoder(response.Body).Decode(&data_locations_dates)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON : %v", err)
	}

	return data_locations_dates, nil
}
