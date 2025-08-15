package main

import (
	"flag"
	"fmt"
)

type CLIArgs struct {
	Input any
}

func ParseFlags() CLIArgs {
	valFlag := flag.String("value", "", "Inspect a single value (bool,int,float or string)")
	jsonFlag := flag.String("json", "", "Inspect a JSON Object or array")
	flag.Parse()

	count := countSet(*valFlag, *jsonFlag)

	if count == 0 {
		fmt.Println("X Please Provide one of --value or --json")
		flag.Usage()
		panic("No Input Provided")
	}
	if count > 1 {
		fmt.Println("Please Provide only one of --value or --json (not both)")
		flag.Usage()
		panic("Conflicting input flags")
	}
	return CLIArgs{}

}
