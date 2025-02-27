package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DataIndex struct {
	ID           int      `json:"id"`           // Rajout d'un id par artist pour ne pas avoir d'erreur dans la console de la page html
	Image        string   `json:"image"`        // Lien de l'image
	Name         string   `json:"name"`         // Nom de l'artiste
	CreationDate int      `json:"creationDate"` // Année de création
	FirstAlbum   string   `json:"firstAlbum"`   // Premier album
	Members      []string `json:"members"`      // Liste des membres
}

type APIResponseLocations struct {
	Index []DataLocations `json:"index"`
}

type DataLocations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"` // URL vers l'API des dates
	Dates     []string `json:"-"`     // Stocke les dates après récupération
}

type APIResponseDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func FetchDataIndex(apiURL string) ([]DataIndex, error) {
	// Faire une requête HTTP GET
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'appel à l'API : %v", err)
	}
	defer response.Body.Close()

	// Désérialiser la réponse JSON
	var data_index []DataIndex
	err = json.NewDecoder(response.Body).Decode(&data_index)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON : %v", err)
	}

	return data_index, nil
}

// Fonction générique pour récupérer du JSON depuis une URL
func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

// Fonction pour récupérer les localisations et leurs dates associées
func GetLocationsWithDates(locationsAPI string) ([]DataLocations, error) {
	var locationsResponse APIResponseLocations

	// Récupération des localisations
	if err := fetchJSON(locationsAPI, &locationsResponse); err != nil {
		return nil, err
	}

	// Boucle pour récupérer les dates associées à chaque localisation
	for i, location := range locationsResponse.Index {
		var datesResponse APIResponseDates

		// Récupération des dates via l'URL fournie
		if err := fetchJSON(location.DatesURL, &datesResponse); err != nil {
			return nil, err
		}

		// Stocker les dates dans la structure
		locationsResponse.Index[i].Dates = datesResponse.Dates
	}

	return locationsResponse.Index, nil
}
