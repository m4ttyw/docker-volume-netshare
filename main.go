package main

import (
	"github.com/m4ttyw/docker-volume-netshare/netshare"
)

var VERSION string = ""
var BUILD_DATE string = ""

func main() {
	netshare.Version = VERSION
	netshare.BuildDate = BUILD_DATE
	netshare.Execute()
}
