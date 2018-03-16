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
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var (
	trelloSecretsFile = flag.String("secrets", "trello_secrets.json", "Trello Secrets configuration")

	limitList        = flag.String("list", "missing", "Limit the output to a single list")
	limitLabel       = flag.String("label", "missing", "Limit the output to a single label")
	limitStakeholder = flag.String("stakeholder", "missing", "Limit the output to a single stakeholder")
)

const (
	database = "testdb"
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

	backlog, err := trello.Tillamook()
	if err != nil {
		log.Fatal(err)
	}

	lists, err := backlog.board.Lists()
	if err != nil {
		log.Fatal("Cannot get lists")
	}

	t := time.Now()

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	labels := map[string]int{}

	for _, l := range lists {
		if !strings.Contains(tillamookLists, l.Name) {
			continue
		}

		cards, err := l.Cards()
		if err != nil {
			log.Fatal("Cannot get cards")
		}

		tags := map[string]string{"lane": strings.Replace(strings.ToLower(l.Name), " ", "", -1)}
		fields := map[string]interface{}{
			"value": len(cards),
		}

		pt, err := client.NewPoint("daily-lists", tags, fields, t)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		for _, c := range cards {
			if l.Name == "Done" {
				continue
			}
			for _, ll := range c.Labels {
				labels[ll.Name]++
			}
		}

		fmt.Printf("insert daily,lane=%s value=%d %d\n", strings.Replace(strings.ToLower(l.Name), " ", "", -1), len(cards), t.UnixNano())
	}

	for key, item := range labels {
		tags2 := map[string]string{"label": strings.Replace(strings.ToLower(key), " ", "-", -1)}
		fields2 := map[string]interface{}{
			"value": item,
		}
		pt2, err := client.NewPoint("daily-labels", tags2, fields2, t)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt2)

	}

	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// insert daily2,lane=backlog value=20 1504225728000123456
}
