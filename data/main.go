package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"os"
)


func outputTweets(status anaconda.SearchResponse) {
	file, err := os.Create("tweets.json")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	fmt.Fprint(file, status.Statuses)
}

func main() {
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
	v.Set("count", "100")
	//v.Set("result_type", "popular")
	searchResult, _ := api.GetSearch("%23malware+OR+suspicious+OR+malware+AND+%22IP%20address%22+OR+%22IP%20addresses%22+OR+IP", v)
	outputTweets(searchResult)
	println("sample tweet:\r\n", searchResult.Statuses[0].FullText)


	for _, tweet := range searchResult.Statuses {

		ipaddrs = extractAddrs(tweet.Text)
		for _, addr := range ipaddrs {
			if addr != "" {
				addr = "{'ipaddr': '" + addr + "'},"
				_, err := f.WriteString(addr)
				if err != nil {
					println("couldn't write IP to file")
				}
			}
		}

	}

}