//This example uses the ORM jet
package main

import (
	"fmt"
	"log"

	"github.com/eaigner/jet"
	"github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	//Make sure you setup the ELEPHANTSQL_URL to be a uri, e.g. 'postgres://user:pass@host/db?options'
	pgURL, err := pq.ParseURL("postgres://elnbxzcj:EhKpzdq9m8jvluh8o8KgXgrH5JFBnD0x@drona.db.elephantsql.com:5432/elnbxzcj")
	logFatal(err)
	db, err := jet.Open("postgres", pgURL)
	logFatal(err)
	var p map[string]interface{}
	err = db.Query("SELECT id, data, role_id FROM users").Rows(&p)
	logFatal(err)
	log.Println(len(p))
	log.Println(fmt.Sprintf("%s", p["id"]))
	log.Println(fmt.Sprintf("%s", p["role_id"]))
	log.Println(fmt.Sprintf("%s", p["data"]))
}
