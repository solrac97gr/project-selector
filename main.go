package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/solrac97gr/project-selector/config"
)

func main() {
	config := config.NewConfig()
	if err := config.LoadConfigFromFile(); err != nil {
		config.SetDefaultConfig()
	}

	// Get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	// Initialize the projectDirs slice
	var projectDirs []string

	// Iterate over the config.ProjectDirs array and load the project directories
	for _, dir := range config.ProjectDirs {
		rootDir := filepath.Join(usr.HomeDir, dir)
		err = findProjectDirs(rootDir, &projectDirs, config.DirsToIgnore)
		if err != nil {
			fmt.Printf("Error walking the path %v: %v\n", rootDir, err)
			return
		}
	}

	// If no projects found
	if len(projectDirs) == 0 {
		fmt.Println("No projects found.")
		return
	}

	// Select a project
	selectedProjectPath, err := selectProject(projectDirs, config)
	if err != nil {
		if strings.Contains(err.Error(), "^C") {
			fmt.Println("Selection canceled.")
			return
		}
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Execute the zed command with the selected project path
	if selectedProjectPath != "" {
		err := executeZedCommand(selectedProjectPath, config.CMD)
		if err != nil {
			fmt.Printf("Error executing zed: %v\n", err)
		}
	}
}

func findProjectDirs(rootDir string, projectDirs *[]string, dirsToIgnore []string) error {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore .git and hidden directories
		if strings.Contains(path, ".git") || strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// Check if the directory is in the dirsToIgnore list
		for _, dir := range dirsToIgnore {
			if strings.Contains(path, dir) {
				return nil
			}
		}

		// Only consider directories that are exactly one level deep inside rootDir
		relPath, _ := filepath.Rel(rootDir, path)
		if info.IsDir() && relPath != "." && !strings.Contains(relPath, string(os.PathSeparator)) {
			// Extract the project name (last part of the path)
			*projectDirs = append(*projectDirs, path)
		}

		return nil
	})

	return err
}

func selectProject(projectDirs []string, config *config.Config) (string, error) {
	// Prepare the project names for the selection list
	var projectNames []string
	for _, dir := range projectDirs {
		_, name := splitProjectPath(dir)
		projectNames = append(projectNames, name)
	}

	// Setup the interactive prompt
	prompt := promptui.Select{
		Label: fmt.Sprintf("Select a Project %s (ctrl+c to cancel)", config.Style.Title.Icon),
		Items: projectNames,
		Size:  config.NumberOfProjects,
		Searcher: func(input string, index int) bool {
			return strings.Contains(projectNames[index], input)
		},
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Label:    config.Style.Title.Template,
			Active:   fmt.Sprintf("%s %s %s", config.Style.Active.Icon, config.Style.Inactive.Icon, config.Style.Active.Template),
			Inactive: fmt.Sprintf(" %s %s", config.Style.Inactive.Icon, config.Style.Inactive.Template),
			Selected: "✅ {{ . | blue }}",
		},
	}

	// Run the prompt
	_, selectedProject, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// Get the full path of the selected project
	selectedProjectPath := getProjectPath(selectedProject, projectDirs)
	return selectedProjectPath, nil
}

func splitProjectPath(path string) (string, string) {
	// Split the path and return the last segment as the project name
	parts := strings.Split(path, "/")
	return parts[len(parts)-2], parts[len(parts)-1]
}

func getProjectPath(selectedName string, projectDirs []string) string {
	// Find the full path of the selected project
	for _, dir := range projectDirs {
		_, name := splitProjectPath(dir)
		if name == selectedName {
			return dir
		}
	}
	return ""
}

func executeZedCommand(projectPath string, openWith string) error {
	// Prepare the zed command with the project path
	cmd := exec.Command(openWith, projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute zed command: %w", err)
	}
	return nil
}
