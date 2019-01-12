package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/hugoamvieira/intercom-users-inviter/config"
	"github.com/hugoamvieira/intercom-users-inviter/data"
)

// Inviter reads through users and outputs them if they are within a certain distance
// from Intercom's office.
type Inviter struct {
	config *config.Config
	users  []*data.User
}

// NewInviter creates the Inviter instance. You must pass the file path for the JSON file,
// the distance threshold in km and the coordinates to check users distance against.
func NewInviter(c *config.Config) *Inviter {
	return &Inviter{
		config: c,
		users:  make([]*data.User, 0),
	}
}

// Start will kick-off the whole process
func (i *Inviter) Start() error {
	// Retrieve the data from file and put it in a queue
	fr := data.NewFileReader(i.config.UsersFilePath)
	err := fr.ReadLinesToQueue()
	if err != nil {
		return err
	}

	for {
		lineBytes, err := fr.Queue.Dequeue()
		if err == data.ErrEmptyQueue {
			// Parsed all users, stop dequeuing.
			break
		}
		if err != nil {
			// This is here because in the future someone could add more error cases to the queue
			// and forget to check them here. This covers that the program will not crash and burn.
			log.Println("Unhandled error captured | Error:", err)
			return err
		}

		user, err := data.NewUserFromJSONBytes(lineBytes)
		if err != nil {
			log.Println("Failed to create new user from JSON line, continuing to next user | Error:", err)
			continue
		}

		isWithin, err := user.IsWithinDistanceFromCoords(i.config.Lat, i.config.Lng, i.config.DistThresholdKm)
		if err != nil {
			log.Println("Failed to determine if user is within distance threshold, continuing to next user | Error:", err)
			continue
		}

		if isWithin {
			i.users = append(i.users, user)
		}
	}

	i.sort()
	i.output()

	return nil
}

func (i *Inviter) sort() {
	sort.Slice(i.users, func(a, b int) bool {
		return i.users[a].ID < i.users[b].ID
	})
}

func (i *Inviter) output() {
	fmt.Println("Users to Invite:")
	for _, user := range i.users {
		fmt.Printf("%v, ID %v\n", user.Name, user.ID)
	}
}
