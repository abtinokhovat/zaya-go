package zaya

import (
	"fmt"
	"time"
)

// Space represents a space in Zaya
type Space struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SpaceListParams represents parameters for listing spaces
type SpaceListParams struct {
	Search string `json:"search,omitempty"`
	Sort   string `json:"sort,omitempty"` // desc, asc
}

// CreateSpaceParams represents parameters for creating a space
type CreateSpaceParams struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// ListSpaces retrieves a list of spaces based on the provided parameters
func (c *Client) ListSpaces(params SpaceListParams) ([]Space, error) {
	var spaces []Space
	err := c.request("GET", "/spaces", params, &spaces)
	return spaces, err
}

// CreateSpace creates a new space
func (c *Client) CreateSpace(params CreateSpaceParams) (*Space, error) {
	var space Space
	err := c.request("POST", "/spaces", params, &space)
	return &space, err
}

// GetSpace retrieves details for a specific space
func (c *Client) GetSpace(id int) (*Space, error) {
	var space Space
	err := c.request("GET", fmt.Sprintf("/spaces/%d", id), nil, &space)
	return &space, err
}

// UpdateSpace updates an existing space
func (c *Client) UpdateSpace(id int, params CreateSpaceParams) (*Space, error) {
	var space Space
	err := c.request("PUT", fmt.Sprintf("/spaces/%d", id), params, &space)
	return &space, err
}

// DeleteSpace deletes a space
func (c *Client) DeleteSpace(id int) error {
	return c.request("DELETE", fmt.Sprintf("/spaces/%d", id), nil, nil)
} 