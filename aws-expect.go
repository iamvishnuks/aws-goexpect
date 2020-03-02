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
	addr = flag.String("address", "", "address of telnet server")
	user = flag.String("user", "", "username to use")
	pass = flag.String("pass", "", "password to use")
	cmd  = flag.String("cmd", "", "command to run")

	userRE   = regexp.MustCompile("username:")
	passRE   = regexp.MustCompile("password:")
	promptRE = regexp.MustCompile("%")
)

func main() {
	flag.Parse()
	fmt.Println(term.Bluef("Telnet 1 example"))

	e, _, err := expect.Spawn(fmt.Sprintf("telnet %s", *addr), -1)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()

	e.Expect(userRE, timeout)
	e.Send(*user + "\n")
	e.Expect(passRE, timeout)
	e.Send(*pass + "\n")
	e.Expect(promptRE, timeout)
	e.Send(*cmd + "\n")
	result, _, _ := e.Expect(promptRE, timeout)
	e.Send("exit\n")

	fmt.Println(term.Greenf("%s: result: %s\n", *cmd, result))
}