package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Author struct {
    gorm.Model
    Name string
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Please provide the file name as an argument")
    }
    fileName := os.Args[1]

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.AutoMigrate(&Author{})

    f, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    r := csv.NewReader(f)

    _, err = r.Read()
    if err != nil {
        log.Fatal(err)
    }

    records, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    for _, record := range records {
        author := Author{Name: record[0]}
        db.Create(&author)
    }

    log.Println("Successfully migrated and seeded the database.")
}