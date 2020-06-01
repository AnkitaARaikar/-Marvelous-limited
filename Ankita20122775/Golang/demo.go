package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type marvelous struct {
	Name      string `json:"name"`
	Character []struct {
		Name     string `json:"name"`
		MaxPower int    `json:"max_power"`
	} `json:"character"`
}

func main() {

	url := "http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b"
	url1 := "http://www.mocky.io/v2/5ed49b033300005400f7a25f"
	url2 := "http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req1, err1 := http.NewRequest(http.MethodGet, url1, nil)
	if err1 != nil {
		log.Fatal(err1)
	}

	req2, err2 := http.NewRequest(http.MethodGet, url2, nil)
	if err2 != nil {
		log.Fatal(err2)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	res1, getErr1 := spaceClient.Do(req1)
	if getErr1 != nil {
		log.Fatal(getErr1)
	}

	res2, getErr2 := spaceClient.Do(req2)
	if getErr2 != nil {
		log.Fatal(getErr2)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	body1, readErr1 := ioutil.ReadAll(res1.Body)
	if readErr1 != nil {
		log.Fatal(readErr1)
	}

	body2, readErr2 := ioutil.ReadAll(res2.Body)
	if readErr2 != nil {
		log.Fatal(readErr2)
	}

	marv1 := marvelous{}
	jsonErr := json.Unmarshal(body, &marv1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(marv1.Name)
	for _, i := range marv1.Character {
		fmt.Println("Name:" + i.Name + " , " + "Max_power:" + strconv.Itoa(i.MaxPower))
	}

	marv2 := marvelous{}
	jsonErr1 := json.Unmarshal(body1, &marv2)
	if jsonErr1 != nil {
		log.Fatal(jsonErr1)
	}
	fmt.Println(marv2.Name)
	for _, i := range marv2.Character {
		fmt.Println("Name:" + i.Name + " , " + "Max_power:" + strconv.Itoa(i.MaxPower))
	}

	marv3 := marvelous{}
	jsonErr2 := json.Unmarshal(body2, &marv3)
	if jsonErr2 != nil {
		log.Fatal(jsonErr2)
	}
	fmt.Println(marv3.Name)
	for _, i := range marv3.Character {
		fmt.Println("Name:" + i.Name + " , " + "Max_power:" + strconv.Itoa(i.MaxPower))
	}
}
