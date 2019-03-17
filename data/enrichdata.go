package main

import (
	"encoding/json"
	"fmt"
	"github.com/williballenthin/govt"
	"os"
)
// adapted from sample govt client found here: https://github.com/williballenthin/govt/blob/master/SampleClients/ipreport/vtIpReport.go


var apikey string = os.Getenv("VT_API_KEY")
var apiurl string = os.Getenv("VT_API_URL")
//var ip string

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getReport (ip string) {
	if ip == "" {
		fmt.Println("-ip=<ip> missing!")
		os.Exit(1)
	}
	c, err := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	check(err)
	r, err := c.GetIpReport(ip)
	check(err)
	j, err := json.MarshalIndent(r, "", "    ")
	check(err)
	fmt.Printf("IP Report: ")
	os.Stdout.Write(j)

}

func setDetails () {

}
