package main

import (
	"fmt"
	"sync"
)


func bootstrap() {
	capitals = make(map[string]string)
	capitals["France"] = "Paris"
	capitals["Germany"] = "Berlin"
	capitals["Japan"] = "Tokyo"
	capitals["Brazil"] = "Brasilia"
	capitals["China"] = "Beijing"
	capitals["USA"] = "Washineton"
	capitals["Poland"] = "Warsaw"
}

/*
var capitals map[string]string
func getCapitalCity(country string) string {
	if capitals == nil {
		bootstrap()
	}
	return capitals[country]
}
*/

/*
var (
	capitals map[string]string
	mutex sync.Mutex
)

func getCapitalCity(country string) string {
	mutex.Lock()
	if capitals == nil {
		bootstrap()
	}
	mutex.Unlock()
	return capitals[country]
}
*/

/*
var (
	capitals map[string]string
	mutex sync.RWMutex
)

func getCapitalCity(country string) string {
	mutex.RLock()
	if capitals != nil {
		city := capitals[country]
		mutex.RUnlock()
		return city
	}
	mutex.RUnlock()
	mutex.Lock()
	if capitals == nil {
		bootstrap()
	}
	mutex.Unlock()
	return getCapitalCity(country)
}
*/
var (
	capitals map[string]string
	once sync.Once
)

func getCapitalCity(country string) string {
	once.Do(bootstrap)
	return capitals[country]
}

func main() {
	fmt.Println(getCapitalCity("Poland"))
	fmt.Println(getCapitalCity("USA"))
	fmt.Println(getCapitalCity("Japan"))
}
