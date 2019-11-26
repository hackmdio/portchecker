package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type NetPort struct {
	Protocol string
	Address  string
	Port     string
}

func NewNetPort(proto string, address string, port string) *NetPort {
	return &NetPort{
		Protocol: proto,
		Address:  address,
		Port:     port,
	}
}

func (n *NetPort) GetNetworkAddress() string {
	if n.Port == "" {
		return n.Address
	}
	return fmt.Sprintf("%s:%s", n.Address, n.Port)
}

func ParseNetworkStringToNetPort(inp string) *NetPort {
	result := ParseNetworkString(inp)
	return NewNetPort(result[1], result[2], result[3])
}

func ParseNetworkString(inp string) []string {
	reg, err := regexp.Compile(`^(?:([a-z]+)://)?(?:.*(?::.*)?@)?([^:/]*)(?::(\d+))?.*$`)
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
