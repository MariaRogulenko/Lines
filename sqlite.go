package lines

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	sqlite3 "github.com/mattn/go-sqlite3"
)

var database *sql.DB
var err error

// StoreItem store
func StoreItem(params *DBComminication) {

	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sqlTable := "CREATE TABLE IF NOT EXISTS Game(Id TEXT NOT NULL PRIMARY KEY,Name TEXT NOT NULL, CurrScore INT, Board TEXT, ActX INT default -1, ActY INT default -1)"
	_, err = database.Exec(sqlTable)
	if err != nil {
		log.Fatal(err)
	}

	sqlAdditem := `
	INSERT OR REPLACE INTO Game(
		Id,
		Name,
		CurrScore,
		Board,
		ActX,
		ActY
	) values(?, ?, ?, ?, ?, ?)
	`

	stmt, err := database.Prepare(sqlAdditem)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(params.id, params.username, params.score, encodeTable(params.table), params.active.x, params.active.y)

	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

// ReadItem c
func ReadItem(id string) *DBComminication {
	sqlReadAll := `
	SELECT * FROM Game 
    WHERE Id=?
	`

	rows, err := database.Query(sqlReadAll, id)
	if err != nil {
		panic(err)
	}
	if rows == nil {
		fmt.Println("row")
		return nil
	}
	defer rows.Close()

	result := new(DBComminication)

	var eTable string
	for rows.Next() {
		err2 := rows.Scan(&result.id, &result.username, &result.score, &eTable, &result.active.x, &result.active.y)
		if err2 != nil {
			panic(err2)
		}
	}
	if result.id == "" {
		fmt.Println(id + "   !zero id!")
		return nil
	}
	result.table = decodeTable(eTable)
	return result
}

// RegisterAndOpenDB c
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