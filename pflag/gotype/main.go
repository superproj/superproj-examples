package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var (
	host string
	port int
)

func main() {
	pflag.StringVarP(&host, "host", "h", "127.0.0.1", "MySQL service host address.")
	pflag.IntVarP(&port, "port", "p", 3306, "MySQL service host port.")
	pflag.Parse()

	fmt.Printf("host: %s, port: %d\n", host, port)
}
