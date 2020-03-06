package main

import (
	"flag"
	"fmt"
	"github.com/hackmdio/portchecker/internal"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var (
	conStr         string
	conStrFromEnv  string
	conStrFromFile string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.IntVar(&internal.DefaultRetryOptions.Attempt, "attempt", 5, "retry Attempt")
	flag.Int64Var(&internal.DefaultRetryOptions.WaitTime, "wait", 3, "if connecting fail, wait how many second to retry")
	flag.StringVar(&conStr, "constr", "", "try to connect")
	flag.StringVar(&conStrFromEnv, "env", "", "environment with connection string")
	flag.StringVar(&conStrFromFile, "file", "", "file path that contains connection string")
}

func main() {
	flag.Parse()

	// conStr or conStrFromEnv should only set one
	if (conStr == "") && (conStrFromEnv == "") && (conStrFromFile == "") {
		log.Fatalln("Error: must specific one type of conStr")
	}

	netPort := getNetPort()
	if netPort == nil {
		log.Fatalln("Error: cannot parsed connection string.")
	}

	attemptTime := internal.DefaultRetryOptions.Attempt
	waitTime := internal.DefaultRetryOptions.WaitTime

	fmt.Printf("Info: try to connect to %s://%s in port %s\n", netPort.Protocol, netPort.Address, netPort.Port)

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
	if conStrFromFile != "" {
		contents, err := ioutil.ReadFile(conStrFromFile)
		if err != nil {
			fmt.Printf("Error: can't read file %s", conStrFromFile)
			os.Exit(1)
		}
		trimConnStr := strings.TrimSpace(string(contents))
		return internal.ParseNetworkStringToNetPort(trimConnStr)
	}
	return nil
}
