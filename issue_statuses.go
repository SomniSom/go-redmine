//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"errors"
	"strings"
)

//easyjson:json
type issueStatusesResult struct {
	IssueStatuses []IssueStatus `json:"issue_statuses"`
}

//easyjson:json
type IssueStatus struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
	IsClosed  bool   `json:"is_closed"`
}

func (c *Client) IssueStatuses() ([]IssueStatus, error) {
	res, err := c.Get(c.endpoint + "/issue_statuses.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	defer printError(res.Body.Close())

	decoder := json.NewDecoder(res.Body)
	var r issueStatusesResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return r.IssueStatuses, nil
}
