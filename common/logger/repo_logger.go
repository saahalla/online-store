package logger

import (
	"log"

	"github.com/doug-martin/goqu/v9"
)

func LogQueryInsert(dataset *goqu.InsertDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}

func LogQuerySelect(dataset *goqu.SelectDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}

func LogQueryDelete(dataset *goqu.DeleteDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}

func LogQueryUpdate(dataset *goqu.UpdateDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}
