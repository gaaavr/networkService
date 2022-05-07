package data

import (
	"bufio"
	"networkService/pkg/service"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func GetResultsEmail() {
	defer wg.Done()
	var email EmailData
	var emailArr []EmailData
	providers := make(map[string][][]EmailData)
	countriesList := service.GetCountriesList()
	providersList := service.GetEmailProvidersList()
	file, err := os.Open("..\\networkService\\simulator\\email.data")
	if err != nil {
		Result.Lock()
		Result.Email = map[string][][]EmailData{}
		Result.Unlock()
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strData := fileScanner.Text()
		if strings.Count(strData, ";") != 2 {
			continue
		}
		element := strings.Split(strData, ";")
		if countriesList[element[0]] && providersList[element[1]] {
			if deliveryTime, err := strconv.Atoi(element[2]); err != nil {
				continue
			} else {
				email.Country = element[0]
				email.Provider = element[1]
				email.DeliveryTime = deliveryTime
				emailArr = append(emailArr, email)
			}
		}
	}
	if emailArr == nil {
		Result.Lock()
		Result.Email = providers
		Result.Unlock()
		return
	}
	countries := []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
	for _, v := range countries {
		speedProviders := make([]EmailData, 0)
		for i, _ := range emailArr {
			if emailArr[i].Country == v {
				speedProviders = append(speedProviders, emailArr[i])
				continue
			} else {
				emailArr = emailArr[i+1:]
				break
			}
		}
		sort.SliceStable(speedProviders, func(i, j int) bool {
			return speedProviders[i].DeliveryTime < speedProviders[j].DeliveryTime
		})
		providers[v] = append(providers[v], speedProviders[:3])
		providers[v] = append(providers[v], speedProviders[len(speedProviders)-3:])
	}
	Result.Lock()
	Result.Email = providers
	Result.Unlock()
	return
}
