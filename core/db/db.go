package db

import (
	arangoDB "github.com/hostelix/aranGO"
)

const (
	DBHOST     = "http://localhost:8529"
	DBUSER     = "root"
	DBPASSWORD = "canaima"
)

func GetSessionDB() *arangoDB.Session {

	//Connection ArangoDB
	s, err := arangoDB.Connect(DBHOST, DBUSER, DBPASSWORD, false)

	if err != nil {
		panic(err)
	}
	return s
}

func GetDatabase(nameDB string) *arangoDB.Database {
	return GetSessionDB().DB(nameDB)
}

func Save(_db *arangoDB.Database, _model arangoDB.Modeler) error {

	err := _db.Col(_model.GetCollection()).Save(_model)

	return err
}
