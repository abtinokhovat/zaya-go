package zaya

import (
	"fmt"
	"time"
)

// Domain represents a domain in Zaya
type Domain struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	IndexPage      string    `json:"index_page,omitempty"`
	NotFoundPage   string    `json:"not_found_page,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// DomainListParams represents parameters for listing domains
type DomainListParams struct {
	Search string `json:"search,omitempty"`
	Sort   string `json:"sort,omitempty"` // desc, asc
}

// CreateDomainParams represents parameters for creating a domain
type CreateDomainParams struct {
	Name         string `json:"name"`
	IndexPage    string `json:"index_page,omitempty"`
	NotFoundPage string `json:"not_found_page,omitempty"`
}

// ListDomains retrieves a list of domains based on the provided parameters
func (c *Client) ListDomains(params DomainListParams) ([]Domain, error) {
	var domains []Domain
	err := c.request("GET", "/domains", params, &domains)
	return domains, err
}

// CreateDomain creates a new domain
func (c *Client) CreateDomain(params CreateDomainParams) (*Domain, error) {
	var domain Domain
	err := c.request("POST", "/domains", params, &domain)
	return &domain, err
}

// GetDomain retrieves details for a specific domain
func (c *Client) GetDomain(id int) (*Domain, error) {
	var domain Domain
	err := c.request("GET", fmt.Sprintf("/domains/%d", id), nil, &domain)
	return &domain, err
}

// UpdateDomain updates an existing domain
func (c *Client) UpdateDomain(id int, params CreateDomainParams) (*Domain, error) {
	var domain Domain
	err := c.request("PUT", fmt.Sprintf("/domains/%d", id), params, &domain)
	return &domain, err
}

// DeleteDomain deletes a domain
func (c *Client) DeleteDomain(id int) error {
	return c.request("DELETE", fmt.Sprintf("/domains/%d", id), nil, nil)
} 