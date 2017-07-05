package flags

import (
	"flag"
)

var Port *string
var Server *string

func Parse() {
	Port = flag.String("port", "1025", "list port")
	Server = flag.String("server", "192.168.12.72", "server address")
}
