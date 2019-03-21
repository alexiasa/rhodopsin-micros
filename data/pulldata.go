package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/alexiasa/rhodopsin-micros/ips/controllers"
	"github.com/alexiasa/rhodopsin-micros/ips/models"
	"github.com/williballenthin/govt"
	"log"
	"net/url"
	"os"
	models2 "rhodopsin-micros/ips/models"
)


func outputTweets(status anaconda.SearchResponse) {
	file, err := os.Create("tweets.json")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	fmt.Fprint(file, status.Statuses)
}

func pullData() {

	var ipaddrs []string

	f, err := os.Create("addrs.json")
	if err != nil {
		print("couldn't create file")
	}
	defer f.Close()

	anaconda.SetConsumerKey(os.Getenv("KEY"))
	anaconda.SetConsumerSecret(os.Getenv("SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TOKEN"), os.Getenv("TOKENSECRET"))

	v := url.Values{}
	v.Set("count", "50")
	// v.Set("result_type", "popular")
	searchResult, _ := api.GetSearch("%23malware+OR+suspicious+OR+malware+AND+%22IP%20address%22+OR+%22IP%20addresses%22+OR+IP", v)
	outputTweets(searchResult)
	println("sample tweet:\r\n", searchResult.Statuses[0].FullText)

	for _, tweet := range searchResult.Statuses {

		ipaddrs = extractAddrs(tweet.Text)
		for _, addr := range ipaddrs {
			if addr != "" {
				// try to get info from vt about it
				i, m, a, l := getReport(addr)
				// create document for ip based on what was returned
				createDocument(i, m, a, l)
				addr = "{'ipaddr': '" + addr + "'},"
				_, err := f.WriteString(addr)
				if err != nil {
					println("couldn't write IP to file")
				}
			}
		}
	}


}

func createDocument(ipaddr string, mal bool, asn string, location string) {
	context := controllers.NewContext()
	defer context.Close()
	col := context.DbCollection("ips")

	// Insert Data
	// err := col.Insert(&models.IpDetails{Ipaddr: ipaddr, Asn: asn, Location: location, })
	err := col.Insert(&models.IpDetails{Ipaddr: ipaddr, Asn: asn, Location: location, Malicious: mal })
	if err != nil {
		panic(err)
	}

}

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// accept an IP addr string. return the IP, malicious boolean value (whether or not the address was found in VT), the asn,
// and the inferred location as a string object.

func getReport (ip string) (string, bool, string, string)  {

	var apikey string = os.Getenv("VT_API_KEY")
	var apiurl string = os.Getenv("VT_API_URL")
	var vmal bool
	var vloc string
	var vasn string

	if ip == "" {
		log.Fatalf("there was no IP")
		return "0", false, "0", "0"
	}
	c, err := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	check(err)
	r, err := c.GetIpReport(ip)
	if r.Status.ResponseCode == '0' {
		return "0", false, "0", "0"
	}
	check(err)

	j, err := json.MarshalIndent(r, "", "    ")

	var dat map[string]interface{}

	// example :  num := dat["num"].(float64)

	if err := json.Unmarshal(j, govt.IpReport{}) ; err != nil {
		panic(err)
		return "0", false, "0", "0"

	} else {

		if dat["response_code"].(int) == 0 {
			vmal = false
		} else {
			vmal = true
		}

		vasn = dat["asn"].(string)
		vloc = dat["country"].(string)

		os.Stdout.Write(j)
		return ip, vmal, vasn, vloc
	}


}