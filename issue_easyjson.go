// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package redmine

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine(in *jlexer.Lexer, out *issuesResult) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "issues":
			if in.IsNull() {
				in.Skip()
				out.Issues = nil
			} else {
				in.Delim('[')
				if out.Issues == nil {
					if !in.IsDelim(']') {
						out.Issues = make([]*Issue, 0, 8)
					} else {
						out.Issues = []*Issue{}
					}
				} else {
					out.Issues = (out.Issues)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Issue
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Issue)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Issues = append(out.Issues, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "total_count":
			out.TotalCount = uint(in.Uint())
		case "offset":
			out.Offset = uint(in.Uint())
		case "limit":
			out.Limit = uint(in.Uint())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine(out *jwriter.Writer, in issuesResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"issues\":"
		out.RawString(prefix[1:])
		if in.Issues == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Issues {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"total_count\":"
		out.RawString(prefix)
		out.Uint(uint(in.TotalCount))
	}
	{
		const prefix string = ",\"offset\":"
		out.RawString(prefix)
		out.Uint(uint(in.Offset))
	}
	{
		const prefix string = ",\"limit\":"
		out.RawString(prefix)
		out.Uint(uint(in.Limit))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v issuesResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v issuesResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *issuesResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *issuesResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine1(in *jlexer.Lexer, out *issueResult) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "issue":
			(out.Issue).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine1(out *jwriter.Writer, in issueResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"issue\":"
		out.RawString(prefix[1:])
		(in.Issue).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v issueResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v issueResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *issueResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *issueResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine1(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine2(in *jlexer.Lexer, out *issueRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "issue":
			(out.Issue).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine2(out *jwriter.Writer, in issueRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"issue\":"
		out.RawString(prefix[1:])
		(in.Issue).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v issueRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v issueRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *issueRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *issueRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine2(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine3(in *jlexer.Lexer, out *JournalDetails) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "property":
			out.Property = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "old_value":
			out.OldValue = string(in.String())
		case "new_value":
			out.NewValue = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine3(out *jwriter.Writer, in JournalDetails) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"property\":"
		out.RawString(prefix[1:])
		out.String(string(in.Property))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"old_value\":"
		out.RawString(prefix)
		out.String(string(in.OldValue))
	}
	{
		const prefix string = ",\"new_value\":"
		out.RawString(prefix)
		out.String(string(in.NewValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JournalDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JournalDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JournalDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JournalDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine3(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine4(in *jlexer.Lexer, out *Journal) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.User)
			}
		case "notes":
			out.Notes = string(in.String())
		case "created_on":
			out.CreatedOn = string(in.String())
		case "details":
			if in.IsNull() {
				in.Skip()
				out.Details = nil
			} else {
				in.Delim('[')
				if out.Details == nil {
					if !in.IsDelim(']') {
						out.Details = make([]JournalDetails, 0, 1)
					} else {
						out.Details = []JournalDetails{}
					}
				} else {
					out.Details = (out.Details)[:0]
				}
				for !in.IsDelim(']') {
					var v4 JournalDetails
					(v4).UnmarshalEasyJSON(in)
					out.Details = append(out.Details, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine4(out *jwriter.Writer, in Journal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		if in.User == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.User)
		}
	}
	{
		const prefix string = ",\"notes\":"
		out.RawString(prefix)
		out.String(string(in.Notes))
	}
	{
		const prefix string = ",\"created_on\":"
		out.RawString(prefix)
		out.String(string(in.CreatedOn))
	}
	{
		const prefix string = ",\"details\":"
		out.RawString(prefix)
		if in.Details == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Details {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Journal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Journal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Journal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Journal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine4(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in *jlexer.Lexer, out *IdName) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out *jwriter.Writer, in IdName) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine6(in *jlexer.Lexer, out *IssueFilter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ProjectId":
			out.ProjectId = string(in.String())
		case "SubprojectId":
			out.SubprojectId = string(in.String())
		case "TrackerId":
			out.TrackerId = string(in.String())
		case "StatusId":
			out.StatusId = string(in.String())
		case "AssignedToId":
			out.AssignedToId = string(in.String())
		case "UpdatedOn":
			out.UpdatedOn = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine6(out *jwriter.Writer, in IssueFilter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ProjectId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ProjectId))
	}
	{
		const prefix string = ",\"SubprojectId\":"
		out.RawString(prefix)
		out.String(string(in.SubprojectId))
	}
	{
		const prefix string = ",\"TrackerId\":"
		out.RawString(prefix)
		out.String(string(in.TrackerId))
	}
	{
		const prefix string = ",\"StatusId\":"
		out.RawString(prefix)
		out.String(string(in.StatusId))
	}
	{
		const prefix string = ",\"AssignedToId\":"
		out.RawString(prefix)
		out.String(string(in.AssignedToId))
	}
	{
		const prefix string = ",\"UpdatedOn\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedOn))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IssueFilter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IssueFilter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IssueFilter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IssueFilter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine6(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine7(in *jlexer.Lexer, out *Issue) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "subject":
			out.Subject = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "project_id":
			out.ProjectId = int(in.Int())
		case "project":
			if in.IsNull() {
				in.Skip()
				out.Project = nil
			} else {
				if out.Project == nil {
					out.Project = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.Project)
			}
		case "tracker_id":
			out.TrackerId = int(in.Int())
		case "tracker":
			if in.IsNull() {
				in.Skip()
				out.Tracker = nil
			} else {
				if out.Tracker == nil {
					out.Tracker = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.Tracker)
			}
		case "parent_issue_id":
			out.ParentId = int(in.Int())
		case "parent":
			if in.IsNull() {
				in.Skip()
				out.Parent = nil
			} else {
				if out.Parent == nil {
					out.Parent = new(Id)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine8(in, out.Parent)
			}
		case "status_id":
			out.StatusId = int(in.Int())
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.Status)
			}
		case "priority_id":
			out.PriorityId = int(in.Int())
		case "priority":
			if in.IsNull() {
				in.Skip()
				out.Priority = nil
			} else {
				if out.Priority == nil {
					out.Priority = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.Priority)
			}
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.Author)
			}
		case "fixed_version":
			if in.IsNull() {
				in.Skip()
				out.FixedVersion = nil
			} else {
				if out.FixedVersion == nil {
					out.FixedVersion = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.FixedVersion)
			}
		case "assigned_to":
			if in.IsNull() {
				in.Skip()
				out.AssignedTo = nil
			} else {
				if out.AssignedTo == nil {
					out.AssignedTo = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.AssignedTo)
			}
		case "category":
			if in.IsNull() {
				in.Skip()
				out.Category = nil
			} else {
				if out.Category == nil {
					out.Category = new(IdName)
				}
				easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine5(in, out.Category)
			}
		case "category_id":
			out.CategoryId = int(in.Int())
		case "notes":
			out.Notes = string(in.String())
		case "status_date":
			out.StatusDate = string(in.String())
		case "created_on":
			out.CreatedOn = string(in.String())
		case "updated_on":
			out.UpdatedOn = string(in.String())
		case "start_date":
			out.StartDate = string(in.String())
		case "due_date":
			out.DueDate = string(in.String())
		case "closed_on":
			out.ClosedOn = string(in.String())
		case "custom_fields":
			if in.IsNull() {
				in.Skip()
				out.CustomFields = nil
			} else {
				in.Delim('[')
				if out.CustomFields == nil {
					if !in.IsDelim(']') {
						out.CustomFields = make([]*CustomField, 0, 8)
					} else {
						out.CustomFields = []*CustomField{}
					}
				} else {
					out.CustomFields = (out.CustomFields)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *CustomField
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(CustomField)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.CustomFields = append(out.CustomFields, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "uploads":
			if in.IsNull() {
				in.Skip()
				out.Uploads = nil
			} else {
				in.Delim('[')
				if out.Uploads == nil {
					if !in.IsDelim(']') {
						out.Uploads = make([]*Upload, 0, 8)
					} else {
						out.Uploads = []*Upload{}
					}
				} else {
					out.Uploads = (out.Uploads)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *Upload
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(Upload)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					out.Uploads = append(out.Uploads, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "journals":
			if in.IsNull() {
				in.Skip()
				out.Journals = nil
			} else {
				in.Delim('[')
				if out.Journals == nil {
					if !in.IsDelim(']') {
						out.Journals = make([]*Journal, 0, 8)
					} else {
						out.Journals = []*Journal{}
					}
				} else {
					out.Journals = (out.Journals)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *Journal
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(Journal)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.Journals = append(out.Journals, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "spent_hours":
			out.SpentHours = float32(in.Float32())
		case "total_spent_hours":
			out.TotalSpentHours = float32(in.Float32())
		case "estimated_hours":
			out.EstimatedHours = float32(in.Float32())
		case "done_ratio":
			out.DoneRatio = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine7(out *jwriter.Writer, in Issue) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"subject\":"
		out.RawString(prefix)
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"project_id\":"
		out.RawString(prefix)
		out.Int(int(in.ProjectId))
	}
	if in.Project != nil {
		const prefix string = ",\"project\":"
		out.RawString(prefix)
		easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.Project)
	}
	{
		const prefix string = ",\"tracker_id\":"
		out.RawString(prefix)
		out.Int(int(in.TrackerId))
	}
	{
		const prefix string = ",\"tracker\":"
		out.RawString(prefix)
		if in.Tracker == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.Tracker)
		}
	}
	if in.ParentId != 0 {
		const prefix string = ",\"parent_issue_id\":"
		out.RawString(prefix)
		out.Int(int(in.ParentId))
	}
	{
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		if in.Parent == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine8(out, *in.Parent)
		}
	}
	{
		const prefix string = ",\"status_id\":"
		out.RawString(prefix)
		out.Int(int(in.StatusId))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		if in.Status == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.Status)
		}
	}
	if in.PriorityId != 0 {
		const prefix string = ",\"priority_id\":"
		out.RawString(prefix)
		out.Int(int(in.PriorityId))
	}
	{
		const prefix string = ",\"priority\":"
		out.RawString(prefix)
		if in.Priority == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.Priority)
		}
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		if in.Author == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.Author)
		}
	}
	{
		const prefix string = ",\"fixed_version\":"
		out.RawString(prefix)
		if in.FixedVersion == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.FixedVersion)
		}
	}
	{
		const prefix string = ",\"assigned_to\":"
		out.RawString(prefix)
		if in.AssignedTo == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.AssignedTo)
		}
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		if in.Category == nil {
			out.RawString("null")
		} else {
			easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine5(out, *in.Category)
		}
	}
	{
		const prefix string = ",\"category_id\":"
		out.RawString(prefix)
		out.Int(int(in.CategoryId))
	}
	{
		const prefix string = ",\"notes\":"
		out.RawString(prefix)
		out.String(string(in.Notes))
	}
	{
		const prefix string = ",\"status_date\":"
		out.RawString(prefix)
		out.String(string(in.StatusDate))
	}
	{
		const prefix string = ",\"created_on\":"
		out.RawString(prefix)
		out.String(string(in.CreatedOn))
	}
	{
		const prefix string = ",\"updated_on\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedOn))
	}
	{
		const prefix string = ",\"start_date\":"
		out.RawString(prefix)
		out.String(string(in.StartDate))
	}
	{
		const prefix string = ",\"due_date\":"
		out.RawString(prefix)
		out.String(string(in.DueDate))
	}
	{
		const prefix string = ",\"closed_on\":"
		out.RawString(prefix)
		out.String(string(in.ClosedOn))
	}
	if len(in.CustomFields) != 0 {
		const prefix string = ",\"custom_fields\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v10, v11 := range in.CustomFields {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"uploads\":"
		out.RawString(prefix)
		if in.Uploads == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Uploads {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"journals\":"
		out.RawString(prefix)
		if in.Journals == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Journals {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.SpentHours != 0 {
		const prefix string = ",\"spent_hours\":"
		out.RawString(prefix)
		out.Float32(float32(in.SpentHours))
	}
	if in.TotalSpentHours != 0 {
		const prefix string = ",\"total_spent_hours\":"
		out.RawString(prefix)
		out.Float32(float32(in.TotalSpentHours))
	}
	if in.EstimatedHours != 0 {
		const prefix string = ",\"estimated_hours\":"
		out.RawString(prefix)
		out.Float32(float32(in.EstimatedHours))
	}
	if in.DoneRatio != 0 {
		const prefix string = ",\"done_ratio\":"
		out.RawString(prefix)
		out.Int(int(in.DoneRatio))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Issue) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Issue) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Issue) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Issue) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine7(l, v)
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine8(in *jlexer.Lexer, out *Id) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine8(out *jwriter.Writer, in Id) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	out.RawByte('}')
}
func easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine9(in *jlexer.Lexer, out *CustomField) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "multiple":
			out.Multiple = bool(in.Bool())
		case "value":
			if m, ok := out.Value.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Value.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Value = in.Interface()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine9(out *jwriter.Writer, in CustomField) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"multiple\":"
		out.RawString(prefix)
		out.Bool(bool(in.Multiple))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if m, ok := in.Value.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Value.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Value))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomField) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomField) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9b96fefEncodeGithubComSomniSomGoRedmine9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomField) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomField) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9b96fefDecodeGithubComSomniSomGoRedmine9(l, v)
}