package model

import "fmt"

// EventType ...
type EventType uint8

// EventStatus ...
type EventStatus uint8

const (
	// InvalidType ...
	InvalidType EventType = iota

	// Created ...
	Created

	// Updated ...
	Updated

	// Removed ...
	Removed

	// InvalidStatus ...
	InvalidStatus EventStatus = iota

	// Idle ...
	Idle

	// Deferred ...
	Deferred

	// Processed ...
	Processed
)

// Request ...
type Request struct {
	ID      uint64 `json:"id,omitempty" db:"id"`
	Service string `json:"service,omitempty" db:"service"`
	User    string `json:"user,omitempty" db:"user"`
	Text    string `json:"desc,omitempty" db:"text"`
}

// RequestEvent ...
type RequestEvent struct {
	ID     uint64      `json:"id,omitempty"`
	Type   EventType   `json:"type,omitempty"`
	Status EventStatus `json:"status,omitempty"`
	Entity *Request    `json:"entiry,omitempty"`
}

var (
	evTypeStr = map[EventType]string{
		Created: "Created",
		Removed: "Removed",
		Updated: "Updated",
	}
	evStatusStr = map[EventStatus]string{
		Idle:      "Idle",
		Deferred:  "Deferred",
		Processed: "Processed",
	}
)

func (r Request) String() string {
	return fmt.Sprintf("id: %d; service: %s; user: %s; text: %s",
		r.ID, r.Service, r.User, r.Text)
}

func (e RequestEvent) String() string {
	return fmt.Sprintf("RequestEvent { id: %d; type: %s; status: %s }",
		e.ID, evTypeStr[e.Type], evStatusStr[e.Status])
}

// EventTypeStrToVal ...
func EventTypeStrToVal(t string) EventType {
	switch t {
	case "Created":
		return Created
	case "Removed":
		return Removed
	case "Updated":
		return Updated
	default:
		return InvalidType
	}
}
