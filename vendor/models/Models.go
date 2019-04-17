package models

import (
	"github.com/jinzhu/gorm"
)

// EnvConstants represents constants provided by the runtime environment
type EnvConstants struct {
	Token      string
	BotToken   string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
}

// URLVerificationModel represents post request body which slack requests during URL verification
type URLVerificationModel struct {
	EventModel
	Challenge string `json:"challenge"`
}

// EventModel represents the basic request body which all events look like
type EventModel struct {
	Token string `json:"token"`
	Event struct {
		Type string `json:"type"`
	} `json:"event"`
	Type string `json:"type"`
}

// UserChangeRequest represents the request when a user changes their information
type UserChangeRequest struct {
	Event struct {
		User `json:"user"`
	} `json:"event"`
}

// UserListResponse represents the response of slack.com/api/users.list
type UserListResponse struct {
	Ok               bool   `json:"ok"`
	Offset           string `json:"offset"`
	Members          []User `json:"members"`
	CacheTs          int    `json:"cache_ts"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

// User represents a user in the installed Slack workspace
type User struct {
	gorm.Model
	ProfileID         uint         `json:"-" gorm:"ForeignKey:id"`
	Profile           SlackProfile `json:"profile"`
	SlackID           string       `json:"id" gorm:"index"`
	TeamID            string       `json:"team_id"`
	Name              string       `json:"name"`
	Deleted           bool         `json:"deleted"`
	Color             string       `json:"color"`
	RealName          string       `json:"real_name"`
	Tz                string       `json:"tz"`
	TzLabel           string       `json:"tz_label"`
	TzOffset          int          `json:"tz_offset"`
	IsAdmin           bool         `json:"is_admin"`
	IsOwner           bool         `json:"is_owner"`
	IsPrimaryOwner    bool         `json:"is_primary_owner"`
	IsRestricted      bool         `json:"is_restricted"`
	IsUltraRestricted bool         `json:"is_ultra_restricted"`
	IsBot             bool         `json:"is_bot"`
	IsAppUser         bool         `json:"is_app_user"`
	Updated           int          `json:"updated"`
	Locale            string       `json:"locale"`
}

// SlackProfile represents the profile of a user
type SlackProfile struct {
	gorm.Model
	UserID                uint   `json:"-"`
	Title                 string `json:"title"`
	Phone                 string `json:"phone"`
	Skype                 string `json:"skype"`
	RealName              string `json:"real_name"`
	RealNameNormalized    string `json:"real_name_normalized"`
	DisplayName           string `json:"display_name"`
	DisplayNameNormalized string `json:"display_name_normalized"`
	StatusText            string `json:"status_text"`
	StatusEmoji           string `json:"status_emoji"`
	StatusExpiration      int    `json:"status_expiration"`
	AvatarHash            string `json:"avatar_hash"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Image24               string `json:"image_24"`
	Image32               string `json:"image_32"`
	Image48               string `json:"image_48"`
	Image72               string `json:"image_72"`
	Image192              string `json:"image_192"`
	Image512              string `json:"image_512"`
	StatusTextCanonical   string `json:"status_text_canonical"`
	Team                  string `json:"team"`
}
