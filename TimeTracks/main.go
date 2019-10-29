package main

import (
	"flag"
	"fmt"
	"github.com/SomniSom/go-redmine"
	"github.com/go-errors/errors"
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

	sliceTimeEntry, err := cl.TimeEntriesWithFilter(redmine.Filter{})
	if err != nil {
		fmt.Println(err.(*errors.Error).ErrorStack())
		return
	}

	for _, tEntry := range sliceTimeEntry {
		fmt.Println(tEntry, tEntry.SpentOn, tEntry.Hours)

	}

}
