package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

type DockerCompose struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Ports []string `yaml:"ports"`
}

func main() {
	outputFile := "out.txt"
	dockerComposeFile := "docker-compose.yml"
	fmt.Printf("Reading %s\n", dockerComposeFile)
	// Read the YAML file
	file, err := os.ReadFile(dockerComposeFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Unmarshal the YAML into a struct
	var compose DockerCompose
	err = yaml.Unmarshal(file, &compose)
	if err != nil {
		fmt.Printf("Error unmarshalling YAML: %v\n", err)
		return
	}

	cmdEmptyFile := exec.Command("bash", "-c", fmt.Sprintf("echo \"\" > %s", outputFile))
	// Execute the command and capture the output
	_, err = cmdEmptyFile.Output()

	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	fmt.Printf("Extracting services from: %s\n", dockerComposeFile)

	// Iterate over the services and print the host ports
	for serviceName, service := range compose.Services {
		if strings.HasPrefix(serviceName, "go-") {
			for _, port := range service.Ports {
				hostPort := port[:strings.Index(port, ":")]
				fmt.Printf("Benchmarking docker service: %s, on port: %s\n", serviceName, hostPort)
				cmdText := fmt.Sprintf("echo \"Name:  %s\" >> %s && oha -j --no-tui -n 100 -c 100 -p 100 -z 30s http://localhost:%s >> %s", serviceName, outputFile, hostPort, outputFile)
				// Define the command with arguments
				cmd := exec.Command("bash", "-c", cmdText)

				// Execute the command and capture the output
				_, err := cmd.Output()
				if err != nil {
					fmt.Printf("Error executing command: %v\n", err)
					return
				}
			}
		}
	}
}
