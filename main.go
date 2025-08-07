package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// Define Configuration structure
type Config struct {
	BaseURL  string `json:"base_url"`
	Email    string `json:"email"`
	APIToken string `json:"api_token"`
}

// Define Jira issue structures
type JiraIssue struct {
	ID     string          `json:"id,omitempty"`
	Key    string          `json:"key,omitempty"`
	Self   string          `json:"self,omitempty"`
	Fields JiraIssueFields `json:"fields"`
}

type JiraIssueFields struct {
	Summary     string        `json:"summary"`
	Description string        `json:"desciption"`
	IssueType   JiraIssueType `json:"issuetype"`
	Project     JiraProject   `json:"project"`
	Priority    *JiraPriority `json:"priority,omitempty`
	Assignee    *JiraUser     `json:"assignee,omitempty`
	Status      *JiraStatus   `json:"status,omitempty"`
	Created     string        `json:"created,omitempty"`
	Updated     string        `json:"updated,omitempty"`
}

type JiraIssueType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type JiraProject struct {
	Key  string `json:"key"`
	Name string `json:"name,omitempty"`
}

type JiraPriority struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type JiraUser struct {
	AccountID    string `json:"accountId,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

type JiraStatus struct {
	Name string `json:"name"`
}

type JiraCreateIssueRequest struct {
	Fields JiraIssueFields `json:"fields"`
}

type JiraUpdateIssueRequest struct {
	Fields map[string]interface{} `json:"fields"`
}

// Define global variables
var (
	config     Config
	configDir  = os.Getenv("HOME") + "/goji"
	configFile = configDir + "/config.json"
)

// Define color definitions
var (
	successColor = color.New(color.FgGreen, color.Bold)
	errorColor   = color.New(color.FgRed, color.Bold)
	warningColor = color.New(color.FgYellow, color.Bold)
	infoColor    = color.New(color.FgCyan)
	headerColor  = color.New(color.FgMagenta, color.Bold)
)

// Define HTTP client
var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

// Define Root command
var rootCmd = &cobra.Command{
	Use:   "goji",
	Short: "An opinionated CLI tool for interacting with the Jira API",
	Long:  "A comprehensive, opinionated command line interface for creating and managing Jira issues",
}

// Define Configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure Jira connection settings",
	Run:   runConfigure,
}

// Define Create issue command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Jira issue",
	Run:   runCreate,
}

// Define Main function
func main() {
	// Load configuration
	loadConfig()

	// Add subcommands
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(createCmd)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		errorColor.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Define Configuration Load function
func loadConfig() {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		warningColor.Printf("Warning: Could not read configuration file: %v\n", err)
		return
	}

	if err := json.Unmarshal(data, &config); err != nil {
		warningColor.Printf("Warning: Could not parse configuration file: %v\n", err)
	}
}

// Define Configuration Save function
func saveConfig() error {
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0600)
}

// Define Configure function
func runConfigure(cmd *cobra.Command, args []string) {
	headerColor.Println("=== Goji Configuration ===")

	reader := bufio.NewReader(os.Stdin)

	// Get Jira base URL
	fmt.Print("Jira Base URL (e.g., https://company.atlassian.net): ")
	baseUrl, _ := reader.ReadString('\n')
	config.BaseURL = strings.TrimSpace(baseUrl)

	// Get email address
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	config.Email = strings.TrimSpace(email)

	// Get API token
	fmt.Print("API Token: ")
	byteToken, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		errorColor.Printf("Error reading API token: %v\n", err)
		return
	}
	fmt.Println()
	config.APIToken = string(byteToken)

	// Save configuration
	if err := saveConfig(); err != nil {
		errorColor.Printf("Error saving configuration: %v\n", err)
		return
	}

	successColor.Println("✓ Configuration saved successfully!")
}

// Define API helper function
func makeJiraRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	if config.BaseURL == "" || config.Email == "" || config.APIToken == "" {
		return nil, fmt.Errorf("Goji configuration has not been set. Run 'goji configure' first")
	}

	url := config.BaseURL + "/rest/api/3" + endpoint
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.Email, config.APIToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return httpClient.Do(req)
}

// Define Create issue function
func runCreate(cmd *cobra.Command, args []string) {
	headerColor.Println("=== Create New Jira Issue ===")

	reader := bufio.NewReader(os.Stdin)

	// Get project key
	fmt.Print("Project Key: ")
	projectKey, _ := reader.ReadString('\n')
	projectKey = strings.TrimSpace(projectKey)

	// Get issue type
	fmt.Print("Issue Type (Bug/Task/Story): ")
	issueType, _ := reader.ReadString('\n')
	issueType = strings.TrimSpace(issueType)

	// Get summary
	fmt.Print("Summary: ")
	summary, _ := reader.ReadString('\n')
	summary = strings.TrimSpace(summary)

	// Get description
	fmt.Print("Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	// Create issue request
	issueRequest := JiraCreateIssueRequest{
		Fields: JiraIssueFields{
			Project: JiraProject{
				Key: projectKey,
			},
			IssueType: JiraIssueType{
				Name: issueType,
			},
			Summary: summary,
			Description: description,
		},
	}

	// Send request
	jsonData, _ := json.Marshal(issueRequest)
	resp, err := makeJiraRequest("POST", "/issue", bytes.NewBuffer(jsonData))
	if err != nil {
		errorColor.Printf("Error creating issue: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		errorColor.Printf("Failed to create issue. Status %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response: %s\n", string(body))
		return
	}

	var createdIssue JiraIssue
	if err := json.NewDecoder(resp.Body).Decode(&createdIssue); err != nil {
		errorColor.Printf("Error parsing response: %v\n", err)
		return
	}

	successColor.Printf("✓ Issue created successfully: %s\n", createdIssue.Key)
	infoColor.Printf("URL: %s/browse/%s\n", config.BaseURL, createdIssue.Key)
}
