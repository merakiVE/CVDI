package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"

	"github.com/merakiVE/CVDI/src/controllers"
	"github.com/merakiVE/CVDI/core/config"
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

	//Load configuration
	config.Load()

	app.Use(APILogger)

	//Init Controllers
	cAuth := controllers.AuthController{Configuration: config.GetConfig() }
	cUser := controllers.UserController{Configuration: config.GetConfig() }
	cNeuron := controllers.NeuronController{Configuration: config.GetConfig() }

	routerUsers := app.Party("/users")
	{
		routerUsers.Get("/", cUser.List)
		routerUsers.Post("/create", cUser.Create)
	}

	routerAdmin := app.Party("/auth")
	{
		routerAdmin.Post("/login", cAuth.Login)
	}

	routerNeuron := app.Party("/neuron")
	{
		routerNeuron.Get("/", cNeuron.List)
		routerNeuron.Post("/subscription", cNeuron.Subscribe)
	}

	app.Run(iris.Addr(PORT_SERVER), iris.WithCharset("UTF-8"))
}
