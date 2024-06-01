package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"toy-manager/application"
	"toy-manager/domain"
	"toy-manager/infrastructure"

	_ "github.com/lib/pq"
)

func main() {
	// Enrutador HTTP
	http.HandleFunc("/toys", getAllToysHandler())
	http.HandleFunc("/toys/add", addToyHandler())

	// Iniciar el servidor HTTP
	log.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Controlador para obtener todos los juguetes
func getAllToysHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//infra
		db := infrastructure.ConnectionOracle()
		toyRepo := infrastructure.NewPostgresToyRepository(db)
		//instancia aplicacion
		toyService := application.NewToyService(toyRepo)
		// Llamar al servicio para obtener todos los juguetes
		toys, err := toyService.GetAllToys()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error al obtener los juguetes: %v", err), http.StatusInternalServerError)
			return
		}
		// Escribir la respuesta HTTP con los juguetes en formato JSON
		// (Aquí deberías usar un paquete como encoding/json para serializar los juguetes a JSON)
		w.Header().Set("Content-Type", "application/json")
		// Escribir la respuesta JSON en el cuerpo de la respuesta HTTP
		// (Aquí deberías usar un paquete como encoding/json para serializar los juguetes a JSON)
		// Por ejemplo:
		json.NewEncoder(w).Encode(toys)
	}
}

// Controlador para agregar un juguete
func addToyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//infra
		db := infrastructure.ConnectionOracle()
		toyRepo := infrastructure.NewPostgresToyRepository(db)
		//instancia aplicacion
		toyService := application.NewToyService(toyRepo)

		var newToy domain.Toy // Utiliza el tipo Toy del paquete de dominio
		err := json.NewDecoder(r.Body).Decode(&newToy)
		if err != nil {
			http.Error(w, "Error al decodificar los datos del juguete", http.StatusBadRequest)
			return
		}

		err = toyService.AddToy(newToy)
		if err != nil {
			http.Error(w, "Error al agregar el juguete", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}
