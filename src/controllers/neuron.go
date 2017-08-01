package controllers

import (
	"github.com/kataras/iris/context"
	"github.com/merakiVE/CVDI/src/models"
	"github.com/kataras/iris"
	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/spf13/viper"
)

type NeuronController struct {
	Configuration *viper.Viper
}

func (this NeuronController) Subscribe(_context context.Context) {
	var _neuron models.NeuronModel

	var err error

	err = _context.ReadJSON(&_neuron)

	if err != nil {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: "Invalid data Neuron",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	success := db.SaveModel(db.GetDatabase(this.Configuration.GetString("DATABASE.DB_NAME")), &_neuron)

	if success {
		_context.StatusCode(iris.StatusOK)
		_context.JSON(types.ResponseAPI{
			Message: "Neuron subscribe successfully",
			Data:    nil,
			Errors:  nil,
		})

	} else {
		_context.StatusCode(iris.StatusOK)
		_context.JSON(types.ResponseAPI{
			Message: "Error subscribing neuron, invalid data",
			Data:    nil,
			Errors:  _neuron.GetValidationErrors(),
		})
	}
}
