package main

import (
	_ "fabric-orderer-seek/server/mysql"
	"fabric-orderer-seek/server/ote"
)

func main() {
	ote.StartOTE()
}
