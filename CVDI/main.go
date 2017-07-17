package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"

	arangoDB "github.com/diegogub/aranGO"

	"meraki/CVDI/types"
	"meraki/CVDI/validator"
)

const (
	DBHOST      = "http://localhost:8529"
	DBUSER      = "root"
	DBPASSWORD  = ""
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

func getSessionDB() *arangoDB.Session {
	//ArangoDB
	s, err := arangoDB.Connect(DBHOST, DBUSER, DBPASSWORD, false)

	if err != nil {
		panic(err)
	}

	return s
}

func getAllUsers(ctx context.Context) {

	result := make([]types.User, 0)
	var err error

	q := arangoDB.NewQuery("FOR i in users RETURN i")
	cur, err := getSessionDB().DB("meraki").Execute(q)

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

	var _user types.User
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
		err = getSessionDB().DB("meraki").Col("users").Save(&_user)

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
