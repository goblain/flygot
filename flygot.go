package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	var processConfig bool
	var originalFile string
	parsedFile := "/tmp/fly-template-123123.out"
	args := os.Args[1:]

	// If seting pipeline, process template and replace with template output
	for _, element := range args {
		if element == "set-pipeline" {
			processConfig = true
			fmt.Println("Process config template")
		}
	}

	if processConfig {
		for index, element := range args {
			if element == "-c" {
				originalFile = args[index+1]
				args[index+1] = parsedFile
			}
		}
	}

	tpl, _ := ioutil.ReadFile(originalFile)
	fmt.Printf("%s", tpl)

	t, err := template.New("pipeline").Parse(string(tpl))
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fh, _ := os.Create(parsedFile)

	vars := map[string]string{"a": "b"}
	err = t.Execute(fh, vars)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Printf("%s", args)

	// Execute original fly command with potentially replaced file
	cmd := exec.Cmd{
		Path: "/usr/bin/fly",
		Args: append([]string{"fly"}, args...),
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
}
