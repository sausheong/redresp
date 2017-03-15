package main

import (
	r "github.com/sausheong/red/responder/core"
)

func main() {
	r.ROUTEID = "GET/_/hola/mundo"
	r.RESPOND = respond
	r.Run()
}

func respond(request r.RequestInfo, response *r.ResponseInfo) {
	response.SetHTML()
	response.Body = "Hola Mundo"
}
