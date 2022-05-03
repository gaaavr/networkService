package data

import (
	"bufio"
	"networkService/pkg/service"
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func GetResultsVC() {
	var voiceCall VoiceCallData
	var voiceCallArr []VoiceCallData
	countriesList := service.GetCountriesList()
	providersList := service.GetVoiceCallProvidersList()
	file, err := os.Open("..\\networkService\\voice.data")
	if err != nil {
		Result.VoiceCall = voiceCallArr
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strData := fileScanner.Text()
		if strings.Count(strData, ";") != 7 {
			continue
		}
		element := strings.Split(strData, ";")
		if countriesList[element[0]] && providersList[element[3]] {
			if connectionParse, err := strconv.ParseFloat(element[4], 32); err != nil {
				continue
			} else {
				voiceCall.ConnectionStability = float32(connectionParse)
			}
			if TTFBParse, err := strconv.Atoi(element[5]); err != nil {
				continue
			} else {
				voiceCall.TTFB = TTFBParse
			}
			if VoicePurityParse, err := strconv.Atoi(element[6]); err != nil {
				continue
			} else {
				voiceCall.VoicePurity = VoicePurityParse
			}
			if MedianOfCallsTimeParse, err := strconv.Atoi(element[7]); err != nil {
				continue
			} else {
				voiceCall.MedianOfCallsTime = MedianOfCallsTimeParse
			}
			voiceCall.Country = element[0]
			voiceCall.Bandwidth = element[1]
			voiceCall.ResponseTime = element[2]
			voiceCall.Provider = element[3]
			voiceCallArr = append(voiceCallArr, voiceCall)
		}
	}
	Result.VoiceCall = voiceCallArr
	return
}
