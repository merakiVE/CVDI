package controllers

import (
	"github.com/kataras/iris"
	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/src/models"

	"github.com/kataras/iris/context"

	arangoDB "github.com/hostelix/aranGO"
	"github.com/spf13/viper"
)

type UserController struct {
	Configuration *viper.Viper
}

func (this UserController) List(_context context.Context) {

	result := make([]models.UserModel, 0)
	var err error

	q := arangoDB.NewQuery(`
		FOR user in users
		RETURN {
			"_key": user._key,
			"_id": user._id,
			"_rev": user._rev,
			"username": user.username,
			"email": user.email
		}
	`)
	cur, err := db.GetDatabase(this.Configuration.GetString("DATABASE.DB_NAME")).Execute(q)

	if err != nil {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: "Fail",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	err = cur.FetchBatch(&result)

	if err != nil {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: "Fail",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	_context.StatusCode(iris.StatusOK)
	_context.JSON(types.ResponseAPI{
		Message: "Success",
		Data:    result,
		Errors:  nil,
	})
}

func (this UserController) Create(_context context.Context) {

	var _user models.UserModel

	var err error

	err = _context.ReadJSON(&_user)

	if err != nil {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: "Invalid data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	success := db.SaveModel(db.GetDatabase(this.Configuration.GetString("DATABASE.DB_NAME")), &_user)

	if success {
		_context.StatusCode(iris.StatusOK)
		_context.JSON(types.ResponseAPI{
			Message: "User created successfully",
			Data:    nil,
			Errors:  nil,
		})

	} else {
		_context.StatusCode(iris.StatusOK)
		_context.JSON(types.ResponseAPI{
			Message: "Error creating user, invalid data",
			Data:    nil,
			Errors:  _user.GetValidationErrors(),
		})
	}

}
