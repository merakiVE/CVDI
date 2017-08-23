package models

import (
	"testing"
)

func TestProcedure(t *testing.T) {

	a1 := []Activity{
		{"Cotizacion de Pedido", "/banesco/debitar", "id1", 1},
		{"Solicitud a maquilador", "/saime/get/cedula", "id2", 2},
	}

	a2 := []Activity{
		{"Envio a cliente", "/banesco/debitar", "id3", 3},
		{"Recepcion de pago", "/saime/get/cedula", "id4", 4},
	}

	lan := []Lane{
		{
			Name:       "Cliente",
			InPool:     false,
			NamePool:   "",
			Activities: a1,
		},
		{
			Name:       "Comercializador",
			InPool:     false,
			NamePool:   "",
			Activities: a2,
		},
	}

	b := Bpmn{
		Lanes: lan,
	}

	ac := b.GetSequenceActivities()

	if len(ac) > 0 {
		t.Log(b.GetSequenceActivities())
	} else {
		t.Error("Error execute test")
	}

}
