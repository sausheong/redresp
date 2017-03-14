package main

import (
	r "github.com/sausheong/red/responder/core"
)

func main() {
	r.ROUTEID = "GET/_/hello/world"
	r.RESPOND = respond
	r.Run()
}

func respond(request r.RequestInfo, response *r.ResponseInfo) {
	response.SetHTML()
	response.Body = "Hello World"
}
