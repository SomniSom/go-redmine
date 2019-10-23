//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

//easyjson:json
type trackersResult struct {
	Trackers []IdName `json:"trackers"`
}

func (c *Client) Trackers() ([]IdName, error) {
	res, err := http.Get(c.endpoint + "/trackers.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	defer printError(res.Body.Close())

	decoder := json.NewDecoder(res.Body)
	var r trackersResult
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
	return r.Trackers, nil
}
