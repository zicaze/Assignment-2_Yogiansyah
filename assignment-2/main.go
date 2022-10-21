package main

import "assignment-2/routers"

func main() {
	var PORT = ":8080"
	routers.ServerOn().Run(PORT)
}
