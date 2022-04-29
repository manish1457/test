package main

import (
	pkg "example/api-call/pkg/routes"
)

func main() {
	pkg.InitDB()
	pkg.Run()

}
