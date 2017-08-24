package models

import (
	arangoDB "github.com/hostelix/aranGO"

	"github.com/merakiVE/CVDI/core/types"
	"github.com/merakiVE/CVDI/core/validator"
	"github.com/merakiVE/CVDI/core/tags"
	"sort"
)

const (
	TYPE_NODE_START        = 1
	TYPE_NODE_END          = 2
	TYPE_NODE_TASK         = 3
	TYPE_NODE_USER_TASK    = 4
	TYPE_NODE_SERVICE_TASK = 5
	TYPE_NODE_GATEWAY      = 6
	TYPE_NODE_TIMER        = 7
	TYPE_NODE_MESSAGE      = 8
)

type Node struct {
	Type        int
	NextNode    *Node
	PreviusNode *Node
	Data        interface{}
}

type TypeNode interface {
	GetType() string
	GetNexNode() *Node
	GetPreviosNode() *Node
	GetData() interface{}
}

//Structs for BPMN
type Lane struct {
	Name       string        `json:"name"`
	InPool     bool          `json:"in_pool"`
	NamePool   string        `json:"name_pool"`
	Activities []Activity    `json:"activities,omitempty" validate:"required"`
}

type Bpmn struct {
	Lanes []Lane `json:"lanes"`
}

func (this Bpmn) GetSequenceActivities() ([]Activity) {
	activities_tmp := make([]Activity, 0)

	for _, lane := range this.Lanes {
		for _, acti := range lane.Activities {
			activities_tmp = append(activities_tmp, acti)
		}
	}

	sort.Slice(activities_tmp, func(i, j int) bool {
		return activities_tmp[i].Sequence < activities_tmp[j].Sequence
	})

	return activities_tmp
}

type Activity struct {
	Name      string `json:"name"`
	NeuronKey string `json:"neuron_key"`
	ActionID  string `json:"action_id"`
	Sequence  int32    `json:"sequence"`
	InputData map[string]interface{} `json:"inputs"`
}

type ProcedureModel struct {
	arangoDB.Document

	ID    string        `json:"id" validate:"required" on_create:"set,auto_uuid"`
	Owner string        `json:"owner,omitempty" validate:"required"`

	types.Timestamps
	ErrorsValidation []map[string]string `json:"errors_validation,omitempty"`
}

func (this ProcedureModel) GetKey() string {
	return this.Key
}

func (this ProcedureModel) GetCollection() string {
	return "procedures"
}

func (this ProcedureModel) GetError() (string, bool) {
	return this.Message, this.Error
}

func (this ProcedureModel) GetValidationErrors() ([]map[string]string) {
	return this.ErrorsValidation
}

func (this *ProcedureModel) PreSave(c *arangoDB.Context) {

	v := validator.New()

	v.Validate(this)

	if v.IsValid() {

		//Tag Process for model
		tags.New().ProcessTags(this)
	} else {

		c.Err["error_validation"] = "Error validating model"
		this.ErrorsValidation = v.GetMessagesValidation()
	}

	return
}
