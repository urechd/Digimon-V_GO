package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urechd/digimon/internal/models"
)

type DigimonTable struct {
	Id    int
	Name  string
	Type  int
	Stage int
}

const (
	dbPath                      = "./resources/digimon.db"
	createDigimonInfoTableQuery = "create table DigimonInfo (Id integer not null primary key, Name text, Type integer, Stage integer);"
	insertDigimonQuery          = "insert into DigimonInfo (Id, Name, Type, Stage) values (?, ?, ?, ?)"
	getDigimonInfoByIdQuery     = "select * from DigimonInfo where Id = ?"
)

var db *sql.DB

func GetFileData(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func InitializeDB() {
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(dbPath)
		if err != nil {
			panic(err.Error())
		}
		file.Close()

		openDB()

		_, err = db.Exec(createDigimonInfoTableQuery)
		if err != nil {
			panic(err.Error())
		}

		insertDigimon()
	} else if err != nil {
		panic(err.Error())
	} else {
		openDB()
	}
}

func GetDigimonInfoById(id int) models.DigimonInfo {
	digiInfo := models.DigimonInfo{}
	row, err := db.Query(getDigimonInfoByIdQuery, id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&(digiInfo.Id), &(digiInfo.Name), &(digiInfo.Type), &(digiInfo.Stage))
	}

	return digiInfo
}

func openDB() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err.Error())
	}
}

func insertDigimon() {
	digimons := []models.Digimon{}
	dat := GetFileData("./resources/digimon.json")
	err := json.Unmarshal(dat, &digimons)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare(insertDigimonQuery)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(digimons); i++ {
		digi := digimons[i].Info
		_, err = stmt.Exec(digi.Id, digi.Name, digi.Type, digi.Stage)
		if err != nil {
			panic(err.Error())
		}
	}
}
