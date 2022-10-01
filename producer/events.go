package main

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/samber/mo"
	uuid "github.com/satori/go.uuid"
)

type EventType string

const (
	EventTypeTrack    EventType = "track"
	EventTypeIdentify EventType = "identify"
	EventTypeGroup    EventType = "group"
	EventTypeScreen   EventType = "screen"
	EventTypePage     EventType = "page"
)

type EventBody struct {
	MessageID   string                 `json:"message_id"`
	TenantID    string                 `json:"tenant_id"`
	UserID      string                 `json:"user_id"`
	GroupID     mo.Option[string]      `json:"group_id"`
	TriggeredAt time.Time              `json:"triggered_at"`
	Type        EventType              `json:"event_type"`
	EventName   mo.Option[string]      `json:"event_name"`
	Properties  map[string]interface{} `json:"properties"`
}

/**
 * Build events
 */
func newRequestCommon(eventType EventType) EventBody {
	return EventBody{
		MessageID:   "m-" + uuid.NewV4().String(),
		TenantID:    RandTenant(),
		UserID:      RandIdentity(),
		TriggeredAt: time.Now().UTC().Add(-1 * time.Duration(rand.Intn(200)) * time.Millisecond),
		Type:        eventType,
	}
}

func NewRequestIdentify() EventBody {
	body := newRequestCommon(EventTypeIdentify)
	body.Properties = RandProperties()
	return body
}

func NewRequestGroup() EventBody {
	body := newRequestCommon(EventTypeGroup)
	body.GroupID = mo.Some("g-" + uuid.NewV4().String())
	body.Properties = RandProperties()
	return body
}

func NewRequestTrack() EventBody {
	body := newRequestCommon(EventTypeTrack)
	body.EventName = mo.Some(RandEventName())
	body.Properties = RandProperties()
	return body
}

func NewRequestScreen() EventBody {
	body := newRequestCommon(EventTypeScreen)
	body.EventName = mo.Some(RandScreenName())
	body.Properties = RandProperties()
	return body
}

func NewRequestPage() EventBody {
	page := RandPage()

	body := newRequestCommon(EventTypePage)
	body.Properties = map[string]interface{}{
		"path":     page.C,         // eg: '/toto/tata/titi'
		"referrer": gofakeit.URL(), // full url
		"search":   page.D,         // eg: '?from=2021-10-05&to=2021-10-05'
		"title":    page.A,
		"url":      page.B,
		"keywords": page.E,
		// "keywords": []string{
		// 	gofakeit.Animal(),
		// 	gofakeit.Fruit(),
		// 	gofakeit.FarmAnimal(),
		// },
	}

	return body
}
