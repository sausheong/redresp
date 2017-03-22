package main

import (
	r "github.com/sausheong/red/responder/core"
)

func main() {
	r.Run(func(req r.RequestInfo, res *r.ResponseInfo) {
		res.SetHTML()
		res.Body = "Hola Mundo"
	})
}
