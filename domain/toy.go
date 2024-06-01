package domain

// Package toy contiene el modelo de dominio relacionado con los juguetes.

// ToyType representa el tipo de juguete.
type ToyType string

// Propiedades del juguete.
type ToyProperties struct {
	Color    string // Color del juguete.
	Size     string // Tamaño del juguete.
	Material string // Material del juguete.
}

// Toy representa un juguete en el sistema.
type Toy struct {
	ID         int           // ID único del juguete.
	Name       string        // Nombre del juguete.
	Type       ToyType       // Tipo de juguete.
	Properties ToyProperties // Propiedades del juguete.
	Location   string        // Ubicación actual del juguete.
	Status     string        // Estado del juguete (nuevo, usado, roto, etc.).
}

// Repository define las operaciones que se pueden realizar sobre la colección de juguetes.
type Repository interface {
	GetToyByID(id int) (*Toy, error)               // Obtener un juguete por su ID.
	SaveToy(toy *Toy) error                        // Guardar un nuevo juguete.
	UpdateToy(toy *Toy) error                      // Actualizar la información de un juguete existente.
	DeleteToy(id int) error                        // Eliminar un juguete.
	GetAllToys() ([]*Toy, error)                   // Obtener todos los juguetes.
	GetToysByType(toyType ToyType) ([]*Toy, error) // Obtener juguetes por tipo.
}
