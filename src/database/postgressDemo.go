package database

import (
	"app/utils"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitPostgressDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	utils.CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	utils.CheckError(err)

	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
	_, err = db.Exec(insertStmt)
	utils.CheckError(err)

	// dynamic
	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	_, err = db.Exec(insertDynStmt, "Jane", 2)
	utils.CheckError(err)

	// update
	updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	_, err = db.Exec(updateStmt, "Mary", 3, 2)
	utils.CheckError(err)

	// delete
	deleteStmt := `delete from "Students" where "id"=$1`
	_, err = db.Exec(deleteStmt, 1)
	utils.CheckError(err)

	// query
	rows, err := db.Query(`select "Name", "Roll" from "Students"`)
	utils.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var name string
		var roll int
		err = rows.Scan(&name, &roll)
		utils.CheckError(err)

		fmt.Println(name, roll)
	}

	utils.CheckError(err)

	fmt.Println("Connected!")
}
