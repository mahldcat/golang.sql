package golangdb

import (
	"context"
	"database/sql"
	"log"
	"time"

	mssql "github.com/microsoft/go-mssqldb"
)

func GetDb(connection string) (*sql.DB, error) {

	db, err := sql.Open("sqlserver", connection)

	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return db, err
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return db, err
	}

	err = db.PingContext(ctx)

	return db, err
}

func GetTaskCt(db *sql.DB) (int, error) {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := "SELECT * FROM Tasks"

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	for rows.Next() {
		var id mssql.UniqueIdentifier
		var msg, status string
		var lastupdate time.Time
		var percentcomplete int

		err := rows.Scan(&id, &msg, &lastupdate, &status, &percentcomplete)
		if err != nil {
			return -1, err
		}

		count++

	}

	return count, nil
}
