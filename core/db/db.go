package db

import (
	"reflect"
	"errors"
	
	arangoDB "github.com/hostelix/aranGO"
	"github.com/merakiVE/CVDI/core/config"
	"github.com/merakiVE/CVDI/core/utils"
)

var DBHOST, DBUSER, DBPASSWORD, DBNAME string
var DBLOG bool

var configGlobal config.Configuration

func init() {

	configGlobal = config.Configuration{}

	configGlobal.Load()

	DBHOST = configGlobal.GetString("DATABASE.DB_HOST") + ":" + configGlobal.GetString("DATABASE.DB_PORT")
	DBUSER = configGlobal.GetString("DATABASE.DB_USER")
	DBPASSWORD = configGlobal.GetString("DATABASE.DB_PASSWORD")
	DBNAME = configGlobal.GetString("DATABASE.DB_NAME")
	DBLOG = false
}

func GetSessionDB() *arangoDB.Session {
	//Connection ArangoDB
	s, err := arangoDB.Connect(DBHOST, DBUSER, DBPASSWORD, DBLOG)

	if err != nil {
		panic(err)
	}
	return s
}

func GetCurrentDatabase() *arangoDB.Database {
	return GetSessionDB().DB(DBNAME)
}

func GetDatabase(nameDB string) *arangoDB.Database {
	return GetSessionDB().DB(nameDB)
}

func Save(db *arangoDB.Database, model arangoDB.Modeler) error {
	return db.Col(model.GetCollection()).Save(model)
}

func SaveModel(db *arangoDB.Database, model arangoDB.Modeler) (bool) {

	if reflect.ValueOf(model).Kind() != reflect.Ptr {
		panic("Check model must be a pointer")
	}

	ctx, err := arangoDB.NewContext(db)

	if err != nil {
		panic(err)
	}

	// save model, returns map of errors or empty map
	// check errors, also Error is saved in Context struct
	if e := ctx.Save(model); len(e) >= 1 {
		return false
	}
	return true
}

func GetModel(db *arangoDB.Database, model arangoDB.Modeler) (bool) {

	if reflect.ValueOf(model).Kind() != reflect.Ptr {
		panic("Check model must be a pointer")
	}

	ctx, err := arangoDB.NewContext(db)

	if err != nil {
		panic(err)
	}

	// save model, returns map of errors or empty map
	// check errors, also Error is saved in Context struct
	if e := ctx.Get(model); len(e) >= 1 {
		return false
	}
	return true
}

func ReplaceModel(db *arangoDB.Database, model arangoDB.Modeler) (error) {

	if utils.IsEmptyString(model.GetKey()) {
		panic("Check key is empty")
	}

	return db.Col(model.GetCollection()).Replace(model.GetKey(), model)
}

func FindModel(db *arangoDB.Database, model arangoDB.Modeler, filter ...interface{}) (error) {
	kindModel := reflect.ValueOf(model).Kind()

	if kindModel == reflect.Slice || kindModel == reflect.Array {
		panic("Check model no must be a slice or array")
	}

	aql := arangoDB.NewAqlStruct()
	aql.For("v", model.GetCollection())
	aql.Filter(filter...)
	aql.Return("v")

	c, err := aql.Execute(db)

	if err != nil {
		return err
	}
	c.FetchOne(model)
	return nil
}

func Model(model arangoDB.Modeler, db *arangoDB.Database) *modelDB {
	return &modelDB{
		modeler: model,
		db:      db,
	}
}

type modelDB struct {
	modeler arangoDB.Modeler
	db      *arangoDB.Database
}

func (this *modelDB) Create(m *arangoDB.Modeler) (error) {
	ctx, err := arangoDB.NewContext(this.db)
	if err != nil {
		return err
	}
	if e := ctx.Save(m); len(e) >= 1 {
		return errors.New("Error save model in database")
	}
	return nil
}

func (this *modelDB) FindOne(out interface{}, filter ...interface{}) (error) {
	aql := arangoDB.NewAqlStruct()
	aql.For("v", this.modeler.GetCollection())
	aql.Filter(filter...)
	aql.Return("v")

	c, err := aql.Execute(this.db)
	if err != nil {
		return err
	}

	if ok := c.FetchOne(out); !ok {
		return errors.New("Error get data")
	}

	return nil
}

func (this *modelDB) Find(out interface{}, filter ...interface{}) (error) {
	aql := arangoDB.NewAqlStruct()
	aql.For("v", this.modeler.GetCollection())
	aql.Filter(filter...)
	aql.Return("v")

	c, err := aql.Execute(this.db)
	if err != nil {
		return err
	}
	return c.FetchBatch(out)
}
