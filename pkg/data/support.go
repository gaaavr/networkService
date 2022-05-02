package data

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"networkService/pkg/service"
	"strings"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickers int    `json:"active_tickets"`
}

func GetResultsSupport() []int {
	var sup *SupportData
	var arrSup []SupportData
	var load []int
	content, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		return load
	}
	if content.StatusCode == 500 {
		return load
	}
	data, err := io.ReadAll(content.Body)
	if err != nil {
		return load
	}
	defer content.Body.Close()
	strData := string(data)
	arrData := service.FormattingString(strData)
	var element []byte
	for i, _ := range arrData {
		element = []byte(strings.Trim(arrData[i], " ,"))
		if err := json.Unmarshal(element, &sup); err != nil {
			return load
		}
		arrSup = append(arrSup, *sup)
	}
	const minOneTicket = 3.33
	var openTickets int
	for _, v := range arrSup {
		openTickets += v.ActiveTickers
	}
	waitingTime := (minOneTicket * float64(openTickets)) / 7
	switch {
	case waitingTime < 9:
		load = append(load, 1)
	case 9 <= waitingTime && waitingTime <= 16:
		load = append(load, 2)
	default:
		load = append(load, 3)
	}
	load = append(load, int(math.Trunc(waitingTime)))
	return load
}
