package config

import (
	"all_inclusive_app/models"
	"encoding/json"
	"io/ioutil"
	"log"
)

var Configure models.App
var ConfigDB models.DBConnection
var ConfigureLogger models.LoggerStruct

func ReadConfigApp(F string) {
	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	err = json.Unmarshal(byteValue, &Configure)
	log.Println(Configure)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

}

func ReadConfigDB(F string) {
	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	err = json.Unmarshal(byteValue, &ConfigDB)
	log.Println(ConfigDB)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

}

func ReadConfigLogger(F string) {
	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	err = json.Unmarshal(byteValue, &ConfigureLogger)
	log.Println(ConfigureLogger)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

}
