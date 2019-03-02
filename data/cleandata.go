package main

import (
	"io/ioutil"
	"regexp"
)

// saveToFile takes a string representing a filename and a string of data and writes the data to a file
// The data is converted into a slice of bytes
// The file is given permissions 0666  for r/w
func saveToFile(filename string, tweets string) error {
	return ioutil.WriteFile(filename, []byte(tweets), 0666)
}

// extractAddrs takes a string of data and extracts IPv4 addresses from it
// The first regex is for a standard address in format xx.xx.xx.xx
// The second regex is for a "sanitized" address in format xx[.]xx[.]xx[.]xx
// It returns a slice of strings consisting of all the matched addresses
func extractAddrs(tweetblock string) []string {
	addrs := []string{""}
	var re = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)

	for _, match := range re.FindAllString(tweetblock, -1) {
		addrs = append(addrs, match)
		// fmt.Println(match, "found at index", i)
	}
	re = regexp.MustCompile(`(?m)\b(?:[0-9]{1,3}\[\.\]){3}[0-9]{1,3}\b`)
	var ex = regexp.MustCompile(`\[.]`)
	for _, match := range re.FindAllString(tweetblock, -1) {
		match = ex.ReplaceAllString(match,".")
		addrs = append(addrs, match)
		// fmt.Println(match, "found at index", i)
	}

	return addrs
}



// will need logic to extract addresses from tweets

// func removePrivateAddrs()


// func recordAddrs()
// will need to record unique addrs for analysis