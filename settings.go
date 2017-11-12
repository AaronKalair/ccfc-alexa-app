package main

import (
	json "encoding/json"
	io "io/ioutil"
)

type Settings struct {
	AppId            string
	DatabaseLocation string
	IcalLocation     string
}

func GetSettings() Settings {
	file, err := io.ReadFile("settings.json")
	if err != nil {
		panic(err)
	}

	var AppSettings Settings
	err2 := json.Unmarshal(file, &AppSettings)

	if err2 != nil {
		panic(err)
	}
	return AppSettings
}
