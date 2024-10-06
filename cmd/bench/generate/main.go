package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

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

	cmdEmptyFile := exec.Command("bash", "-c", fmt.Sprintf(": > %s", outputFile))
	// Execute the command and capture the output
	_, err = cmdEmptyFile.Output()

	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	fmt.Printf("Extracting services from: %s\n", dockerComposeFile)
	services := make([]string, 0, len(compose.Services))
	for serviceName := range compose.Services {
		if strings.HasPrefix(serviceName, "go-") {
			services = append(services, serviceName)
		}
	}

	for i, serviceName := range services {
		service := compose.Services[serviceName]
		for _, port := range service.Ports {
			hostPort := port[:strings.Index(port, ":")]
			fmt.Printf("Benchmarking docker service: %s, on port: %s\n", serviceName, hostPort)
			cmdText := fmt.Sprintf("echo \"\nName:  %s\" >> %s && oha -j --no-tui -n 120000 -c 1000 -p 500 -z 1m http://localhost:%s >> %s", serviceName, outputFile, hostPort, outputFile)
			cmd := exec.Command("bash", "-c", cmdText)

			_, err := cmd.Output()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				return
			}

			if i < len(services)-1 {
				coolDown := 30 * time.Second
				fmt.Printf("Waiting (%s) for resources to cooldown\n", coolDown)
				time.Sleep(coolDown)
			}
		}
	}
}
