package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Artist représente la structure JSON
type Artist struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// ResponseWrapper encapsule la réponse HTTP et offre des méthodes pratiques
type ResponseWrapper struct {
	Body []byte
}

// GetBody extrait le corps de la réponse HTTP
func (r *ResponseWrapper) GetBody() []byte {
	return r.Body
}

// Unmarshal convertit le corps JSON en une structure Go
func (r *ResponseWrapper) Unmarshal(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// FetchURL fait une requête GET et retourne un ResponseWrapper
func FetchURL(url string) (*ResponseWrapper, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête : %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du corps : %w", err)
	}

	return &ResponseWrapper{Body: body}, nil
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists" // Récupération de l'URL pour les infos des artistes

	// Faire une requête et obtenir la réponse encapsulée
	response, err := FetchURL(url)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	// Désérialiser en une liste d'artistes
	var artists []Artist
	err = response.Unmarshal(&artists)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON :", err)
		return
	}

	// Afficher les artistes
	for _, artist := range artists {
		fmt.Printf("Nom : %s, ID : %d, Babebibobu : %s\n", artist.Name, artist.ID)
	}
}
