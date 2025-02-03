package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Données principales des artistes
type Data struct {
	ID           int      `json:"id"`           // Rajout d'un id par artist pour ne pas avoir d'erreur dans la console de la page html
	Image        string   `json:"image"`        // Lien de l'image
	Name         string   `json:"name"`         // Nom de l'artiste
	CreationDate int      `json:"creationDate"` // Année de création
	FirstAlbum   string   `json:"firstAlbum"`   // Premier album
	Members      []string `json:"members"`      // Liste des membres
}

func FetchData(apiURL string) ([]Data, error) {
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
	var data []Data
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON : %v", err)
	}

	return data, nil
}
