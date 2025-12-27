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
	fmt.Print(`Welcome to Runner!

`)
	for {
		appTypeInput := getUserInput(`1. Apps
2. Tests
Choose a number: `, 1, 2)

		if appTypeInput == 1 {
			runExecutableApp()
		} else {
			runTests()
		}

		runOrNot := getUserInput(`1. Yes
2. No
Do you want to continue? `, 1, 2)

		if runOrNot == 2 {
			fmt.Println("\nBye!")
			break
		}

		fmt.Println(strings.Repeat("-", 50))
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println(strings.Repeat("-", 50))
	}
}

func getUserInput(prompt string, min, max int) int {
	var input int
	for {
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Reader error: %v\n", err)
			continue
		}
		inputStr = strings.TrimSpace(inputStr)
		input, err = strconv.Atoi(inputStr)
		if err != nil {
			log.Printf("Input error: expected 'a number (0-9)', got '%v'\n\n", inputStr)
			continue
		}
		if input < min || input > max {
			log.Printf("Validation error: Insert a number between %d and %d\n\n", min, max)
			continue
		}
		break
	}
	return input
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
		fmt.Print("No 'main.go' files found in 'examples' directory\n\n")
		return
	}

	var builder strings.Builder
	builder.WriteString("\n")
	for i, v := range runnableFiles {
		fmt.Fprintf(&builder, "%d: '%s'\n", i+1, v)
	}
	builder.WriteString("Choose a number: ")
	appNumberInput := getUserInput(builder.String(), 1, appSize)

	chosenFile := runnableFiles[appNumberInput-1]

	fmt.Println("\n"+strings.Repeat("=", 18), "Go RUN", strings.Repeat("=", 18))
	cmd := exec.Command("go", "run", chosenFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("Program execution failed: %v\n", err)
	}
	fmt.Println(strings.Repeat("=", 44))
}

func runTests() {
	var testPackages []string
	packagesMap := make(map[string]bool)

	if err := filepath.WalkDir("examples", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), "_test.go") {
			dir := filepath.Dir(path)
			pkgPath := filepath.ToSlash(dir)
			if !strings.HasPrefix(pkgPath, "./") {
				pkgPath = "./" + pkgPath
			}

			if !packagesMap[pkgPath] {
				packagesMap[pkgPath] = true
				testPackages = append(testPackages, pkgPath)
			}
		}
		return nil
	}); err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}

	testsSize := len(testPackages)
	if testsSize == 0 {
		fmt.Print("No '*_test.go' files found in 'examples' directory\n\n")
		return
	}

	var builder strings.Builder
	builder.WriteString("\n")
	for i, pkg := range testPackages {
		fmt.Fprintf(&builder, "%d: '%s'\n", i+1, pkg)
	}
	builder.WriteString("Choose a number: ")
	pkgNumber := getUserInput(builder.String(), 1, testsSize)

	chosenPackage := testPackages[pkgNumber-1]

	fmt.Println("\n"+strings.Repeat("=", 18), "Go TEST", strings.Repeat("=", 18))
	cmd := exec.Command("go", "test", "-v", chosenPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("Test execution failed: %v\n", err)
	}
	fmt.Println(strings.Repeat("=", 45))
}
