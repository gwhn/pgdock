package main

import (
	"net/http"
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
//postgres://postgres:1234@localhost/test_pg?sslmode=verify-full
func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "connecting to postgres")
	db, err := sql.Open(
		"postgres", 
		"user=postgres password=1234 dbname=test_db host=pg_db sslmode=disable")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()
	fmt.Fprintln(w, "running query")
	rows, err := db.Query(`select r.id, r.name, c.name 
		from recipes r, categories c 
		where r.category_id=c.id`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()
	fmt.Fprintln(w, "scanning rows")
	for rows.Next() {
		var id int
		var rname, cname string
		err = rows.Scan(&id, &rname, &cname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "id=%d, recipe_name=%q, category_name=%q\n", id, rname, cname)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "completed successfully!")
}
