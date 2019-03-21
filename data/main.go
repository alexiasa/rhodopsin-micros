package main

import "time"

func main() {
	on := true
	if on == true {
		pullData()
		time.Sleep(1 * time.Minute)
		//get ips with missing details
		//enrichData()
		//updateDB()
	}
}

