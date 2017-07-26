package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"

	arangoDB "github.com/hostelix/aranGO"

	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/CVDI/core/auth"
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

	routerAdmin := app.Party("/auth")
	{
		routerAdmin.Post("/login", login)
	}

	app.Run(iris.Addr(PORT_SERVER), iris.WithCharset("UTF-8"))
}

func login(ctx context.Context) {
	var _form types.UserLogin
	var _user models.UserModel

	err := ctx.ReadJSON(&_form)

	if err != nil {

		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Invalid data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	sq := fmt.Sprintf("FOR user in users FILTER user.username == '%s' RETURN user", _form.Username)

	q := arangoDB.NewQuery(sq)
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

	_err := cur.FetchOne(&_user)

	if !_err {
		fmt.Println("Error get user")
	}

	if auth.VerifyPassword([]byte(_user.Password), []byte(_form.Password)) {
		fmt.Println("Es el password del usuario")
	} else {
		fmt.Println(_form)
		fmt.Println("No es el password")
	}
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
