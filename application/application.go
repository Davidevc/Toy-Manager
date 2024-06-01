package application

import (
	"fmt"

	toy "toy-manager/domain"
)

// ToyService es un servicio que contiene la lógica de aplicación relacionada con los juguetes.
type ToyService struct {
	toyRepo toy.Repository
}

// NewToyService crea una nueva instancia de ToyService con el repositorio dado.
func NewToyService(repo toy.Repository) *ToyService {
	return &ToyService{toyRepo: repo}
}

// AddToy agrega un nuevo juguete al sistema. Tambien podria ir varios parametros
func (s *ToyService) AddToy(newToy toy.Toy) error {
	// Guardar el nuevo juguete utilizando el repositorio.
	if err := s.toyRepo.SaveToy(&newToy); err != nil {
		return fmt.Errorf("error al agregar el juguete: %w", err)
	}

	return nil
}

// GetAllToys devuelve todos los juguetes.
func (s *ToyService) GetAllToys() ([]*toy.Toy, error) {
	return s.toyRepo.GetAllToys()
}
