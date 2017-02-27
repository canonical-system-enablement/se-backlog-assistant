/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	trelloSecretsFile = flag.String("secrets", "trello_secrets.json", "Trello Secrets configuration")

	limitList        = flag.String("list", "missing", "Limit the output to a single list")
	limitLabel       = flag.String("label", "missing", "Limit the output to a single label")
	limitStakeholder = flag.String("stakeholder", "missing", "Limit the output to a single stakeholder")
)

func Print(stories BacklogStories) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, s := range stories {
		fmt.Fprintf(w, " %d\t| %s\t| %s\t| %dd\t| %s\t| %s\t\n", s.Position, s.List, s.Stakeholder, s.Age, s.Title, strings.Join(s.Labels, ", "))
	}
	w.Flush()
}

func main() {

	flag.Parse()

	trelloSecrets, err := NewTrelloSecrets(*trelloSecretsFile)
	if err != nil {
		log.Fatal(err)
	}

	trello, err := NewSeClient(*trelloSecrets)
	if err != nil {
		log.Fatal(err)
	}

	backlog, err := trello.Backlog()
	if err != nil {
		log.Fatal(err)
	}

	sfb, err := backlog.Stories()
	if err != nil {
		log.Fatal(err)
	}

	bq := BacklogQuery{}
	sfb = bq.Limit(sfb, *limitList, *limitLabel, *limitStakeholder)
	Print(sfb)
}
