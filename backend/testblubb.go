package main

import (
	"log"
	"reflect"
)

func SagMirWasDuBist(wert any) {
	t := reflect.ValueOf(wert).Type()
	log.Println(t)
	if t.Kind() == reflect.Struct {
		for _, v := range reflect.VisibleFields(t) {
			log.Printf("GORM TAG FOR %s: %s", v.Name, v.Tag.Get("json"))
		}
	}
}

func init() {
	/*
		SagMirWasDuBist(6)
		SagMirWasDuBist("HI")
		SagMirWasDuBist(models.Streamer{})
		SagMirWasDuBist(twitch.StreamInformation{})
		SagMirWasDuBist(fmt.Errorf("Kein Plan"))
	*/
}
