package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/albanul/advent_of_code_2024/day05"
)

type Rules map[int][]int
type Update []int

func main() {
	rules, updates, err := GetRulesAndUpdatesFromFile("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := SumCorrectedMiddlePages(rules, updates)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum: ", sum)
}

func GetRulesAndUpdatesFromFile(filepath string) (Rules, []Update, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	rules, pages, err := GetRulesAndUpdates(file)
	if err != nil {
		return nil, nil, err
	}

	return rules, pages, nil
}

func GetRulesAndUpdates(r io.Reader) (Rules, []Update, error) {
	scanner := bufio.NewScanner(r)

	rules := make(Rules)

	// parse rules
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		var x, y int

		_, err := fmt.Sscanf(line, "%d|%d", &x, &y)
		if err != nil {
			return nil, nil, err
		}

		if _, ok := rules[x]; !ok {
			rules[x] = make([]int, 0)
		}
		rules[x] = append(rules[x], y)
	}

	updates := make([]Update, 0)

	// parse pages
	for scanner.Scan() {
		line := scanner.Text()

		pages := make([]int, 0)

		split := strings.Split(line, ",")

		for _, s := range split {
			page, err := strconv.Atoi(s)
			if err != nil {
				return nil, nil, err
			}

			pages = append(pages, page)
		}

		updates = append(updates, pages)
	}

	return rules, updates, nil
}

func SumCorrectedMiddlePages(rules Rules, updates []Update) (int, error) {
	incorrectUpdates, err := GetIncorrectUpdates(rules, updates)
	if err != nil {
		return 0, err
	}

	correctedUpdates := make([]Update, 0)

	for _, incorrectUpdate := range incorrectUpdates {
		correctedUpdate, err := CorrectUpdate(rules, incorrectUpdate)
		if err != nil {
			return 0, err
		}

		correctedUpdates = append(correctedUpdates, correctedUpdate)
	}

	sum := 0
	for _, correctUpdate := range correctedUpdates {
		middleIndex := len(correctUpdate) / 2
		sum += correctUpdate[middleIndex]
	}

	return sum, nil
}

func GetIncorrectUpdates(rules Rules, updates []Update) ([]Update, error) {
	incorrectUpdates := make([]Update, 0)

	for _, update := range updates {
		isCorrect := true

		filteredRules := FilterRules(rules, update)

		for i := 0; i < len(update)-1; i++ {
			ok, err := CanFindRule(filteredRules, update[i], update[i+1])
			if err != nil {
				return nil, err
			}

			if !ok {
				isCorrect = false
				break
			}
		}

		if isCorrect {
			continue
		}

		incorrectUpdates = append(incorrectUpdates, update)
	}

	return incorrectUpdates, nil
}

func CorrectUpdate(rules Rules, origUpdate Update) (Update, error) {
	filteredRules := FilterRules(rules, origUpdate)

	update := make(Update, len(origUpdate))
	copy(update, origUpdate)

	for j := 1; j < len(update); j++ {
		isCorrect := true

		for i := 0; i < len(update)-j; i++ {
			ok, err := CanFindRule(filteredRules, update[i], update[i+1])
			if err != nil {
				return nil, err
			}

			if !ok {
				update[i], update[i+1] = update[i+1], update[i]
				isCorrect = false
			}
		}

		if isCorrect {
			break
		}
	}
	return update, nil
}

func FilterRules(rules Rules, updates Update) Rules {
	filteredRules := make(Rules)

	for _, elem := range updates {
		links, ok := rules[elem]
		if !ok {
			continue
		}

		filteredLinks := make([]int, 0)

		for _, link := range links {
			if slices.Contains(updates, link) {
				filteredLinks = append(filteredLinks, link)
			}
		}

		filteredRules[elem] = filteredLinks
	}

	return filteredRules
}

func CanFindRule(rules Rules, l, r int) (bool, error) {
	q := utils.NewQueue()
	q.Enqueue(l)

	for !q.IsEmpty() {
		next, err := q.Dequeue()
		if err != nil {
			return false, err
		}

		//fmt.Println(next)

		if next == r {
			return true, nil
		}

		links := rules[next]

		for _, ll := range links {
			q.Enqueue(ll)
		}
	}

	return false, nil
}
