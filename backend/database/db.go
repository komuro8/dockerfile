package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    dbHost := os.Getenv("POSTGRES_HOST")
    dbPort := os.Getenv("POSTGRES_PORT")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPass := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")

    psqlInfo := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPass, dbName,
    )

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal("Error conectando a la base de datos:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("No se pudo hacer ping a la base de datos:", err)
    }

    log.Println("Conectado a la base de datos Postgres")
}
