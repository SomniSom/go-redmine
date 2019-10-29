package main

import (
	"flag"
	"fmt"
	"github.com/SomniSom/go-redmine"
	"github.com/ztrue/tracerr"
	"strconv"
	"strings"
)

var (
	endpoint      string
	key           string
	users         string
	acceptedUsers []int
)

func init() {
	flag.StringVar(&endpoint, "e", "", "Endpoint Redmine")
	flag.StringVar(&key, "k", "", "Api key Redmine")
	flag.StringVar(&users, "u", "", "Accepted users")
}

func acceptUser(userId int) bool {
	if len(acceptedUsers) == 0 {
		return true
	}

	for _, au := range acceptedUsers {
		if au == userId {
			return true
		}
	}
	return false
}

func timePercent(done int, estimate, spent float32) (bool, int) {
	if (estimate - spent - float32(done)) == 0 {
		return true, 0
	}
	if estimate == 0 {
		return false, 0
	}
	proc := int((spent / estimate) * 100)
	return done >= (proc - 10), proc
}

func main() {
	flag.Parse()
	if endpoint == "" || key == "" {
		flag.PrintDefaults()
		return
	}

	for _, userId := range strings.Split(users, ",") {
		if id, err := strconv.Atoi(userId); err == nil {
			acceptedUsers = append(acceptedUsers, id)
		}
	}

	cl := redmine.NewClient(endpoint, key)
	cl.Offset = -1
	cl.Limit = 100

	issues, err := cl.IssuesByFilter(&redmine.IssueFilter{StatusId: "2"})
	if err != nil {
		tracerr.PrintSourceColor(err)
		tracerr.StackTrace(err)
		return
	}

	var data = make(map[string][]*redmine.Issue)

	for _, issue := range issues {
		if issue.AssignedTo != nil {
			if !acceptUser(issue.AssignedTo.Id) {
				continue
			}
			issue.Update(cl)
			if issue == nil {
				continue
			}
			if _, ok := data[issue.AssignedTo.Name]; !ok {
				data[issue.AssignedTo.Name] = make([]*redmine.Issue, 0, 3)
			}
			data[issue.AssignedTo.Name] = append(data[issue.AssignedTo.Name], issue)
		}
	}

	for user, issues := range data {
		fmt.Println("\n", user, issues[0].AssignedTo.Id)
		for _, issue := range issues {
			correct, tPercent := timePercent(issue.DoneRatio, issue.EstimatedHours, issue.SpentHours)
			fmt.Printf("%v Done: %v(%v) Estimate: %v Spent: %v Correct: %v\n", issue.GetTitle(), issue.DoneRatio, tPercent, issue.EstimatedHours, issue.SpentHours, correct)
		}
	}

}
