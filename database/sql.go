package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/otaxhu/api-rest-golang/settings"
)

func NewSqlConnection(dbSettings *settings.Database) (*sql.DB, error) {
	return sql.Open(dbSettings.Driver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			dbSettings.User,
			dbSettings.Password,
			dbSettings.Host,
			dbSettings.Port,
			dbSettings.Name,
		),
	)
}
