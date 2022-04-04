package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"post/database"
	"post/models"
	"strconv"

	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func main() {
	Db = database.Db()
	csvFile, _ := os.Open("test.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		a, _ := strconv.ParseInt(line[1], 10, 64)
		n := line[0]
		g := line[2]

		Db.Create(&models.Person{
			Name:   n,
			Age:    a,
			Gender: g,
		})

	}
}
