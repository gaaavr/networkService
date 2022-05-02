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

func GetResultsIncident() []IncidentData {
	var incident *IncidentData
	var arrIncident []IncidentData
	content, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		return arrIncident
	}
	if content.StatusCode == 500 {
		return arrIncident
	}
	data, err := io.ReadAll(content.Body)
	if err != nil {
		return arrIncident
	}
	defer content.Body.Close()
	strData := string(data)
	arrData := service.FormattingString(strData)
	var element []byte
	for i, _ := range arrData {
		element = []byte(strings.Trim(arrData[i], " ,"))
		if err := json.Unmarshal(element, &incident); err != nil {
			return arrIncident
		}
		arrIncident = append(arrIncident, *incident)
	}
	sort.SliceStable(arrIncident, func(i, j int) bool {
		return arrIncident[i].Status < arrIncident[j].Status
	})
	return arrIncident
}
