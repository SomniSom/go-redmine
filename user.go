//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"strconv"
)

//easyjson:json
type userResult struct {
	User User `json:"user"`
}

//easyjson:json
type usersResult struct {
	Users []User `json:"users"`
}

//easyjson:json
type User struct {
	Id           int            `json:"id"`
	Login        string         `json:"login"`
	Firstname    string         `json:"firstname"`
	Lastname     string         `json:"lastname"`
	Mail         string         `json:"mail"`
	CreatedOn    string         `json:"created_on"`
	LatLoginOn   string         `json:"last_login_on"`
	Memberships  []Membership   `json:"memberships"`
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}

func (c *Client) Users() ([]User, error) {
	statusCode, body, err := c.fhttp.Get(nil, c.endpoint+"/users.json?key="+c.apikey+c.getPaginationClause())
	if err != nil {
		return nil, err
	}

	var r usersResult
	if statusCode != 200 {
		return nil, checkStatusCode(statusCode, body)
	} else {
		err = json.Unmarshal(body, &r)
	}

	return r.Users, err
}

func (c *Client) User(id int) (*User, error) {
	statusCode, body, err := c.fhttp.Get(nil, c.endpoint+"/users/"+strconv.Itoa(id)+".json?key="+c.apikey)
	if err != nil {
		return nil, err
	}

	var r userResult
	if statusCode != 200 {
		return nil, checkStatusCode(statusCode, body)
	} else {
		err = json.Unmarshal(body, &r)
	}
	if err != nil {
		return nil, err
	}
	return &r.User, nil
}
