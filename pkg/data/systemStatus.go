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

var data ResultSetT

func GetResultsData() ResultSetT {
	var result ResultSetT
	result.SMS = GetResultsSMS()
	result.MMS = GetResultsMMS()
	result.VoiceCall = GetResultsVC()
	result.Email = GetResultsEmail()
	result.Billing = GetResultsBilling()
	result.Support = GetResultsSupport()
	result.Incident = GetResultsIncident()
	return result
}

func NetworkService() {
	data = GetResultsData()
	ticker := time.Tick(30 * time.Second)
	go func() {
		for range ticker {
			data = GetResultsData()
		}
	}()
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)
	http.ListenAndServe("localhost:8282", r)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	var status ResultT
	if data.SMS == nil || data.MMS == nil || data.VoiceCall == nil || len(data.Email) == 0 {
		status.Error = "Error on collect data"
	} else if data.Billing == (BillingData{}) || data.Support == nil || data.Incident == nil {
		status.Error = "Error on collect data"
	} else {
		status.Status = true
		status.Data = data
	}
	finalResult, err := json.Marshal(status)
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(finalResult)
	return

}
