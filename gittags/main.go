package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// Command to get all tags for the current commit
	cmd := exec.Command("git", "tag", "--points-at", "HEAD")

	// Run the command and capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running git command:", err)
		return
	}

	// Convert output to a string and trim any whitespace
	tags := strings.TrimSpace(out.String())

	// Check if there are tags
	if tags == "" {
		fmt.Println("No tags found for the current commit.")
		return
	}

	// Regular expression to match tags in the format "$package/vX.X.X"
	re := regexp.MustCompile(`^(.+?)/v(\d+\.\d+\.\d+)$`)

	// Check if we're running in a GitHub Actions environment
	githubEnv := os.Getenv("GITHUB_ENV")
	if githubEnv == "" {
		fmt.Println("GITHUB_ENV is not set. Are you running this in a GitHub Actions environment?")
		return
	}

	// Open the GITHUB_ENV file and prepare to append variables
	file, err := os.OpenFile(githubEnv, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error opening GITHUB_ENV file:", err)
		return
	}
	defer file.Close()

	// Split tags into lines and process each one
	for _, tag := range strings.Split(tags, "\n") {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		// Find matches
		matches := re.FindStringSubmatch(tag)
		if len(matches) == 3 {
			packageName := matches[1]
			version := matches[2]

			// Create the environment variable name and value
			envVarName := fmt.Sprintf("GIT_TAG_%s_LATEST", strings.ToUpper(packageName))
			envVarValue := fmt.Sprintf("v%s", version)

			// Write to the GITHUB_ENV file
			_, err = file.WriteString(fmt.Sprintf("%s=%s\n", envVarName, envVarValue))
			if err != nil {
				fmt.Println("Error writing to GITHUB_ENV file:", err)
				return
			}

			fmt.Printf("%s environment variable set successfully.\n", envVarName)
		}
	}
}
