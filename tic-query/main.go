package main

import "tic-chain/tic-query/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
