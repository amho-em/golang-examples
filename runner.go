package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var runnableFiles []string

	if err := filepath.WalkDir("examples", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() == "main.go" {
			runnableFiles = append(runnableFiles, path)
		}
		return nil
	}); err != nil {
		log.Fatalf("filepath error: %v\n", err)
	}

	appSize := len(runnableFiles)
	if appSize == 0 {
		fmt.Println("No 'main.go' files found in 'examples' directory")
		os.Exit(0)
	}

	for i, v := range runnableFiles {
		fmt.Printf("%d: '%s'\n", i+1, v)
	}
	fmt.Print("Choose a number: ")
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("bufio error: %v\n", err)
	}
	inputStr = strings.TrimSpace(inputStr)
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		log.Fatalf("Input error: %v\n", err)
	}

	if input < 1 || input > appSize {
		log.Fatalln("Input error: the number is out of the range")
	}
	chosenFile := runnableFiles[input-1]

	fmt.Println("\n"+strings.Repeat("=", 18), "Go", strings.Repeat("=", 18))
	cmd := exec.Command("go", "run", chosenFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("Program execution failed: %v", err)
		os.Exit(1)
	}
	fmt.Println(strings.Repeat("=", 40))
}
