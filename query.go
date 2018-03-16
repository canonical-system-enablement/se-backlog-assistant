package main

import (
	. "gopkg.in/ahmetalpbalkan/go-linq.v3"
	"strings"
)

type BacklogQuery struct{}

func (q *BacklogQuery) Limit(stories BacklogStories, list string, label string, stakeholder string) BacklogStories {
	var retval = stories
	if list != "missing" {
		retval = q.LimitList(retval, list)
	}
	if label != "missing" {
		retval = q.limitLabel(retval, label)
	}
	if stakeholder != "missing" {
		retval = q.limitStakeholder(retval, stakeholder)
	}
	return retval
}

func (q *BacklogQuery) LimitList(stories BacklogStories, list string) BacklogStories {
	var retval BacklogStories
	From(stories).Where(func(c interface{}) bool {
		return strings.Contains(c.(BacklogEntry).List, list)
	}).ToSlice(&retval)
	return retval
}
func (q *BacklogQuery) limitStakeholder(stories BacklogStories, stakeholder string) BacklogStories {
	var retval BacklogStories
	From(stories).Where(func(c interface{}) bool {
		return strings.Contains(c.(BacklogEntry).Stakeholder, stakeholder)
	}).ToSlice(&retval)
	return retval
}
func (q *BacklogQuery) limitLabel(stories BacklogStories, label string) BacklogStories {
	var retval BacklogStories
	From(stories).Where(func(c interface{}) bool {
		return strings.Contains(strings.Join(c.(BacklogEntry).Labels, " "), label)
	}).ToSlice(&retval)
	return retval
}
