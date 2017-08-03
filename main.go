package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"

	"github.com/merakiVE/CVDI/src/controllers"
	packageConfig "github.com/merakiVE/CVDI/core/config"
	"github.com/merakiVE/CVDI/core"
)

const (
	PORT_SERVER = ":8101"
)

func main() {

	///Iris
	app := iris.New()

	// Init Configuration var
	config := packageConfig.Configuration{}

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

	//Load configuration
	config.Load()

	//Context Controller
	contextController := core.ContextController{App: app, Config: config}

	//Init Controllers
	cAuth := controllers.NewAuthController(contextController)
	cUser := controllers.NewUserController(contextController)
	cNeuron := controllers.NewNeuronController(contextController)

	//Routers
	routerUsers := app.Party("/users")
	{
		routerUsers.Get("/", cUser.List)
		routerUsers.Post("/", cUser.Create)
	}

	routerAdmin := app.Party("/auth")
	{
		routerAdmin.Post("/login", cAuth.Login)
	}

	routerNeuron := app.Party("/neurons")
	{
		routerNeuron.Get("/", cNeuron.List)
		routerNeuron.Get("/{key:string}", cNeuron.Get)
		routerNeuron.Post("/subscription", cNeuron.Subscribe)

		//Action Neuron
		routerNeuron.Get("/{key:string}/actions", cNeuron.Actions)
	}

	app.Run(iris.Addr(PORT_SERVER), iris.WithCharset("UTF-8"))
}
