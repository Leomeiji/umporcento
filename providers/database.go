package providers

import (
	"database/sql"
	"strings"
)

type SQLProvider interface {
	Connect() error
}

func ExecuteSQLScriptLineByLine(script string, db *sql.DB) error {
	for _, entry := range strings.Split(script, "\n") {
		if _, err := db.Exec(entry); err != nil {
			return err
		}
	}

	return nil
}
