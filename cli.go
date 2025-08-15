package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
		os.Exit(1)
	}
	if count > 1 {
		fmt.Println("Please Provide only one of --value or --json (not both)")
		flag.Usage()
		os.Exit(1)
	}

	switch {
	case *jsonFlag != "":
		var data any
		if err := json.Unmarshal([]byte(*jsonFlag), &data); err != nil {
			fmt.Printf("❌ Invalid JSON: %s\n", err)
			os.Exit(1)
		}
		return CLIArgs{Input: data}
	case *valFlag != "":
		return CLIArgs{Input: parseDynamicValue(*valFlag)}
	}
	return CLIArgs{}

}
