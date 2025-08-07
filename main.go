package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
}

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
	configDir  = os.Getenv("HOME") + "./goji"
	configFile = configDir + "/config.json"
)
