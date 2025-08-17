package main

import (
	"encoding/json"
	"time"
)

// Notification represents a single notification
type Notification struct {
	ID        string    `json:"id"`
	pipeline  string    `json:"pipeline"`
	status    string    `json:"status"`
	timestamp time.Time `json:"timestamp"`
	message   string    `json:"message"`
}

// NotifierConfig represents the configuration for the notifier
type NotifierConfig struct {
	WebhookURL  string `json:"webhook_url"`
	APIKey     string `json:"api_key"`
	APIURL     string `json:"api_url"`
	PipelineID string `json:"pipeline_id"`
}

// DevOpsPipeline represents a DevOps pipeline
type DevOpsPipeline struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Project  string `json:"project"`
	Status   string `json:"status"`
	LastRun  time.Time `json:"last_run"`
	NextRun  time.Time `json:"next_run"`
	Notifier NotifierConfig `json:"notifier"`
}

// DevOpsNotifier represents the DevOps notifier
type DevOpsNotifier struct {
	Pipeline DevOpsPipeline `json:"-"`
	Notifications []Notification `json:"-"`
}

func (dn *DevOpsNotifier) AddNotification(n Notification) {
	dn.Notifications = append(dn.Notifications, n)
}

func (dn *DevOpsNotifier) SendNotifications() error {
	// TO DO: implement notification sending logic using webhook or API
	return nil
}

func main() {
	// TO DO: initialize and use the DevOpsNotifier
}