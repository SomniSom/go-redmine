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
		fmt.Println("\n", user)
		for _, issue := range issues {
			fmt.Println(issue.AssignedTo.Id, issue.GetTitle(), issue.DoneRatio, issue.EstimatedHours, issue.SpentHours)
		}
	}

}
