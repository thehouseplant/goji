package models

import "time"

// Issue represents a Jira issue
type Issue struct {
	ID     string      `json:"id"`
	Key    string      `json:"key"`
	Self   string      `json:"self"`
	Fields IssueFields `json:"fields"`
}

// IssueFields contains the field data for a Jira issue
type IssueFields struct {
	Summary     string      `json:"summary"`
	Description *string     `json:"desciption,omitempty"`
	IssueType   *IssueType  `json:"issuetype,omitempty"`
	Project     *Project    `json:"project,omitempty"`
	Status      *Status     `json:"status,omitempty"`
	Priority    *Priority   `json:"priority,omitempty"`
	Assignee    *User       `json:"assignee,omitempty"`
	Reporter    *User       `json:"reporter,omitempty`
	Created     *time.Time  `json:"created,omitempty"`
	Updated     *time.Time  `json:"updated,omitempty"`
	Labels      []string    `json:"labels,omitempty"`
	Components  []Component `json:"components,omitempty"`
}

// IssueType represents a Jira issue type
type IssueType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
}

// Preoject represnts a Jira project
type Project struct {
	ID   string `json:",id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

// Status represents a Jira issue status
type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Priority represents a Jira issue priority
type Priority struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User represents a Jira user
type User struct {
	AccountID    string `json:"accountId"`
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

// Component represents a Jira project component
type Component struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateIssueRequest represents the request to create a Jira issue
type CreateIssueRequest struct {
	Fields CreateIssueFields `json:"fields"`
}

// CreateIssueFields contains the fields for creating a Jira issue
type CreateIssueFields struct {
	Summary     string      `json:"summary"`
	Description *string     `json:"description,omitempty"`
	IssueType   IssueType   `json:"issuetype"`
	Project     Project     `json:"project"`
	Priority    *Priority   `json:"priority,omitempty"`
	Assignee    *User       `json:"assignee,omitempty"`
	Labels      []string    `json:"labels,omitempty"`
	Components  []Component `json:"componenets,omitempty"`
}

// UpdateIssueRequest represents the request to update a Jira issue
type UpdateIssueRequest struct {
	Fields map[string]interface{} `json:"fields"`
}

// SearchResponse represents the response from a search request
type SearchResponse struct {
	Issues     []Issue `json:"issues"`
	Total      int     `json:"total"`
	MaxResults int     `json:"maxResults"`
	StartAt    int     `json:"startAt"`
}

// CreateIssueResponse represents the response from creating a Jira issue
type CreateIssueReponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

// ErrorResponse represents an error from Jira
type ErrorResponse struct {
	ErrorMessages []string          `json:"errorMessages,omitempty"`
	Errors        map[string]string `json:"errors,omitempty"`
}
