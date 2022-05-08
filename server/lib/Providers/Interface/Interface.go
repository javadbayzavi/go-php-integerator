package Interface

import "database/sql"

type ProviderInterface interface {
	Connect() bool
	Query(params map[string]string)
	GetResult() *sql.Rows
}
