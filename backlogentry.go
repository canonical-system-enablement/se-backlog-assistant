package main

import (
	"github.com/VojtechVitek/go-trello"
	"log"
	"strconv"
	"strings"
	"time"
)

type BacklogEntry struct {
	Position    uint     `json:"position"`
	Title       string   `json:"title"`
	List        string   `json:"list"`
	Points      uint     `json:"points"`
	Age         uint     `json:"age"`
	Stakeholder string   `json:"stakeholder"`
	Labels      []string `json:"labels"`
}

type BacklogStories []BacklogEntry

func CreateBacklogEntry(list string, pos uint, card trello.Card) BacklogEntry {
	be := BacklogEntry{}

	be.Position = pos
	be.List = list
	be.Points = getPointsFromTitle(card.Name)
	be.Stakeholder = getStakeholderFromTitle(card.Name)
	be.Title = getPrettyTitleFromTitle(card.Name)

	// labels
	for _, label := range card.Labels {
		be.Labels = append(be.Labels, label.Name)
	}

	// age
	epoch, err := strconv.ParseInt(card.Id[:8], 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	t0 := time.Unix(epoch, 0)
	be.Age = uint(time.Now().Sub(t0).Hours()/24 + 0.5)

	return be
}

func getPointsFromTitle(title string) uint {
	closingBracket := strings.Index(title, ")")
	if closingBracket == -1 || closingBracket > 4 {
		return 0
	}
	retval, err := strconv.Atoi(title[1:closingBracket])
	if err != nil {
		return 0
	}
	return uint(retval)
}

// Card stakeholder is tracked as:
// 'As <stakeholder> I want|would REST OF THE TITLE'"
func getStakeholderFromTitle(title string) string {
	closingBracket := strings.Index(title, ")")
	if closingBracket != -1 && closingBracket < 5 {
		title = title[closingBracket+2:]
	}
	tmp := strings.Split(strings.TrimPrefix(title, "As "), " I ")

	if len(tmp[0]) > 12 {
		return tmp[0][:12]
	} else {
		return strings.TrimRight(tmp[0], ", ")
	}
}

// Card points are tracked as '(points) REST OF TITLE'"
func get(Name string) (string, string) {
	snapName := ""
	snapVersion := "null"

	tmp := strings.Split(Name, " - ")

	if len(tmp) > 0 {
		snapName = tmp[0]
	}

	if len(tmp) > 1 {
		snapVersion = tmp[1]
	}

	return snapName, snapVersion
}

func getPrettyTitleFromTitle(title string) string {
	tmp := strings.Split(title, "I want")
	retval := ""

	if len(tmp) == 1 {
		tmp = strings.Split(title, "I would like")
		if len(tmp) == 1 {
			retval = tmp[0]
		} else {
			retval = tmp[1]
		}
	} else {
		retval = tmp[1]
	}

	if len(retval) > 80 {
		return retval[:80]
	} else {
		return retval
	}
}
