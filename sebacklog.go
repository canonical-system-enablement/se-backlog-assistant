package main

import (
	"github.com/VojtechVitek/go-trello"
	"strings"
)

// SE Backlog board ID
const backlogBoardId string = "57c986881996509dca1b0f4d"
const tillamookBoardId string = "5a575ac018ab97ee883a9c87"

// Written as a single line of text, not a slice, to be able to leverage the
// strings.Contains(string, substring). Otherwise it would require writing own
// function. It is safe because there are no swimlanes with names that could
// cause collision.
var backlogLists = "Dartboard Next Sprint Candidates Yes Please Maybe Someday"
var tillamookLists = "Backlog Next Current In Review Done Blocked"

type SeBacklog struct {
	board trello.Board
}

func (t *SeClient) Backlog() (*SeBacklog, error) {
	board, err := t.trello.Board(backlogBoardId)
	if err != nil {
		return nil, err
	}

	return &SeBacklog{board: *board}, nil
}

func (t *SeClient) Tillamook() (*SeBacklog, error) {
	board, err := t.trello.Board(tillamookBoardId)
	if err != nil {
		return nil, err
	}

	return &SeBacklog{board: *board}, nil
}

func (sb *SeBacklog) Stories() (BacklogStories, error) {
	lists, err := sb.board.Lists()
	if err != nil {
		return nil, err
	}

	retval := BacklogStories{}

	for _, l := range lists {
		if !strings.Contains(backlogLists, l.Name) {
			continue
		}

		cards, err := l.Cards()
		if err != nil {
			return nil, err
		}

		for i, c := range cards {
			bc := CreateBacklogEntry(l.Name, uint(i), c)
			retval = append(retval, bc)
		}
	}

	return retval, nil
}
