package actions

import (
	"fmt"
	"net"

	"github.com/urfave/cli/v2"
)

func NameServers(c *cli.Context) error {
	ns, err := net.LookupNS(c.String("host"))
	if err != nil {
		return err
	}

	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}
	return nil
}

func IPAddress(c *cli.Context) error {
	ip, err := net.LookupIP(c.String("host"))
	if err != nil {
		return err
	}
	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}
	return nil
}

func CNAME(c *cli.Context) error {
	cname, err := net.LookupCNAME(c.String("host"))
	if err != nil {
		return err
	}
	fmt.Println(cname)
	return nil
}

func MXRecords(c *cli.Context) error {
	mx, err := net.LookupMX(c.String("host"))
	if err != nil {
		return err
	}
	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i].Host, mx[i].Pref)
	}
	return nil
}
