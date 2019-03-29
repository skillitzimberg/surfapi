package datautil

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetBouyData() []byte {
	url := "https://www.ndbc.noaa.gov/data/realtime2/46029.spec"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func removeEmptySpace(arrayWithSpaces []string) []string {
	cleanedData := make([]string, 0, len(arrayWithSpaces))
	for j := 0; j < len(arrayWithSpaces); j++ {
		if arrayWithSpaces[j] != "" {
			cleanedData = append(cleanedData, arrayWithSpaces[j])
		}
	}
	return cleanedData
}

func HandleRawData(inputText []byte) [][]string {
	text := string(inputText)
	stringsOnReturn := strings.Split(text, "\n")
	var stringsOnTabs []string
	cleanedData := make([]string, 0, len(stringsOnTabs))
	dataRows := make([][]string, 0, len(stringsOnReturn))

	for i := 2; i < 24; i++ {
		stringsOnTabs = strings.Split(stringsOnReturn[i], " ")
		cleanedData = removeEmptySpace(stringsOnTabs)
		dataRows = append(dataRows, cleanedData)
	}

	return dataRows
}

func DataToStructs(dataRows [][]string) []SurfData {
	var surfDataStructs []SurfData

	for i := 0; i < len(dataRows); i++ {
		hourlyData := dataRows[i]
		for j := 0; j < len(hourlyData); j++ {
		}

		year, _ := strconv.Atoi(hourlyData[0])
		month, _ := strconv.Atoi(hourlyData[1])
		day, _ := strconv.Atoi(hourlyData[2])
		hour, _ := strconv.Atoi(hourlyData[3])
		minute, _ := strconv.Atoi(hourlyData[4])
		WVHT, _ := strconv.ParseFloat(hourlyData[5], 64)
		SwH, _ := strconv.ParseFloat(hourlyData[6], 64)
		SwP, _ := strconv.ParseFloat(hourlyData[7], 64)
		WWH, _ := strconv.ParseFloat(hourlyData[8], 64)
		WWP, _ := strconv.ParseFloat(hourlyData[9], 64)
		SwD := hourlyData[10]
		WWD := hourlyData[11]
		steepness := hourlyData[12]
		APD, _ := strconv.ParseFloat(hourlyData[13], 64)
		MWD, _ := strconv.Atoi(hourlyData[14])

		var surfData = SurfData{
			year,
			month,
			day,
			hour,
			minute,
			WVHT,
			SwH,
			SwP,
			WWH,
			WWP,
			SwD,
			WWD,
			steepness,
			APD,
			MWD,
		}

		surfDataStructs = append(surfDataStructs, surfData)
	}
	return surfDataStructs
}

func GetSwellHeight(cleanedData []string) float64 {
	convertSwHToInt, err := strconv.ParseFloat(cleanedData[6], 64)
	if err != nil {
		log.Fatal(err)
	}
	return convertSwHToInt
}

func GetSwellPeriod(cleanedData []string) float64 {
	convertSwPToInt, err := strconv.ParseFloat(cleanedData[7], 64)
	if err != nil {
		log.Fatal(err)
	}
	return convertSwPToInt
}

func GetWindDirection(cleanedData []string) string {
	return cleanedData[11]
}
