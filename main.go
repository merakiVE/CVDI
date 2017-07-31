package main

import (
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"

	"github.com/merakiVE/CVDI/src/controllers"
	"github.com/merakiVE/CVDI/core/utils"
)

const (
	PORT_SERVER = ":8101"
)

var (
	SecrectKey, PublicKey []byte
)

func initKeys() {
	var err error

	SecrectKey, err = utils.ReadSecrectKey()
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	PublicKey, err = utils.ReadPublicKey()
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func main() {

	initKeys()

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

	//Init Controllers
	cAuth := controllers.AuthController{}
	cUser := controllers.UserController{}
	cNeuron := controllers.NeuronController{}

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
		routerNeuron.Get("/")
		routerNeuron.Post("/subscription", cNeuron.Subscribe)
	}

	app.Run(iris.Addr(PORT_SERVER), iris.WithCharset("UTF-8"))
}
