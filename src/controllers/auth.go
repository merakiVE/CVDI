package controllers

import (
	"fmt"
	"log"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris"

	arangoDB "github.com/hostelix/aranGO"

	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/auth"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/utils"
	"github.com/spf13/viper"
)

type AuthController struct {
	Configuration *viper.Viper
}

func (this AuthController) Login(_context context.Context) {

	var _form types.UserCredentials
	var _user models.UserModel

	err := _context.ReadJSON(&_form)

	if err != nil {

		_context.StatusCode(iris.StatusBadRequest)
		_context.JSON(types.ResponseAPI{
			Message: "Invalid data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	sq := fmt.Sprintf("FOR user IN users FILTER user.username == '%s' RETURN user", _form.Username)

	q := arangoDB.NewQuery(sq)
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

	if !cur.FetchOne(&_user) {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: "Error for get data user",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	if auth.VerifyPassword([]byte(_user.Password), []byte(_form.Password)) {

		//Read Private Key
		_secret, err := utils.ReadBinaryFile(this.Configuration.GetString("PRIVATE_KEY_PATH"))
		if err != nil {
			log.Fatal("Error reading private key")
			return
		}

		_token := auth.CreateTokenJWT(map[string]interface{}{"id": _user.Id, "key": _user.Key, "username": _user.Username }, _secret)

		_context.StatusCode(iris.StatusOK)

		_context.JSON(types.ResponseAPI{
			Message: "Login success",
			Data: types.JsonObject{
				"token": _token,
			},
			Errors: nil,
		})
		return

	} else {

		_context.StatusCode(iris.StatusForbidden)
		return
	}
}
