package db

import (
	"database/sql"
	"errors"
	"time"

	// Blank Import
	"github.com/acky666/CredentialsCapture/config"
	_ "github.com/go-sql-driver/mysql"

	l "github.com/acky666/ackyLog"
)

var databaseConnection *sql.DB
var MaxDatabaseOpenConnections = 25
var MaxDatabaseIdleConnections = 25
var DatabaseIdleTimeout = 5 * time.Minute
var ReconnectionAttempts = 3

func OpenDatabaseConnection() (*sql.DB, error) {

	if config.Get("DatabaseDSN") == "" {
		return nil, errors.New("Empty DatabaseDSN")
	}

	// Return the Connection, if we are connected!
	if databaseConnection != nil {
		return databaseConnection, nil
	}

	var DatabaseConnectionError error
	var dbConn *sql.DB

	for i := 0; i < ReconnectionAttempts; i++ {
		dbConn, DatabaseConnectionError = sql.Open("mysql", config.Get("DatabaseDSN"))

		if DatabaseConnectionError == nil {
			//
			// Verify your connnected
			//
			PingCheckErr := dbConn.Ping()

			if PingCheckErr == nil {
				break
			}

			l.ERROR("Ping to database connection '%s' failed  attempt %d, %s", config.Get("DatabaseDSN_LOG"), i, PingCheckErr.Error())
			time.Sleep(500 * time.Millisecond)
		}

		time.Sleep(500 * time.Millisecond)
		l.ERROR("Database connection '%s' failed", config.Get("DatabaseDSN_LOG"))
	}

	if DatabaseConnectionError != nil {
		return nil, DatabaseConnectionError
	}

	databaseConnection = dbConn
	databaseConnection.SetMaxOpenConns(MaxDatabaseOpenConnections)
	databaseConnection.SetMaxIdleConns(MaxDatabaseIdleConnections)
	databaseConnection.SetConnMaxLifetime(DatabaseIdleTimeout)

	return databaseConnection, nil
}

func ExecutePrepared(sqlQuery string, parameters ...interface{}) (int64, int64, error) {

	StopWatch := l.TIMED("[[F-MAGENTA]SQL[F-NORMAL]] Execute - '%s'", sqlQuery)

	DatabaseConnection, err := OpenDatabaseConnection()
	if err != nil {
		return 0, 0, err
	}

	Result, err := DatabaseConnection.Exec(sqlQuery, parameters...)
	if err != nil {
		return 0, 0, err
	}
	StopWatch()

	LastInsertedID, _ := Result.LastInsertId()
	RowsAffected, _ := Result.RowsAffected()

	l.DEBUG("Execute LastID %v and Rows Effected %v", LastInsertedID, RowsAffected)

	return LastInsertedID, RowsAffected, nil
}
