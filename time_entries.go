//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
	"strings"
)

//easyjson:json
type timeEntriesResult struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}

//easyjson:json
type timeEntryResult struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

//easyjson:json
type timeEntryRequest struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

//easyjson:json
type TimeEntry struct {
	Id           int            `json:"id"`
	Project      IdName         `json:"project,omitempty"`
	Issue        Id             `json:"issue"`
	User         IdName         `json:"user,omitempty"`
	Activity     IdName         `json:"activity,omitempty"`
	Hours        float32        `json:"hours"`
	Comments     string         `json:"comments,omitempty"`
	SpentOn      string         `json:"spent_on,omitempty"`
	CreatedOn    string         `json:"created_on,omitempty"`
	UpdatedOn    string         `json:"updated_on,omitempty"`
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}

// TimeEntriesWithFilter send query and return parsed result
func (c *Client) TimeEntriesWithFilter(filter Filter) ([]TimeEntry, error) {
	uri, err := c.URLWithFilter("/time_entries.json", filter)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	req.URI().Update(uri)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	res := fasthttp.AcquireResponse()

	err = c.fhttp.Do(req, res)
	if err != nil {
		return nil, err
	}

	var r timeEntriesResult
	if res.StatusCode() == 404 {
		return nil, errors.New("Not Found ")
	}

	if res.StatusCode() != 200 {
		var er errorsResult
		err = json.Unmarshal(res.Body(), &er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = json.Unmarshal(res.Body(), &r)
	}
	if err != nil {
		return nil, err
	}
	return r.TimeEntries, nil
}

func (c *Client) TimeEntries(projectId int) ([]TimeEntry, error) {
	req := fasthttp.AcquireRequest()
	req.URI().Update(c.endpoint + "/projects/" + strconv.Itoa(projectId) + "/time_entries.json?key=" + c.apikey + c.getPaginationClause())
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	res := fasthttp.AcquireResponse()

	err := c.fhttp.Do(req, res)
	if err != nil {
		return nil, err
	}

	var r timeEntriesResult
	if res.StatusCode() == 404 {
		return nil, errors.New("Not Found ")
	}
	if res.StatusCode() != 200 {
		var er errorsResult
		err = json.Unmarshal(res.Body(), &er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = json.Unmarshal(res.Body(), &r)
	}
	if err != nil {
		return nil, err
	}
	return r.TimeEntries, nil
}

func (c *Client) TimeEntry(id int) (*TimeEntry, error) {
	statusCode, body, err := c.fhttp.Get(nil, c.endpoint+"/time_entries/"+strconv.Itoa(id)+".json?key="+c.apikey)
	if err != nil {
		return nil, err
	}

	var r timeEntryResult
	if statusCode != 200 {
		if statusCode == 404 {
			return nil, errors.New("Not Found ")
		}
		var er errorsResult
		err = json.Unmarshal(body, &er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = json.Unmarshal(body, &r)
	}
	if err != nil {
		return nil, err
	}
	return &r.TimeEntry, nil
}

func (c *Client) CreateTimeEntry(timeEntry TimeEntry) (*TimeEntry, error) {
	var ir timeEntryRequest
	ir.TimeEntry = timeEntry
	s, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/time_entries.json?key="+c.apikey, strings.NewReader(string(s)))
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
	var r timeEntryResult
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
	return &r.TimeEntry, nil
}

func (c *Client) UpdateTimeEntry(timeEntry TimeEntry) error {
	var ir timeEntryRequest
	ir.TimeEntry = timeEntry
	s, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", c.endpoint+"/time_entries/"+strconv.Itoa(timeEntry.Id)+".json?key="+c.apikey, strings.NewReader(string(s)))
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

func (c *Client) DeleteTimeEntry(id int) error {
	req, err := http.NewRequest("DELETE", c.endpoint+"/time_entries/"+strconv.Itoa(id)+".json?key="+c.apikey, strings.NewReader(""))
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
