package service

import "strings"

func GetEmailProvidersList() map[string]bool {
	providersList := make(map[string]bool)
	arr := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX",
		"Protonmail", "Yandex", "Mail.ru"}
	for _, v := range arr {
		providersList[v] = true
	}
	return providersList
}

func GetFullNamesCountries() map[string]string {
	return map[string]string{
		"RU": "Russian Federation",
		"US": "United States",
		"GB": "United Kingdom",
		"FR": "France",
		"BL": "Saint Barthélemy",
		"AT": "Austria",
		"BG": "Bulgaria",
		"DK": "Denmark",
		"CA": "Canada",
		"ES": "Spain",
		"CH": "China",
		"TR": "Turkey",
		"PE": "Peru",
		"NZ": "New Zealand",
		"MC": "Monaco",
	}
}

func GetCountriesList() map[string]bool {
	countriesList := make(map[string]bool)
	arr := []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
	for _, v := range arr {
		countriesList[v] = true
	}
	return countriesList
}

func GetSMSProvidersList() map[string]bool {
	providersList := make(map[string]bool)
	arr := []string{"Topolo", "Rond", "Kildy"}
	for _, v := range arr {
		providersList[v] = true
	}
	return providersList
}

func GetVoiceCallProvidersList() map[string]bool {
	providersList := make(map[string]bool)
	arr := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	for _, v := range arr {
		providersList[v] = true
	}
	return providersList
}

func FormattingString(s string) []string {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	arrData := strings.SplitAfter(s, "}")
	arrData = arrData[:len(arrData)-1]
	return arrData
}
