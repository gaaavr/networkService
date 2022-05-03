package data

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}
type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incident  []IncidentData           `json:"incident"`
}

var Result ResultSetT

func GetResultsData() {
	go GetResultsSMS()
	go GetResultsMMS()
	go GetResultsVC()
	go GetResultsEmail()
	go GetResultsBilling()
	go GetResultsSupport()
	go GetResultsIncident()
}

func NetworkService() {
	GetResultsData()
	ticker := time.Tick(30 * time.Second)
	go func() {
		for range ticker {
			GetResultsData()
		}
	}()
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)
	http.ListenAndServe("localhost:8282", r)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	var status ResultT
	if Result.SMS == nil || Result.MMS == nil || Result.VoiceCall == nil || len(Result.Email) == 0 {
		status.Error = "Error on collect data"
	} else if Result.Billing == (BillingData{}) || Result.Support == nil || Result.Incident == nil {
		status.Error = "Error on collect data"
	} else {
		status.Status = true
		status.Data = Result
	}
	finalResult, err := json.Marshal(status)
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(finalResult)
	return

}
