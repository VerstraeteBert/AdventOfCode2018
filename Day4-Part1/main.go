package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type event struct {
	guardId int
	timestamp time.Time
	kind eventKind
}

type eventKind byte

const (
	eventStart eventKind = iota
	eventSleep
	eventAwake
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Failed to open input file")
	}

	lines := strings.Split(string(file), "\n")
	sort.Strings(lines)

	var events []event

	for _, line := range lines {
		dateEnd := strings.Index(line, "]")
		date, err := time.Parse("2006-01-02 15:04", line[1:dateEnd])
		if err != nil {
			log.Fatalf("Failed to parse date: %v", err)
		}

		var event event
		event.timestamp = date

		rest := strings.Fields(line[dateEnd + 1:])
		switch rest[0] {
			case "Guard":
				guardId, err := strconv.Atoi(rest[1][1:])
				if err != nil {
					log.Fatalf("Failed to parse guard id to int")
				}
				event.guardId = guardId
				event.kind = eventStart
				break
			case "falls":
				event.guardId = events[len(events) - 1].guardId
				event.kind = eventSleep
				break
			case "wakes":
				event.guardId = events[len(events) - 1].guardId
				event.kind = eventAwake
		}

		events = append(events, event)
	}

	guardId := findMostAsleep(events)
	mostAsleepMinute := findMostAsleepMinute(events, guardId)

	fmt.Printf("Guard %d was asleep the most, mostly on minute %d", guardId, mostAsleepMinute)
	fmt.Println("")
}

func findMostAsleepMinute(events []event, guardId int) int {
	minuteMap := make([]int, 60)

	for idx, event := range events {
		if event.guardId != guardId || event.kind != eventAwake {
			continue
		}

		for i := events[idx - 1].timestamp.Minute(); i < event.timestamp.Minute(); i++ {
			minuteMap[i]++
		}
	}

	return maxMinute(minuteMap)
}

func maxMinute(minutes []int) (maxMin int) {
	var maxMins int

	for min, times := range minutes {
		if times > maxMins {
			maxMin = min
			maxMins = times
		}
	}

	return
}

func findMostAsleep(events []event) (id int) {
	sleepTimes := make(map[int]time.Duration)

	for idx, event := range events {
		if event.kind == eventAwake {
			if events[idx - 1].kind != eventSleep || events[idx - 1].guardId != event.guardId {
				log.Fatal("Guard woke up when he wasn't asleep or guard id doesn't match")
			}
			sleepTimes[event.guardId] += event.timestamp.Sub(events[idx - 1].timestamp)
		}
	}

	return maxDuration(sleepTimes)
}

func maxDuration(durations map[int]time.Duration) (guardId int) {
	var currMaxDuration time.Duration

	for currGuardId, duration := range durations {
		if duration > currMaxDuration {
			currMaxDuration = duration
			guardId = currGuardId
		}
	}

	return
}

