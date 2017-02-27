SE Backlog Assistant
====================

This utility gets the SE Backlog and prints the stories there on the
screen. It is possible to limit the results to a particular lane,
stakeholder or label.

Building
========

It uses the:
 - "github.com/VojtechVitek/go-trello" Trello library for Go.
 - "gopkg.in/ahmetalpbalkan/go-linq.v3" LINQ for Go

You can install them with:

```
go get github.com/VojtechVitek/go-trello
go get gopkg.in/ahmetalpbalkan/go-linq.v3
```

Requirements
============

trello_secrets.json file that contains the Application Key and Token.

Follow the instructions here https://trello.com/app-key to obtain an API
key and token.


```
konrad at annapurna in snap-release-email-generator (master) % cat trello_secrets.json 
{"app_id":"APP ID","token":"TOKEN"}

```


Usage
=====

```
Usage of ./se-backlog-assistant:
  -label string
        Limit the output to a single label (default "missing")
  -list string
        Limit the output to a single list (default "missing")
  -secrets string
        Trello Secrets configuration (default "trello_secrets.json")
  -stakeholder string
        Limit the output to a single stakeholder (default "missing")
```

Output
======

It prints information such as:
 - story position
 - the lane the story is at
 - the stakeholder
 - age of the story in days
 - the story topic
 - list of labels assigned

Note that the topic is prettyized by removing the opening "As <stakeholder>
I want" sentence. Also the ending "so that <explanation>" part is dropped.

Exampe output:

```
 0  | Dartboard | Stakeholder | 6d   |  story topic | Label                               
 1  | Dartboard | Stakeholder | 23d  |  story topic | Label, Label                     
```
