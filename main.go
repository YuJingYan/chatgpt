package main

import "chatgpt/routers"

func main() {

	r := routers.SetupRouter()

	r.Run(":3001")
}
