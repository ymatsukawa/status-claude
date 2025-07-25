package core

import (
	"testing"

	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
)

func TestIsElevatedErrors(t *testing.T) {
	tests := []struct {
		name     string
		items    []*gofeed.Item
		expected bool
	}{
		{
			name: "should return true for elevated errors without resolved status",
			items: []*gofeed.Item{
				{
					Title:       "Elevated errors in Claude API",
					Description: "We are experiencing elevated errors",
					Link:        "https://claude.example.com/incidents/123",
				},
			},
			expected: true,
		},
		{
			name: "should return false for resolved elevated errors",
			items: []*gofeed.Item{
				{
					Title:       "Elevated errors in Claude API",
					Description: "Issue resolved &gt;Resolved&lt",
					Link:        "https://claude.example.com/incidents/123",
				},
			},
			expected: false,
		},
		{
			name: "should return false for non-elevated error titles",
			items: []*gofeed.Item{
				{
					Title:       "Service maintenance",
					Description: "Scheduled maintenance",
					Link:        "https://claude.example.com/incidents/124",
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := NewAnthropicStatus(tt.items)
			assert.Equal(t, tt.expected, status.IsElevatedErrors())
		})
	}
}

func TestGetErrorMessage(t *testing.T) {
	tests := []struct {
		name       string
		errorTitle *string
		errorURL   *string
		expected   string
	}{
		{
			name:       "should return formatted message with title and URL",
			errorTitle: stringPtr("Elevated errors in Claude API"),
			errorURL:   stringPtr("https://claude.example.com/incidents/123"),
			expected:   "\nElevated errors in Claude API: https://claude.example.com/incidents/123\n",
		},
		// just functional test cases
		{
			name:       "should handle different error titles and URLs",
			errorTitle: stringPtr("Service disruption"),
			errorURL:   stringPtr("https://status.anthropic.com/incidents/456"),
			expected:   "\nService disruption: https://status.anthropic.com/incidents/456\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := &AnthropicStatus{
				ErrorTitle: tt.errorTitle,
				ErrorURL:   tt.errorURL,
			}
			result := status.GetErrorMessage()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
