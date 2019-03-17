package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/alexiasa/rhodopsin-micros/ips/controllers"
	"github.com/alexiasa/rhodopsin-micros/ips/models"
	"log"
	"net/url"
	"os"
	"os/exec"
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
	v.Set("count", "20")
	// v.Set("result_type", "popular")
	searchResult, _ := api.GetSearch("%23malware+OR+suspicious+OR+malware+AND+%22IP%20address%22+OR+%22IP%20addresses%22+OR+IP", v)
	outputTweets(searchResult)
	println("sample tweet:\r\n", searchResult.Statuses[0].FullText)

	for _, tweet := range searchResult.Statuses {

		ipaddrs = extractAddrs(tweet.Text)
		for _, addr := range ipaddrs {
			if addr != "" {
				createDocument(addr)
				addr = "{'ipaddr': '" + addr + "'},"
				_, err := f.WriteString(addr)
				if err != nil {
					println("couldn't write IP to file")
				}
			}
		}
	}
}

func createDocument(ipaddr string) {
	context := controllers.NewContext()
	defer context.Close()
	col := context.DbCollection("ips")

	//VT Lookup
	getReport(ipaddr)

	// Insert Data
	err := col.Insert(&models.IpDetails{Ipaddr: ipaddr})

	if err != nil {
		panic(err)
	}


}

func loadJSON(filename string) {

	user := os.Getenv("MONGOUSER")
	db := os.Getenv("DBNAME")
	host := os.Getenv("MONGOHOST")
	pass := os.Getenv("MONGOPASS")


	cmd := exec.Command("sh", "-c", "/usr/bin/mongoimport " + "--host " + host + " --port" + " 27017 " + "--db " + db + " --collection" + " ips" + " --authenticationDatabase" +
		" admin" + " --username " + user + " --password " + pass + " --jsonArray" + " --file " + filename)
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
		}
// mongoimport --db test --collection inventory ^
//          --authenticationDatabase admin --username <user> --password <password> ^
//          --drop --file ~\downloads\inventory.crud.json