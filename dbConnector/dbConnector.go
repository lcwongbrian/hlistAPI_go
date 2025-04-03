package dbConnector

import (
	"context"
	"database/sql"
	"encoding/binary"

	"fmt"
	"math"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBHandler struct {
	DB *sql.DB
}

func NewDBHandler() (*DBHandler, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require",
		os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DBHandler{DB: db}, nil
}

func (dh *DBHandler) Initialize(ctx context.Context) error {
	_, err := dh.DB.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS hlist (surface_id SERIAL PRIMARY KEY, vertices BYTEA);`)
	if err != nil {
		return err
	}

	fmt.Println("[PG] Successfully access DB")
	return nil
}

func (dh *DBHandler) FetchSurfaceBinaryByID(ctx context.Context, id int32) ([]float32, error) {
	var vertices []byte
	err := dh.DB.QueryRowContext(ctx, "SELECT vertices FROM hlist WHERE surface_id = $1", id).Scan(&vertices)
	if err != nil {
		return nil, err
	}

	result := make([]float32, len(vertices)/4)
	for i := 0; i < len(vertices); i += 4 {
		bits := binary.LittleEndian.Uint32(vertices[i : i+4])
		result[i/4] = math.Float32frombits(bits)
	}

	return result, nil
}

func (dh *DBHandler) Close() error {
	return dh.DB.Close()
}
