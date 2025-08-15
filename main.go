package main

func main() {
	args := ParseFlags()
	Inspect(args.Input)
}
