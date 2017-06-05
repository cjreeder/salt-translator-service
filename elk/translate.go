package elk

import (
	"errors"
	"log"
	"regexp"
	"strings"

	ei "github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"github.com/byuoitav/salt-translator-service/salt"
)

//converts salt events to API events
func translate(event salt.Event) (Event, error) {
	beacon := regexp.MustCompile(`\/beacon`)
	present := regexp.MustCompile(`(\/presence)(\/present)`)
	change := regexp.MustCompile(`(\/presence)(\/change)`)

	if present.MatchString(event.Tag) {
		return translatePresent(event)
	} else if change.MatchString(event.Tag) {
		return translateChange(event)
	} else if beacon.MatchString(event.Tag) {
		return translateBeacon(event)
	} else {
		return Event{}, errors.New("Could not translate event! Unrecognized tag.")
	}

}

func translatePresent(event salt.Event) (Event, error) {

	log.Printf("Identified salt presence/present event, tag: %s", event.Tag)

	data := event.Data["_stamp"]
	timestamp := data.(string)

	return Event{
		Building:  "Multiple",
		Room:      "Multiple",
		Cause:     ei.AUTOGENERATED.String(),
		Category:  "Heartbeat",
		Hostname:  "Multiple",
		Timestamp: timestamp,
		Data:      event.Data,
	}, nil
}

func translateChange(event salt.Event) (Event, error) {

	log.Printf("Identified salt presence/change event, tag: %s", event.Tag)

	data := event.Data["_stamp"]
	timestamp := data.(string)

	return Event{
		Building:  "Multiple",
		Room:      "Multiple",
		Cause:     ei.AUTOGENERATED.String(),
		Category:  "Heartbeat",
		Hostname:  "Multiple",
		Timestamp: timestamp,
		Data:      event.Data,
	}, nil
}

func translateBeacon(event salt.Event) (Event, error) {

	log.Printf("Identified salt beacon event, tag: %s", event.Tag)

	tag := strings.Split(event.Tag, "/")
	hostname := strings.Split(tag[2], "-")

	return Event{
		Building:  hostname[0],
		Room:      hostname[1],
		Cause:     ei.AUTOGENERATED.String(),
		Category:  isHeartbeat(event),
		Hostname:  tag[2],
		Timestamp: tag[4],
		Data:      event.Data,
	}, nil
}

func isHeartbeat(event salt.Event) string {
	return "Heartbeat"
}
