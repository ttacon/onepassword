package onepassword

import "time"

// NOTE(ttacon): I'm not sure how I feel about this cursor structure.
// On one hand, from a usage perspective it's "clean enough". On the other hand,
// it's two params for passing state in...

type EventsService interface {
	Introspect() (*IntrospectionResponse, error)
	GetItemUsages(resetCursor *ResetCursor, currCursor string) (*ItemUsageResponse, error)
	GetSignInAttempts(resetCursor *ResetCursor, currCursor string) (*SignInAttemptsResponse, error)
}

type IntrospectionResponse struct {
	UUID     string    `json:"UUID"`
	IssuedAt time.Time `json:"IssuedAt"`
	Features []string  `json:"features"`
}

type ResetCursor struct {
	Limit     int        `json:"limit"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time,omitempty"`
}

type ItemUsageResponse struct {
	Items   []ItemUsage `json:"items"`
	Cursor  string      `json:"cursor"`
	HasMore bool        `json:"has_more"`
}

type ItemUsage struct {
	UUID        string     `json:"uuid"`
	Timestamp   time.Time  `json:"timestamp"`
	UsedVersion int        `json:"used_version"`
	VaultUUID   string     `json:"vault_uuid"`
	ItemUUID    string     `json:"item_uuid"`
	User        User       `json:"user"`
	ClientInfo  ClientInfo `json:"client"`
}

type User struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ClientInfo struct {
	AppName         string `json:"app_name"`
	AppVersion      string `json:"app_version"`
	PlatformName    string `json:"platform_name"`
	PlatformVersion string `json:"platform_version"`
	OSName          string `json:"os_name"`
	OSVersion       string `json:"os_version"`
	IPAddress       string `json:"ip_address"`
}

type SignInAttemptsResponse struct {
	Items   []SignInAttempt `json:"items"`
	Cursor  string          `json:"cursor"`
	HasMore bool            `json:"has_more"`
}

type SignInAttempt struct {
	UUID        string            `json:"uuid"`
	SessionUUID string            `json:"session_uuid"`
	Timestamp   time.Time         `json:"timestamp"`
	Category    string            `json:"category"`
	Type        string            `json:"type"`
	Country     string            `json:"country"`
	TargetUser  User              `json:"target_user"`
	ClientInfo  ClientInfo        `json:"client"`
	Details     map[string]string `json:"details"`
}
