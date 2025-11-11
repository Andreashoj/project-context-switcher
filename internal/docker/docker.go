package docker

import (
	"fmt"
	"os"
	"path/filepath"
	"project-context-switcher/internal/models"
	"strings"

	"gopkg.in/yaml.v3"
)

func GetFile(path string) ([]byte, error) {
	// Convert path to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed converting path to absolute path: %w", err)
	}

	// In case of user having inputted a path ending with a file
	info, err := os.Stat(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving path when checking directory: %w", err)
	}

	// Ensure path given is a directory
	dir := absPath
	if !info.IsDir() {
		dir = filepath.Dir(absPath)
	}

	// Ensure directory exists
	_, err = os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("failed accesing file path directory: %w", err)
	}

	// Ensure docker compose file is present
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed trying to read directory files: %w", err)
	}

	dockerComposeExists := false
	var fileName string
	for i := 0; i < len(files); i++ {
		if strings.Contains(files[i].Name(), "docker-compose") &&
			(strings.HasSuffix(files[i].Name(), ".yml") || strings.HasSuffix(files[i].Name(), ".yaml")) {
			dockerComposeExists = true
			fileName = files[i].Name()
			break
		}
	}

	if !dockerComposeExists {
		return nil, fmt.Errorf("directory must hold docker-compose file")
	}

	file, err := os.ReadFile(absPath + "/" + fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading docker-compose file")
	}

	return file, nil
}

func GetContainers(file []byte) (*models.DockerCompose, error) {
	var compose models.DockerCompose
	if err := yaml.Unmarshal(file, &compose); err != nil {
		return nil, fmt.Errorf("failed decoding the document: %w", err)
	}

	// Begin seeing if I can update the environment variables!
	// Post this to the frontend!

	return &compose, nil
}
