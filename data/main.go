package main

import "time"

func main() {
	on := true
	if on == true {
		pullData()
		time.Sleep(1 * time.Hour)
		//get ips with missing details
		//enrichData()
		//updateDB()
	}
}

