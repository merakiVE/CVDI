package pool

import (
	"github.com/gocraft/work"
	"fmt"
	"net/url"
	"io/ioutil"
	"net/http"
	//"github.com/merakiVE/CVDI/core/db"
	//"github.com/merakiVE/CVDI/src/models"
)

type Context struct {
	customerID int64
}

func CreateEnqueuer() *work.Enqueuer {
	/*var procedure models.ProcedureModel
	var neuron models.NeuronModel

	procedure.SetKey("2771182")

	success := db.GetModel(db.GetCurrentDatabase(), &procedure)

	if !success {
		return
	}*/

	redisPool := NewPoolRedis()

	var enqueuer = work.NewEnqueuer("cvdi_namespace", redisPool)

	return enqueuer
	/*for _, value := range procedure.Activities {


		qq := arangoDB.NewQuery(`FOR doc IN neurons FILTER doc.id == '` + value.NeuronKey + `' RETURN doc`)

		cur , _ := db.GetCurrentDatabase().Execute(qq)

		cur.FetchOne(&neuron)

		for _,act := range neuron.Actions {


			if act.ID == value.ActionID {

				if value.Type == "activity" {

					_, err := enqueuer.Enqueue("process_activity_bpmn",
						work.Q{
							"url":       neuron.Host,
							"end_point": act.EndPoint,
							"params": act.Params,
						})

					if err != nil {
						log.Fatal(err)
					}

				}

			}
		}

	}*/
}

func CreateWorkerPool() *work.WorkerPool {
	redisPool := NewPoolRedis()

	pool := work.NewWorkerPool(Context{}, 10, "cvdi_namespace", redisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)

	// Map the name of jobs to handler functions
	pool.Job("worker_activity_bpmn", (*Context).ProcessBPMNActivity)

	// Customize options:
	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	//pool.Start()
	//pool.Stop()

	return pool
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (c *Context) ProcessBPMNActivity(job *work.Job) error {
	if err := job.ArgError(); err != nil {
		return err
	}

	urll := job.ArgString("url")
	end_point := job.ArgString("end_point")
	params := job.Args["params"]

	url_total := CreateURL(urll, end_point, params)

	resp, err := http.Get(url_total)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(url_total, string(body))

	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}

func CreateURL(base string, endPoint string, params interface{}) (string) {

	host, _ := url.Parse(base)
	action, _ := url.Parse(endPoint)

	u := host.ResolveReference(action)

	if params != nil {
		q := u.Query()

		dict_params := params.(map[string]interface{})

		for key, value := range dict_params {
			q.Set(key, value.(string))
		}

		u.RawQuery = q.Encode()
	}
	return u.String()
}
