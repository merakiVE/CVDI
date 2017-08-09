package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core"
	"github.com/merakiVE/CVDI/src/models"

	arangoDB "github.com/hostelix/aranGO"
)

type UserController struct {
	context core.ContextController
}

func NewUserController(cc core.ContextController) (UserController) {
	controller := UserController{}
	controller.SetContext(cc)
	controller.RegisterRouters()
	return controller
}

func (this *UserController) RegisterRouters() {

	app := this.context.App

	//User Routers
	routerUsers := app.Party("/users")
	routerUsers.Get("/{userKey:string}", this.GetUser)
	routerUsers.Get("/", this.ListUsers)
	routerUsers.Post("/", this.CreateUser)
}

func (this *UserController) SetContext(cc core.ContextController) {
	this.context = cc
}

func (this UserController) GetUser(_context context.Context) {

	var user models.UserModel

	key_user := _context.Params().Get("userKey")

	user.SetKey(key_user)
	success := db.GetModel(db.GetCurrentDatabase(), &user)

	if !success {

		_context.StatusCode(iris.StatusNotFound)
		_context.JSON(types.ResponseAPI{
			Message: "Error get user, key not found",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	_context.StatusCode(iris.StatusOK)
	_context.JSON(types.ResponseAPI{
		Message: "User " + key_user,
		Data:    user,
		Errors:  nil,
	})
}

func (this UserController) ListUsers(_context context.Context) {

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
	cur, err := db.GetDatabase(this.context.Config.GetString("DATABASE.DB_NAME")).Execute(q)

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

func (this UserController) CreateUser(_context context.Context) {

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

	success := db.SaveModel(db.GetDatabase(this.context.Config.GetString("DATABASE.DB_NAME")), &_user)

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
