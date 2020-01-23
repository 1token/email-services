package database

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"reflect"
)

func List(db *sql.DB, table string, result interface{}, clause ...string) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}
	slicev := resultv.Elem()
	elemt := slicev.Type().Elem()
	query := "SELECT data FROM " + table

	if len(clause) > 0 {
		for _, v := range clause {
			query = query + " " + v
		}
	}
	log.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		jsonStr := ""
		err := rows.Scan(&jsonStr)
		if err != nil {
			return err
		}
		elemp := reflect.New(elemt)
		json.Unmarshal([]byte(jsonStr), elemp.Interface())
		slicev = reflect.Append(slicev, elemp.Elem())
		i++
	}
	resultv.Elem().Set(slicev.Slice(0, i))
	return nil
}
