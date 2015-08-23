package main

import (
	"fmt"

	"github.com/bndr/gopencils"
)

func main() {
	api := gopencils.Api("http://localhost:5080/api")
	resp := new(interface{})

	// Perform a GET request
	// URL Requested: /api/users
	//Authorization: Bearer W8Ijq0B78JMzYIF26XdakL7od9a4OR
	res := api.Res("todos", resp)
	res.SetHeader("Authorization", "Bearer O6fm69RSKPICKQ2xcNXM7wTp5pMsHl")
	res, _ = res.Get()
	fmt.Println(*resp)
}
