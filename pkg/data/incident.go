package data

import (
	"encoding/json"
	"io"
	"net/http"
	"networkService/pkg/service"
	"sort"
	"strings"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func GetResultsIncident() {
	var incident *IncidentData
	var arrIncident []IncidentData
	content, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		Result.Incident = arrIncident
		return
	}
	if content.StatusCode == 500 {
		Result.Incident = arrIncident
		return
	}
	data, err := io.ReadAll(content.Body)
	if err != nil {
		Result.Incident = arrIncident
		return
	}
	defer content.Body.Close()
	strData := string(data)
	arrData := service.FormattingString(strData)
	var element []byte
	for i, _ := range arrData {
		element = []byte(strings.Trim(arrData[i], " ,"))
		if err := json.Unmarshal(element, &incident); err != nil {
			Result.Incident = arrIncident
			return
		}
		arrIncident = append(arrIncident, *incident)
	}
	sort.SliceStable(arrIncident, func(i, j int) bool {
		return arrIncident[i].Status < arrIncident[j].Status
	})
	Result.Incident = arrIncident
	return
}
