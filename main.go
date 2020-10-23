package main

import (
	"flag"
	"fmt"
	"regexp"
)

func main() {

	flag.Parse()
	adress := flag.Args()

	bitcoinadress := adress[0]
	isValidAddress := validateAddress(bitcoinadress)
	if !isValidAddress {
		fmt.Println("⚠️Invalid btcAddress:" + bitcoinadress)
	} else {
		fmt.Println("👍Valid btcAddress:" + bitcoinadress)
	}

}

func validateAddress(address string) bool {

	re := regexp.MustCompile("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$")
	return re.MatchString(address)
}
