package main

import "fmt"
import "strings"
import "./classpath/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.
		fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
