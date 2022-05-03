package data

import (
	"bufio"
	"networkService/pkg/service"
	"os"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func GetResultsSMS() {
	var sms SMSData
	var smsArr []SMSData
	var finalData [][]SMSData
	countriesList := service.GetCountriesList()
	providersList := service.GetSMSProvidersList()
	file, err := os.Open("..\\networkService\\sms.data")
	if err != nil {
		Result.SMS = finalData
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strData := fileScanner.Text()
		if strings.Count(strData, ";") != 3 {
			continue
		}
		element := strings.Split(strData, ";")
		if countriesList[element[0]] && providersList[element[3]] {
			sms.Country = element[0]
			sms.Bandwidth = element[1]
			sms.ResponseTime = element[2]
			sms.Provider = element[3]
			smsArr = append(smsArr, sms)
		}
	}
	countries := service.GetFullNamesCountries()
	for i, _ := range smsArr {
		smsArr[i].Country = countries[smsArr[i].Country]
	}
	sort.SliceStable(smsArr, func(i, j int) bool {
		return smsArr[i].Provider < smsArr[j].Provider
	})
	providerSorted := make([]SMSData, len(smsArr))
	copy(providerSorted, smsArr)
	finalData = append(finalData, providerSorted)
	sort.SliceStable(smsArr, func(i, j int) bool {
		return smsArr[i].Country < smsArr[j].Country
	})
	finalData = append(finalData, smsArr)
	Result.SMS = finalData
	return
}
