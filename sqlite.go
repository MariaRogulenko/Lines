package lines

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"errors"

	sqlite3 "github.com/mattn/go-sqlite3"
)

var database *sql.DB
var err error

func createTable() {
	sqlTable := "CREATE TABLE IF NOT EXISTS Game(Id TEXT NOT NULL PRIMARY KEY,Name TEXT NOT NULL,BestScore INT, CurrScore INT, Board TEXT, ActX INT default -1, ActY INT default -1, NextColors TEXT)"
	_, err = database.Exec(sqlTable)
	if err != nil {
		log.Fatal(err)
	}
}

// StoreItem writes to database
func StoreItem(params *DBCommunication) error {
	sqlAdditem := `
	INSERT OR REPLACE INTO Game(
		Id,
		Name,
		BestScore,
		CurrScore,
		Board,
		ActX,
		ActY,
		NextColors
	) values(?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := database.Prepare(sqlAdditem)
	if err != nil {
		return errors.New("Failed to prepare DB for write")
	}
	_, err = stmt.Exec(params.id, params.username, params.bestScore, params.score, encodeTable(params.table), params.active.x, params.active.y, encodeTable(params.nextColors))
	if err != nil {
		return errors.New("Failed to write to DB")
	}
	return nil
}

// ReadItem reads to database
func ReadItem(id string) (*DBCommunication, error) {
	sqlReadAll := `
	SELECT * FROM Game 
    WHERE Id=?
	`
	rows, err := database.Query(sqlReadAll, id)
	fmt.Println(rows)
	if err != nil {
		return nil, errors.New("Failed to read from DB")
	}
	defer rows.Close()
	result := new(DBCommunication)
	var eTable string
	var eColors string
	for rows.Next() {
		err = rows.Scan(&result.id, &result.username, &result.bestScore, &result.score, &eTable, &result.active.x, &result.active.y, &eColors)
		if err != nil {
			return nil, errors.New("Scan")
		}
	}
	if result.id == "" {
		return nil, nil
	}
	result.table = decodeTable(eTable)
	result.nextColors = decodeTable(eColors)
	return result, nil
}

// RegisterAndOpenDB registers and opens database
func RegisterAndOpenDB() {
	var DBDriver string
	sql.Register(DBDriver, &sqlite3.SQLiteDriver{})
	database, err = sql.Open(DBDriver, "../db/lines.db")
	if err != nil {
		fmt.Println("Failed to create the handle")
	}
	if err2 := database.Ping(); err2 != nil {
		fmt.Println("Failed to keep connection alive")
	}
	createTable()
}

func encodeTable(dTable []int32) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dTable)), ","), "[]")
}

func decodeTable(eTable string) []int32 {
	var temp = strings.Split(eTable, ",")
	var dTable = []int32{}
	for _, i := range temp {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		dTable = append(dTable, int32(j))
	}
	return dTable
}
