//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

//easyjson:json
type issueRelationsResult struct {
	IssueRelations []IssueRelation `json:"relations"`
}

//easyjson:json
type issueRelationResult struct {
	IssueRelation IssueRelation `json:"issue_relation"`
}

//easyjson:json
type issueRelationRequest struct {
	IssueRelation IssueRelation `json:"issue_relation"`
}

//easyjson:json
type IssueRelation struct {
	Id           int    `json:"id"`
	IssueId      string `json:"issue_id"`
	IssueToId    string `json:"issue_to_id"`
	RelationType string `json:"relation_type"`
	Delay        string `json:"delay"`
}

func (c *Client) IssueRelations(issueId int) ([]IssueRelation, error) {
	res, err := c.Get(c.endpoint + "/issue/" + strconv.Itoa(issueId) + "/relations.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	defer printError(res.Body.Close())

	decoder := json.NewDecoder(res.Body)
	var r issueRelationsResult
	if res.StatusCode == 404 {
		return nil, errors.New("Not Found ")
	}
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
	return r.IssueRelations, nil
}

func (c *Client) IssueRelation(id int) (*IssueRelation, error) {
	res, err := c.Get(c.endpoint + "/relations/" + strconv.Itoa(id) + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer printError(res.Body.Close())

	decoder := json.NewDecoder(res.Body)
	var r issueRelationResult
	if res.StatusCode == 404 {
		return nil, errors.New("Not Found ")
	}
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
	return &r.IssueRelation, nil
}

func (c *Client) CreateIssueRelation(issueRelation IssueRelation) (*IssueRelation, error) {
	var ir issueRelationRequest
	ir.IssueRelation = issueRelation
	s, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/relations.json?key="+c.apikey, strings.NewReader(string(s)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer printError(res.Body.Close())

	decoder := json.NewDecoder(res.Body)
	var r issueRelationResult
	if res.StatusCode != 201 {
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
	return &r.IssueRelation, nil
}

func (c *Client) UpdateIssueRelation(issueRelation IssueRelation) error {
	var ir issueRelationRequest
	ir.IssueRelation = issueRelation
	s, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", c.endpoint+"/relations/"+strconv.Itoa(issueRelation.Id)+".json?key="+c.apikey, strings.NewReader(string(s)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer printError(res.Body.Close())

	if res.StatusCode == 404 {
		return errors.New("Not Found ")
	}
	if res.StatusCode != 200 {
		decoder := json.NewDecoder(res.Body)
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	}
	if err != nil {
		return err
	}
	return err
}

func (c *Client) DeleteIssueRelation(id int) error {
	req, err := http.NewRequest("DELETE", c.endpoint+"/relations/"+strconv.Itoa(id)+".json?key="+c.apikey, strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer printError(res.Body.Close())

	if res.StatusCode == 404 {
		return errors.New("Not Found ")
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	}
	return err
}
