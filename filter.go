package redmine

import (
	"strings"
	"time"
)

type Filter struct {
	UserID    string
	ProjectID string
	SpentOn   time.Time
	From      time.Time
	To        time.Time
	filters   map[string]string
}

//NewFilter
func NewFilter(args ...string) *Filter {
	f := &Filter{}
	if len(args)%2 == 0 {
		for i := 0; i < len(args); i += 2 {
			f.AddPair(args[i], args[i+1])
		}
	}
	return f
}

func (f *Filter) AddPair(key, value string) {
	if f.filters == nil {
		f.filters = make(map[string]string)
	}
	f.filters[key] = encode4Redmine(value)
}

func (f *Filter) ToURLParams() string {
	builder := strings.Builder{}
	if !f.SpentOn.IsZero() {
		builder.WriteString("&spent_on=" + f.SpentOn.Format("2006-01-02"))
	}

	if f.UserID != "" {
		builder.WriteString("&user_id=" + f.UserID)
	}

	if f.ProjectID != "" {
		builder.WriteString("&project_id=" + f.ProjectID)
	}

	if !f.From.IsZero() {
		builder.WriteString("&from=" + f.From.Format("2006-01-02"))
	}

	if !f.To.IsZero() {
		builder.WriteString("&to=" + f.To.Format("2006-01-02"))
	}

	for k, v := range f.filters {
		builder.WriteString("&" + k + "=" + v)
	}
	return builder.String()
}

func encode4Redmine(s string) string {
	a := strings.Replace(s, ">", "%3E", -1)
	a = strings.Replace(a, "<", "%3C", -1)
	a = strings.Replace(a, "=", "%3D", -1)
	return a
}
