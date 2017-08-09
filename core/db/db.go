package db

import (
	"reflect"
	"log"

	arangoDB "github.com/hostelix/aranGO"
	"github.com/merakiVE/CVDI/core/config"
	"github.com/merakiVE/CVDI/core/utils"
)

var DBHOST, DBUSER, DBPASSWORD string

var configGlobal config.Configuration

func init() {

	configGlobal = config.Configuration{}

	configGlobal.Load()

	DBHOST = configGlobal.GetString("DATABASE.DB_HOST") + ":" + configGlobal.GetString("DATABASE.DB_PORT")
	DBUSER = configGlobal.GetString("DATABASE.DB_USER")
	DBPASSWORD = configGlobal.GetString("DATABASE.DB_PASSWORD")

}

func GetSessionDB() *arangoDB.Session {

	//Connection ArangoDB
	s, err := arangoDB.Connect(DBHOST, DBUSER, DBPASSWORD, false)

	if err != nil {
		panic(err)
	}
	return s
}

func GetCurrentDatabase() *arangoDB.Database {
	return GetSessionDB().DB(configGlobal.GetString("DATABASE.DB_NAME"))
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

func GetModel(_db *arangoDB.Database, _model arangoDB.Modeler) (bool) {

	if reflect.ValueOf(_model).Kind() != reflect.Ptr {
		panic("Check model must be a pointer")
	}

	ctx, err := arangoDB.NewContext(_db)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// save model, returns map of errors or empty map
	e := ctx.Get(_model)

	// check errors, also Error is saved in Context struct
	if len(e) >= 1 {
		return false
	}

	return true
}

func ReplaceModel(_db *arangoDB.Database, _model arangoDB.Modeler) (bool) {

	if utils.IsEmptyString(_model.GetKey()) {
		panic("Check key is empty")
	}

	err := _db.Col(_model.GetCollection()).Replace(_model.GetKey(), _model)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}