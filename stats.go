package zaya

import (
	"fmt"
	"time"
)

// StatsParams represents parameters for retrieving statistics
type StatsParams struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

// TotalStats represents total statistics for a link
type TotalStats struct {
	Total int `json:"total"`
}

// ClickStats represents click statistics for a link
type ClickStats struct {
	Date  string `json:"date"`
	Clicks int   `json:"clicks"`
}

// ReferrerStats represents referrer statistics for a link
type ReferrerStats struct {
	Referrer string `json:"referrer"`
	Clicks   int    `json:"clicks"`
}

// CountryStats represents country statistics for a link
type CountryStats struct {
	Country string `json:"country"`
	Clicks  int    `json:"clicks"`
}

// LanguageStats represents language statistics for a link
type LanguageStats struct {
	Language string `json:"language"`
	Clicks   int    `json:"clicks"`
}

// BrowserStats represents browser statistics for a link
type BrowserStats struct {
	Browser string `json:"browser"`
	Clicks  int    `json:"clicks"`
}

// DeviceStats represents device statistics for a link
type DeviceStats struct {
	Device string `json:"device"`
	Clicks int    `json:"clicks"`
}

// OperatingSystemStats represents operating system statistics for a link
type OperatingSystemStats struct {
	OS     string `json:"os"`
	Clicks int    `json:"clicks"`
}

// GetTotalStats retrieves total statistics for a link
func (c *Client) GetTotalStats(linkID int, params StatsParams) (*TotalStats, error) {
	var stats TotalStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/total", linkID), params, &stats)
	return &stats, err
}

// GetClickStats retrieves click statistics for a link
func (c *Client) GetClickStats(linkID int, params StatsParams) ([]ClickStats, error) {
	var stats []ClickStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/clicks", linkID), params, &stats)
	return stats, err
}

// GetReferrerStats retrieves referrer statistics for a link
func (c *Client) GetReferrerStats(linkID int, params StatsParams) ([]ReferrerStats, error) {
	var stats []ReferrerStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/referrers", linkID), params, &stats)
	return stats, err
}

// GetCountryStats retrieves country statistics for a link
func (c *Client) GetCountryStats(linkID int, params StatsParams) ([]CountryStats, error) {
	var stats []CountryStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/countries", linkID), params, &stats)
	return stats, err
}

// GetLanguageStats retrieves language statistics for a link
func (c *Client) GetLanguageStats(linkID int, params StatsParams) ([]LanguageStats, error) {
	var stats []LanguageStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/languages", linkID), params, &stats)
	return stats, err
}

// GetBrowserStats retrieves browser statistics for a link
func (c *Client) GetBrowserStats(linkID int, params StatsParams) ([]BrowserStats, error) {
	var stats []BrowserStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/browsers", linkID), params, &stats)
	return stats, err
}

// GetDeviceStats retrieves device statistics for a link
func (c *Client) GetDeviceStats(linkID int, params StatsParams) ([]DeviceStats, error) {
	var stats []DeviceStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/devices", linkID), params, &stats)
	return stats, err
}

// GetOperatingSystemStats retrieves operating system statistics for a link
func (c *Client) GetOperatingSystemStats(linkID int, params StatsParams) ([]OperatingSystemStats, error) {
	var stats []OperatingSystemStats
	err := c.request("GET", fmt.Sprintf("/links/%d/stats/os", linkID), params, &stats)
	return stats, err
} 