package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

type RetryOptions struct {
	waitTime int64
	attempt  int
}

type NetPort struct {
	protocol string
	address  string
	port     string
}

type SchemaPortMapping struct {
	schema            string
	transportProtocol string
	port              int
}

var SchemaPortTable = map[string]*SchemaPortMapping{
	"http":     {"http", "tcp", 80},
	"https":    {"https", "tcp", 443},
	"postgres": {"postgres", "tcp", 5432},
	"mysql":    {"mysql", "tcp", 3306},
	"mariadb":  {"mariadb", "tcp", 3306},
	"redis":    {"redis", "tcp", 6379},
	"mssql":    {"mssql", "tcp", 1433},
	"ftp":      {"ftp", "tcp", 21},
	"":         {"tcp", "tcp", 80},
	"tcp":      {"tcp", "tcp", 80},
	"udp":      {"tcp", "tcp", 80},
}

func (n *NetPort) GetNetworkAddress() string {
	if n.port == "" {
		return n.address
	}
	return fmt.Sprintf("%s:%s", n.address, n.port)
}

func ParseNetworkString(inp string) []string {
	reg, err := regexp.Compile(`^(?:([a-z]+):\/\/)?(?:(?:\w+(?::\w+)?@)?([A-Za-z0-9.]+))(?::([0-9]+))?$`)
	if err != nil {
		fmt.Printf("%s (%s), error: %s", "Cannot parse network stirng", inp, err)
		os.Exit(1)
	}

	result := reg.FindStringSubmatch(inp)

	if len(result) != 4 {
		result = []string{"", "", "", ""}
	}

	if mapping := SchemaPortTable[result[1]]; mapping != nil {
		result[1] = mapping.transportProtocol
		if result[2] == "" {
			result[2] = "localhost"
		}
		if result[3] == "" {
			result[3] = strconv.Itoa(mapping.port)
		}
	}

	return result
}

func ParseNetworkStringToNetPort(inp string) *NetPort {
	result := ParseNetworkString(inp)

	var ret = &NetPort{
		protocol: result[1],
		address:  result[2],
		port:     result[3],
	}

	return ret
}

func makeDefaultRetryOptions() *RetryOptions {
	return &RetryOptions{
		attempt:  5,
		waitTime: 3,
	}
}

var retryOptions = makeDefaultRetryOptions()
var conStr string

func init() {
	flag.IntVar(&retryOptions.attempt, "attempt", 5, "retry attempt")
	flag.Int64Var(&retryOptions.waitTime, "wait", 3, "if connecting fail, wait how many second to retry")
	flag.StringVar(&conStr, "constr", "", "try to connect")
}

func main() {
	flag.Parse()

	if conStr == "" {
		fmt.Println("Error: must specific constr")
		os.Exit(1)
	}

	netPort := ParseNetworkStringToNetPort(conStr)

	for i := 0; i < retryOptions.attempt; i++ {
		c, err := net.Dial(netPort.protocol, netPort.GetNetworkAddress())
		if err != nil {
			fmt.Printf("%v\n", err)

			time.Sleep(time.Duration(retryOptions.waitTime) * time.Second)
			continue
		}
		_ = c.Close()
		fmt.Printf("dial %s %s: connect: connection success\n", netPort.protocol, netPort.GetNetworkAddress())
		os.Exit(0)
	}
	os.Exit(1)
}
