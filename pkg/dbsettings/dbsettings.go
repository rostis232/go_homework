package dbsettings

import "github.com/upper/db/v4/adapter/postgresql"

var Settings = postgresql.ConnectionURL{
	Database: `go_homework`,
	Host:     `localhost:5432`,
	User:     `admin`,
	Password: `admin`,
}
