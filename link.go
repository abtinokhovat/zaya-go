package zaya

import (
	"encoding/json"
	"fmt"
	"time"
)

// Link represents a shortened link
type Link struct {
	ID              int       `json:"id"`
	URL             string    `json:"url"`
	Alias           string    `json:"alias"`
	Password        string    `json:"password,omitempty"`
	Space           int       `json:"space,omitempty"`
	Domain          int       `json:"domain,omitempty"`
	Disabled        bool      `json:"disabled"`
	Public          bool      `json:"public"`
	Description     string    `json:"description,omitempty"`
	ExpirationURL   string    `json:"expiration_url,omitempty"`
	ExpirationDate  time.Time `json:"expiration_date,omitempty"`
	ExpirationTime  string    `json:"expiration_time,omitempty"`
	ExpirationClicks int      `json:"expiration_clicks,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// LinkListParams represents parameters for listing links
type LinkListParams struct {
	Search     string `json:"search,omitempty"`
	By         string `json:"by,omitempty"` // title, alias, url
	IDs        string `json:"ids,omitempty"`
	Status     int    `json:"status,omitempty"`
	Space      int    `json:"space,omitempty"`
	Domain     int    `json:"domain,omitempty"`
	Favorites  bool   `json:"favorites,omitempty"`
	Sort       string `json:"sort,omitempty"` // desc, asc, max, min
}

// CreateLinkParams represents parameters for creating a link
type CreateLinkParams struct {
	Alias            string    `json:"alias,omitempty"`
	Password         string    `json:"password,omitempty"`
	Space            int       `json:"space,omitempty"`
	Domain           int       `json:"domain,omitempty"`
	Disabled         bool      `json:"disabled,omitempty"`
	Public           bool      `json:"public,omitempty"`
	Description      string    `json:"description,omitempty"`
	ExpirationURL    string    `json:"expiration_url,omitempty"`
	ExpirationDate   time.Time `json:"expiration_date,omitempty"`
	ExpirationTime   string    `json:"expiration_time,omitempty"`
	ExpirationClicks int       `json:"expiration_clicks,omitempty"`
}

// ListLinks retrieves a list of links based on the provided parameters
func (c *Client) ListLinks(params LinkListParams) ([]Link, error) {
	var links []Link
	err := c.request("GET", "/links", params, &links)
	return links, err
}

// CreateLink creates a new shortened link
func (c *Client) CreateLink(url string, params CreateLinkParams) (*Link, error) {
	paramsMap := map[string]interface{}{
		"url": url,
	}
	
	// Convert params to map
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	
	var paramsMap2 map[string]interface{}
	if err := json.Unmarshal(paramsJSON, &paramsMap2); err != nil {
		return nil, fmt.Errorf("failed to unmarshal params: %w", err)
	}
	
	// Merge maps
	for k, v := range paramsMap2 {
		if v != nil && v != "" && v != 0 && v != false {
			paramsMap[k] = v
		}
	}

	var link Link
	err = c.request("POST", "/links", paramsMap, &link)
	return &link, err
}

// GetLink retrieves details for a specific link
func (c *Client) GetLink(id int) (*Link, error) {
	var link Link
	err := c.request("GET", fmt.Sprintf("/links/%d", id), nil, &link)
	return &link, err
}

// UpdateLink updates an existing link
func (c *Client) UpdateLink(id int, params CreateLinkParams) (*Link, error) {
	var link Link
	err := c.request("PUT", fmt.Sprintf("/links/%d", id), params, &link)
	return &link, err
}

// DeleteLink deletes a link
func (c *Client) DeleteLink(id int) error {
	return c.request("DELETE", fmt.Sprintf("/links/%d", id), nil, nil)
} 