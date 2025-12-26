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
	fmt.Println(`1. Apps
2. Tests`)
	appTypeInput, err := getUserInput()
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}
	if appTypeInput != 1 && appTypeInput != 2 {
		log.Fatalln("Input error: the number is out of the range")
	}

	if appTypeInput == 1 {
		runExecutableApp()
		return
	}
	// todo
	log.Println("Run Tests here ...")
}

func getUserInput() (int, error) {
	fmt.Print("Choose a number: ")
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	inputStr = strings.TrimSpace(inputStr)
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return 0, fmt.Errorf("expected 'a number (0-9)', got '%v'", inputStr)
	}
	return input, nil
}

func runExecutableApp() {
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
		fmt.Printf("\n%d: '%s'\n", i+1, v)
	}
	appNumberInput, err := getUserInput()
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}
	if appNumberInput < 1 || appNumberInput > appSize {
		log.Fatalln("Input error: the number is out of the range")
	}
	chosenFile := runnableFiles[appNumberInput-1]

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
