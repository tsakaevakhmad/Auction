package dbcontext

import "database/sql"
import _ "github.com/lib/pq"

func getContext() *sql.DB {
	context, err := sql.Open("postgres", "user=postgres password=password dbname=auction sslmode=disable")
	if err != nil {
		panic(err)
	}
	return context
}
