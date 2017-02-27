SE Backlog Assistant
====================

This utility gets the SE Backlog and prints the stories there on the
screen. It is possible to limit the results to a particular lane,
stakeholder or label.

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
