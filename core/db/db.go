package db

import (
	arangoDB "github.com/hostelix/aranGO"
	"reflect"
	"github.com/merakiVE/CVDI/core/config"
)

var DBHOST, DBUSER, DBPASSWORD string

func init() {
	config.Load()

	DBHOST = config.GetString("DATABASE.DB_HOST") + ":" + config.GetString("DATABASE.DB_PORT")
	DBUSER = config.GetString("DATABASE.DB_USER")
	DBPASSWORD = config.GetString("DATABASE.DB_PASSWORD")

}

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

func SaveModel(_db *arangoDB.Database, _model arangoDB.Modeler) (bool) {

	if reflect.ValueOf(_model).Kind() != reflect.Ptr {
		panic("Check model must be a pointer")
	}

	ctx, err := arangoDB.NewContext(_db)

	if err != nil {
		panic(err)
	}

	// save model, returns map of errors or empty map
	e := ctx.Save(_model)

	// check errors, also Error is saved in Context struct
	if len(e) >= 1 {
		return false
	}

	return true
}
