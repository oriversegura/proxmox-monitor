package database

import (
	"database/sql"
	"log"
)

func CreateSchema(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS Node (id INTEGER PRIMARY KEY AUTOINCREMENT, 
	name text,
	 ip text,
	 imestamp text,
	 status text, 
	hostname text )	
	`)
	if err != nil {
		log.Fatal(err)
	}
	
	
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS Vm (id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name text, 
	 ip text, 
	 timestamp text, 
	 status text, 
	 vm_id integer, 
	kind text )`)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
