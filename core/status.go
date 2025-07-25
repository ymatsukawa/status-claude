package core

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

type IAnthropicStatus interface {
	IsElevatedErrors() bool
	GetErrorMessage() string
}

type AnthropicStatus struct {
	Items      []*gofeed.Item
	ErrorTitle *string
	ErrorURL   *string
}

func NewAnthropicStatus(items []*gofeed.Item) IAnthropicStatus {
	return &AnthropicStatus{
		Items:      items,
		ErrorTitle: nil,
		ErrorURL:   nil,
	}
}

func (a *AnthropicStatus) IsElevatedErrors() bool {
	target := a.Items[0] // top item

	if !strings.Contains(target.Title, "Elevated errors") {
		return false
	}

	if strings.Contains(target.Description, "&gt;Resolved&lt") {
		return false
	}

	a.ErrorTitle = &target.Title
	a.ErrorURL = &target.Link

	return true
}

func (a *AnthropicStatus) GetErrorMessage() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(*a.ErrorTitle)
	sb.WriteString(": ")
	sb.WriteString(*a.ErrorURL)
	sb.WriteString("\n")

	return sb.String()
}
