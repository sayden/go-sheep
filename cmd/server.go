package main

import "net/rpc"

func launchServer() {
	if *server {
		server := rpc.NewServer()
		var maths Maths
		err := server.RegisterName("multiply", &maths)
		if err != nil {
			panic(err)
		}

		server.HandleHTTP("tcp", ":8080")
		return
	}

	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	args := &Args{
		A: 5,
		B: 8,
	}

	var res int
	client.Call("multiply", args, &res)
}
