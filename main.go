package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"

	arangoDB "github.com/diegogub/aranGO"

	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
	"github.com/merakiVE/CVDI/src/models"

	"github.com/fatih/structs"
	"fmt"
)

const (
	PORT_SERVER = ":8101"
)

func main() {

	///Iris
	app := iris.New()

	//app.Configure(iris.WithConfiguration(iris.YAML("./config_iris.yml")))

	APILogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(APILogger)

	routerUsers := app.Party("/users")
	{
		routerUsers.Get("/", getAllUsers)
		routerUsers.Post("/create", createUser)
	}

	routerAdmin := app.Party("/admin")
	{
		routerAdmin.Get("/neuron", getAllUsers)
	}

	app.Run(iris.Addr(PORT_SERVER), iris.WithCharset("UTF-8"))
}

func getAllUsers(ctx context.Context) {

	result := make([]models.UserModel, 0)
	var err error

	q := arangoDB.NewQuery("FOR i in users RETURN i")
	cur, err := db.GetDatabase("meraki").Execute(q)

	if err != nil {

		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Fail",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	err = cur.FetchBatch(&result)

	if err != nil {

		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Fail",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(types.ResponseAPI{
		Message: "Success",
		Data:    result,
		Errors:  nil,
	})
}

func createUser(ctx context.Context) {

	var _user models.UserModel

	var err error

	err = ctx.ReadJSON(&_user)

	if err != nil {

		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Invalid data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	v := validator.CreateValidator()

	v.Validate(&_user)

	if v.IsValid() {

		a := tags.New()
		a.ProcessTags(&_user)

		db.Save(db.GetDatabase("meraki"), _user)

		if err != nil {

			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(types.ResponseAPI{
				Message: "Error creating user",
				Data:    nil,
				Errors:  nil,
			})
		} else {
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(types.ResponseAPI{
				Message: "User created successfully",
				Data:    nil,
				Errors:  nil,
			})
		}
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(types.ResponseAPI{
			Message: "Error creating user, invalid data",
			Data:    nil,
			Errors:  v.GetMessagesValidation(),
		})
	}

}
