package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"networkService/pkg/service"
	"sort"
	"strings"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func GetResultsMMS() [][]MMSData {
	var mms *MMSData
	var arrMMS []MMSData
	var finalData [][]MMSData
	countriesList := service.GetCountriesList()
	providersList := service.GetSMSProvidersList()
	content, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return finalData
	}
	if content.StatusCode == 500 {
		return finalData
	}
	data, err := io.ReadAll(content.Body)
	if err != nil {
		return finalData
	}
	defer content.Body.Close()
	strData := string(data)
	arrData := service.FormattingString(strData)
	var element []byte
	for i, _ := range arrData {
		element = []byte(strings.Trim(arrData[i], " ,"))
		if err := json.Unmarshal(element, &mms); err != nil {
			fmt.Println(finalData)
			return [][]MMSData{}
		}
		if countriesList[mms.Country] && providersList[mms.Provider] {
			arrMMS = append(arrMMS, *mms)
		}
	}
	countries := service.GetFullNamesCountries()
	for i, _ := range arrMMS {
		arrMMS[i].Country = countries[arrMMS[i].Country]
	}
	sort.SliceStable(arrMMS, func(i, j int) bool {
		return arrMMS[i].Provider < arrMMS[j].Provider
	})
	providerSorted := make([]MMSData, len(arrMMS))
	copy(providerSorted, arrMMS)
	finalData = append(finalData, providerSorted)
	sort.SliceStable(arrMMS, func(i, j int) bool {
		return arrMMS[i].Country < arrMMS[j].Country
	})
	finalData = append(finalData, arrMMS)
	return finalData

}
