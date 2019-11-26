package main

import (
	"flag"
	"fmt"
	"github.com/hackmdio/portchecker/internal"
	"log"
	"net"
	"os"
	"time"
)

var conStr string
var conStrFromEnv string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.IntVar(&internal.DefaultRetryOptions.Attempt, "Attempt", 5, "retry Attempt")
	flag.Int64Var(&internal.DefaultRetryOptions.WaitTime, "wait", 3, "if connecting fail, wait how many second to retry")
	flag.StringVar(&conStr, "constr", "", "try to connect")
	flag.StringVar(&conStrFromEnv, "env", "", "environment with connection string")
}

func main() {
	flag.Parse()

	// conStr or conStrFromEnv should only set one
	if (conStr == "") == (conStrFromEnv == "") {
		log.Fatalln("Error: must specific one type of conStr")
	}

	netPort := getNetPort()
	if netPort == nil {
		log.Fatalln("Error: cannot parsed connection string.")
	}

	attemptTime := internal.DefaultRetryOptions.Attempt
	waitTime := internal.DefaultRetryOptions.WaitTime
	for i := 0; i < attemptTime; i++ {
		c, err := net.Dial(netPort.Protocol, netPort.GetNetworkAddress())
		if err != nil {
			log.Printf("%v\n    wait %d seconds retry", err, waitTime)
			time.Sleep(time.Duration(waitTime) * time.Second)
			continue
		}
		_ = c.Close()
		log.Printf("dial %s %s: connect: connection success\n", netPort.Protocol, netPort.GetNetworkAddress())
		os.Exit(0)
	}
	log.Fatalf("Exceeded maximum retry attempts (%d), connot connect to %s\n", attemptTime, netPort.GetNetworkAddress())
}

func getNetPort() *internal.NetPort {
	if conStr != "" {
		return internal.ParseNetworkStringToNetPort(conStr)
	}
	if conStrFromEnv != "" {
		envValue := os.Getenv(conStrFromEnv)
		if envValue == "" {
			fmt.Printf("Error: value of env: %s is empty!", conStrFromEnv)
			os.Exit(1)
		}
		return internal.ParseNetworkStringToNetPort(envValue)
	}
	return nil
}
