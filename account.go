package zaya

import (
	"time"
)

// Account represents a Zaya account
type Account struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAccount retrieves the current account details
func (c *Client) GetAccount() (*Account, error) {
	var account Account
	err := c.request("GET", "/account", nil, &account)
	return &account, err
} 