package infrastructure

import (
	"fmt"

	toy "toy-manager/domain"

	"github.com/jmoiron/sqlx"
)

// PostgresToyRepository es una implementación de la interfaz toy.Repository para PostgreSQL.
type PostgresToyRepository struct {
	DB *sqlx.DB
}

// NewPostgresToyRepository crea una nueva instancia de PostgresToyRepository.
func NewPostgresToyRepository(db *sqlx.DB) *PostgresToyRepository {
	return &PostgresToyRepository{DB: db}
}

func ConnectionOracle() *sqlx.DB {
	var err error

	dsn := "postgres://postgres:password@192.168.1.11/toys_db_development?sslmode=disable"

	driverName := "postgres"
	conn, err := sqlx.Open(driverName, dsn)
	if err != nil {
		return nil
	}
	fmt.Println("Conexión a la base de datos abierta correctamente")
	return conn
}

// SaveToy guarda un nuevo juguete en la base de datos PostgreSQL.
func (r *PostgresToyRepository) SaveToy(toy *toy.Toy) error {
	defer r.close()
	_, err := r.DB.Exec("INSERT INTO toys (name, type, color, size, material, location, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		toy.Name, toy.Type, toy.Properties.Color, toy.Properties.Size, toy.Properties.Material, toy.Location, toy.Status)
	if err != nil {
		return fmt.Errorf("error al guardar el juguete en PostgreSQL: %w", err)
	}
	return nil
}

// Si comento desde aqui abajo, ocurrirá un error
// GetToyByID obtiene un juguete por su ID desde la base de datos PostgreSQL.
func (r *PostgresToyRepository) GetToyByID(id int) (*toy.Toy, error) {
	return nil, fmt.Errorf("método GetToyByID no implementado para PostgreSQL")
}

// UpdateToy actualiza un juguete existente en la base de datos PostgreSQL.
func (r *PostgresToyRepository) UpdateToy(toy *toy.Toy) error {
	return fmt.Errorf("método UpdateToy no implementado para PostgreSQL")
}

// DeleteToy elimina un juguete de la base de datos PostgreSQL.
func (r *PostgresToyRepository) DeleteToy(id int) error {
	return fmt.Errorf("método DeleteToy no implementado para PostgreSQL")
}

// GetAllToys obtiene todos los juguetes desde la base de datos PostgreSQL.
func (r *PostgresToyRepository) GetAllToys() ([]*toy.Toy, error) {
	defer r.close()
	rows, err := r.DB.Query("SELECT id, name, type, color, size, material, location, status FROM toys")
	if err != nil {
		return nil, fmt.Errorf("error al obtener los juguetes desde PostgreSQL: %w", err)
	}
	defer rows.Close()

	var toys []*toy.Toy
	for rows.Next() {
		var t toy.Toy
		if err := rows.Scan(&t.ID, &t.Name, &t.Type, &t.Properties.Color, &t.Properties.Size, &t.Properties.Material, &t.Location, &t.Status); err != nil {
			return nil, fmt.Errorf("error al escanear los juguetes desde PostgreSQL: %w", err)
		}
		toys = append(toys, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre los resultados desde PostgreSQL: %w", err)
	}
	return toys, nil
}

// GetToysByType obtiene juguetes por tipo desde la base de datos PostgreSQL.
func (r *PostgresToyRepository) GetToysByType(toyType toy.ToyType) ([]*toy.Toy, error) {
	return nil, fmt.Errorf("método GetToysByType no implementado para PostgreSQL")
}

func (r *PostgresToyRepository) close() {
	fmt.Println("Cerrando conexión a la base de datos...")
	if err := r.DB.Close(); err != nil {
		fmt.Printf("Error al cerrar la conexión a la base de datos: %v", err)
	} else {
		fmt.Println("Conexión a la base de datos cerrada correctamente")
	}
}
