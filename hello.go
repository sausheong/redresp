package main

import (
	r "github.com/sausheong/red/responder/core"
)

func main() {
	r.Run(respond)
}

func respond(request r.RequestInfo, response *r.ResponseInfo) {
	response.SetHTML()
	response.Body = "Hello World  "
}
