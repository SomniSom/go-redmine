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
	for _, au := range acceptedUsers {
		if au == userId {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	if endpoint == "" || key == "" || users == "" {
		flag.PrintDefaults()
		return
	}

	for _, userId := range strings.Split(users, ",") {
		if id, err := strconv.Atoi(userId); err == nil {
			acceptedUsers = append(acceptedUsers, id)
		}
	}
	if len(acceptedUsers) == 0 {
		fmt.Println("Accepted users not found")
		return
	}

	cl := redmine.NewClient(endpoint, key)
	cl.Offset = -1
	cl.Limit = 100
	//fmt.Println(cl.Projects())
	//fmt.Println(cl.Users())
	issues, err := cl.IssuesByFilter(&redmine.IssueFilter{StatusId: "2"})
	if err != nil {
		tracerr.PrintSourceColor(err)
		tracerr.StackTrace(err)
		return
	}

	var data []*redmine.Issue

	for _, issue := range issues {
		if issue.AssignedTo != nil {
			if !acceptUser(issue.AssignedTo.Id) {
				continue
			}
			data = append(data, &issue)
			fmt.Println(issue.GetTitle(), " == ", issue.AssignedTo.Name, issue.AssignedTo.Id)
		}
	}
}
