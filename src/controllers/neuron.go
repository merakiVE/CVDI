package controllers

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris"

	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core"

	arangoDB "github.com/hostelix/aranGO"

	"fmt"
)

type NeuronController struct {
	context core.ContextController
}

func NewNeuronController(cc core.ContextController) (NeuronController) {
	controller := NeuronController{}
	controller.SetContext(cc)

	return controller
}

func (this *NeuronController) SetContext(cc core.ContextController) {
	this.context = cc
}

func (this NeuronController) Get(_context context.Context) {

	var result models.NeuronModel
	var err error

	key_neuron := _context.Params().Get("key")

	query := fmt.Sprintf(`FOR neuron IN neurons FILTER neuron._key == '%s' RETURN neuron`, key_neuron)

	q := arangoDB.NewQuery(query)

	cur, err := db.GetDatabase(this.context.Config.GetString("DATABASE.DB_NAME")).Execute(q)

	if err != nil {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: err.Error(),
			Data:    nil,
			Errors:  nil,
		})
		return
	}
	success := cur.FetchOne(&result)

	if !success {

		_context.StatusCode(iris.StatusNotFound)
		_context.JSON(types.ResponseAPI{
			Message: "Error get neuron, key not found",
			Data:    nil,
			Errors:  map[string]string{"error": "key not found"},
		})
		return
	}

	_context.StatusCode(iris.StatusOK)
	_context.JSON(types.ResponseAPI{
		Message: "Neuron " + key_neuron,
		Data:    result,
		Errors:  nil,
	})
}

func (this NeuronController) List(_context context.Context) {

	result := make([]models.NeuronModel, 0)
	var err error

	q := arangoDB.NewQuery(`
		FOR neuron in neurons
		RETURN neuron
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

func (this NeuronController) Actions(_context context.Context) {

	var result models.NeuronModel
	var err error

	key_neuron := _context.Params().Get("key")

	query := fmt.Sprintf(`FOR neuron IN neurons FILTER neuron._key == '%s' RETURN neuron`, key_neuron)

	q := arangoDB.NewQuery(query)

	cur, err := db.GetDatabase(this.context.Config.GetString("DATABASE.DB_NAME")).Execute(q)

	if err != nil {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: err.Error(),
			Data:    nil,
			Errors:  nil,
		})
		return
	}
	success := cur.FetchOne(&result)

	if !success {

		_context.StatusCode(iris.StatusInternalServerError)
		_context.JSON(types.ResponseAPI{
			Message: "Error get actions, key not found",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	_context.StatusCode(iris.StatusOK)
	_context.JSON(types.ResponseAPI{
		Message: "List Actions Neuron",
		Data:    result.Actions,
		Errors:  nil,
	})
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

	success := db.SaveModel(db.GetDatabase(this.context.Config.GetString("DATABASE.DB_NAME")), &_neuron)

	if success {
		_context.StatusCode(iris.StatusCreated)
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
