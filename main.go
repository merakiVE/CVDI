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
	// Secrect key testing
	PRIVATE_KEY = `AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSU
					GPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3
					Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XA
					t3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/En
					mZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbx
					NrRFi9wrf+M7Q==`
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
	var _form types.UserCredentials
	var _user models.UserModel

	err := ctx.ReadJSON(&_form)

	if err != nil {

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(types.ResponseAPI{
			Message: "Invalid data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	sq := fmt.Sprintf("FOR user IN users FILTER user.username == '%s' RETURN user", _form.Username)

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

	if !cur.FetchOne(&_user) {

		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Error for get data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	if auth.VerifyPassword([]byte(_user.Password), []byte(_form.Password)) {

		_token := auth.CreateTokenJWT(map[string]interface{}{"id": _user.Key, "username": _user.Username}, []byte(PRIVATE_KEY))

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(types.ResponseAPI{
			Message: "Login success",
			Data: types.JsonObject{
				"token": _token,
			},
			Errors: nil,
		})
		return

	} else {

		ctx.StatusCode(iris.StatusForbidden)
		return
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
