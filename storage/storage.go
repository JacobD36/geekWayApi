package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver es el driver de la base de datos
type Driver string

// Drivers
const (
	MySQL    Driver = "mysql"
	Postgres Driver = "postgres"
)

// New crea una nueva conexión con la base de datos
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// Inicia una nueva conexión a la Base de datos Postgres
func newPostgresDB() {
	once.Do(func() {
		//Toda esta sección se ejecutará una sola vez - patrón Singleton
		var err error
		db, err = gorm.Open(postgres.Open("postgres://jaime:kbjnfqfsfy79@localhost:5432/innomedicerp?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Fatalf("No se pudo realizar la conexión con base de datos: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

// Inicia una nueva conexión a la base de datos MySQL
func newMySQLDB() {
	once.Do(func() {
		//Toda esta sección se ejecutará una sola vez - patrón Singleton
		var err error
		db, err = gorm.Open(mysql.Open("bay:bayental2019@tcp(localhost:3306)/innomedicerp?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Fatalf("No se pudo realizar la conexión con la base de datos: %v", err)
		}
	})

	fmt.Println("Conectado a MySQL")
}

// DB retorna una instancia única a la base de datos
func DB() *gorm.DB {
	return db
}
