//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"errors"
	"strings"
)

//easyjson:json
type issuePrioritiesResult struct {
	IssuePriorities []IssuePriority `json:"issue_priorities"`
}

//easyjson:json
type IssuePriority struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"is_default"`
}

func (c *Client) IssuePriorities() ([]IssuePriority, error) {
	res, err := c.Get(c.endpoint + "/enumerations/issue_priorities.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	defer printError(res.Body.Close())

	decoder := json.NewDecoder(res.Body)
	var r issuePrioritiesResult
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
	return r.IssuePriorities, nil
}
