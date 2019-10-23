//go:generate easyjson --all
package redmine

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//easyjson:json
type projectRequest struct {
	Project Project `json:"project"`
}

//easyjson:json
type projectResult struct {
	Project Project `json:"project"`
}

//easyjson:json
type projectsResult struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Identifier  string `json:"identifier"`
	Description string `json:"description"`
	CreatedOn   string `json:"created_on"`
	UpdatedOn   string `json:"updated_on"`
}

func (c *Client) Project(id int) (*Project, error) {
	statusCode, body, err := c.fhttp.Get(nil, c.endpoint+"/projects/"+strconv.Itoa(id)+".json?key="+c.apikey)
	//res, err := c.Get(c.endpoint + "/projects/" + strconv.Itoa(id) + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	//defer printError(res.Body.Close())

	//decoder := json.NewDecoder(res.Body)
	var r projectResult
	if statusCode != 200 {
		var er errorsResult
		err = json.Unmarshal(body, &er) //decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = json.Unmarshal(body, &r) //decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Project, nil
}

func (c *Client) Projects() ([]Project, error) {
	fmt.Println(c.endpoint + "/projects.json?key=" + c.apikey + c.getPaginationClause())
	statusCode, body, err := c.fhttp.Get(nil, c.endpoint+"/projects.json?key="+c.apikey+c.getPaginationClause())
	if err != nil {
		return nil, err
	}

	var r projectsResult
	if statusCode != 200 {
		var er errorsResult
		err = json.Unmarshal(body, &er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = json.Unmarshal(body, &r)
	}

	return r.Projects, err
}

func (c *Client) CreateProject(project Project) (*Project, error) {
	var ir projectRequest
	ir.Project = project
	s, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/projects.json?key="+c.apikey, strings.NewReader(string(s)))
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
	var r projectRequest
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
	return &r.Project, nil
}

func (c *Client) UpdateProject(project Project) error {
	var ir projectRequest
	ir.Project = project
	s, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", c.endpoint+"/projects/"+strconv.Itoa(project.Id)+".json?key="+c.apikey, strings.NewReader(string(s)))
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

func (c *Client) DeleteProject(id int) error {
	req, err := http.NewRequest("DELETE", c.endpoint+"/projects/"+strconv.Itoa(id)+".json?key="+c.apikey, strings.NewReader(""))
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
