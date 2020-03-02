package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
	"github.com/google/goterm/term"
)

const (
	timeout = 10 * time.Minute
)

var (
	accesskeyid = flag.String("accessid", "ddfd", "address of telnet server")
	secret = flag.String("accesskey", "fdfd", "username to use")
	profile = flag.String("profile", "", "password to use")
	region = flag.String("region","ff-ff-f","region for the cluster")
	output = flag.String("output","json","output format")

	accesskeyidRE   = regexp.MustCompile("AWS Access Key ID:")
	secretRE   = regexp.MustCompile("AWS Secret Access Key:")
	regionRE = regexp.MustCompile("Default region name:")
	outputRE = regexp.MustCompile("Default output format:")
	promptRE = regexp.MustCompile("%")
)

func main() {
	flag.Parse()
	fmt.Println(term.Bluef("AWS CLI configure"))

	e, _, err := expect.Spawn(fmt.Sprintf("aws configure --profile", *profile), -1)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()

	e.Expect(accesskeyidRE, timeout)
	e.Send(*accesskeyid + "\n")
	e.Expect(secretRE, timeout)
	e.Send(*secret + "\n")
	e.Expect(regionRE, timeout)
	e.Send(*region + "\n")
	e.Expect(outputRE, timeout)
	e.Send(*output + "\n")
	result, _, _ := e.Expect(promptRE, timeout)
	e.Send("exit\n")

	fmt.Println(term.Greenf("result: %s\n", result))
}