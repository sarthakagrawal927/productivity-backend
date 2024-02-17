package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	validators "todo/pkg/utils"

	db "todo/pkg/database"
	googlecalendar "todo/pkg/googlecalendar"
	"todo/pkg/models"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"google.golang.org/api/calendar/v3"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddMaxPriorityTaskToCalendar(c echo.Context) error {

	var tasks []models.Task
	status := c.Get("status").(uint)
	var queryResult *gorm.DB
	if status == 0 {
		queryResult = db.DB_CONNECTION.GetDB().Find(&tasks)
	} else {
		queryResult = db.DB_CONNECTION.GetDB().Order("Priority desc").First(&tasks)
	}

	// ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// Get a Google Calendar client
	client := googlecalendar.GetClient(config)

	// Convert Event to Google Calendar event format
	googleEvent := &calendar.Event{
		Summary:  "testing",
		Location: "testing",
		Start: &calendar.EventDateTime{
			DateTime: time.RFC3339,
			TimeZone: "America/Los_Angeles", // Or your preferred timezone
		},
		End: &calendar.EventDateTime{
			DateTime: time.RFC3339,
			TimeZone: "America/Los_Angeles",
		},
	}

	//Insert the event to Google Calendar
	// Insert event into calendar
	service, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	event, err := service.Events.Insert("primary", googleEvent).Do()
	if err != nil {
		return fmt.Errorf("failed to insert event into calendar: %v %v", err, event)
	}

	// Optionally, store the event in your database after successful Google Calendar insertion
	//err = db.DB.Create(&event).Error
	return validators.HandleQueryResult(queryResult, c, validators.RequestResponse{Message: "Success", Data: tasks}, true)
}

func RetrieveEventFromCalendar() {
	ctx := context.Background()
	b, err := os.ReadFile("../../credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := googlecalendar.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}
