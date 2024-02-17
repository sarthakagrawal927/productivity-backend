package service

import (
	"context"
	"time"
	db "todo/pkg/database"
	googlecalendar "todo/pkg/googlecalendar"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func AddEventToCalendar(event Event) error {
	// Get a Google Calendar client
	client := googlecalendar.GetClient(config)

	// Convert Event to Google Calendar event format
	googleEvent := &calendar.Event{
		Summary:  event.Summary,
		Location: event.Location,
		Start: &calendar.EventDateTime{
			DateTime: event.Start.Format(time.RFC3339),
			TimeZone: "America/Los_Angeles", // Or your preferred timezone
		},
		End: &calendar.EventDateTime{
			DateTime: event.End.Format(time.RFC3339),
			TimeZone: "America/Los_Angeles",
		},
	}

	// Insert the event to Google Calendar
	service := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	_, err := service.Events.Insert("primary", googleEvent).Do()
	if err != nil {
		return err
	}

	// Optionally, store the event in your database after successful Google Calendar insertion
	err = db.DB.Create(&event).Error
	return err
}
