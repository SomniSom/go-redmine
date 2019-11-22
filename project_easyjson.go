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

func easyjson7b166cadDecodeGithubComSomniSomGoRedmine(in *jlexer.Lexer, out *projectsResult) {
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
		case "projects":
			if in.IsNull() {
				in.Skip()
				out.Projects = nil
			} else {
				in.Delim('[')
				if out.Projects == nil {
					if !in.IsDelim(']') {
						out.Projects = make([]Project, 0, 1)
					} else {
						out.Projects = []Project{}
					}
				} else {
					out.Projects = (out.Projects)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Project
					(v1).UnmarshalEasyJSON(in)
					out.Projects = append(out.Projects, v1)
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
func easyjson7b166cadEncodeGithubComSomniSomGoRedmine(out *jwriter.Writer, in projectsResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"projects\":"
		out.RawString(prefix[1:])
		if in.Projects == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Projects {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v projectsResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v projectsResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *projectsResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *projectsResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine(l, v)
}
func easyjson7b166cadDecodeGithubComSomniSomGoRedmine1(in *jlexer.Lexer, out *projectResult) {
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
		case "project":
			(out.Project).UnmarshalEasyJSON(in)
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
func easyjson7b166cadEncodeGithubComSomniSomGoRedmine1(out *jwriter.Writer, in projectResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"project\":"
		out.RawString(prefix[1:])
		(in.Project).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v projectResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v projectResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *projectResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *projectResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine1(l, v)
}
func easyjson7b166cadDecodeGithubComSomniSomGoRedmine2(in *jlexer.Lexer, out *projectRequest) {
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
		case "project":
			(out.Project).UnmarshalEasyJSON(in)
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
func easyjson7b166cadEncodeGithubComSomniSomGoRedmine2(out *jwriter.Writer, in projectRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"project\":"
		out.RawString(prefix[1:])
		(in.Project).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v projectRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v projectRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *projectRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *projectRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine2(l, v)
}
func easyjson7b166cadDecodeGithubComSomniSomGoRedmine3(in *jlexer.Lexer, out *Project) {
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
		case "identifier":
			out.Identifier = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "created_on":
			out.CreatedOn = string(in.String())
		case "updated_on":
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
func easyjson7b166cadEncodeGithubComSomniSomGoRedmine3(out *jwriter.Writer, in Project) {
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
		const prefix string = ",\"identifier\":"
		out.RawString(prefix)
		out.String(string(in.Identifier))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Project) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Project) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b166cadEncodeGithubComSomniSomGoRedmine3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Project) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Project) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b166cadDecodeGithubComSomniSomGoRedmine3(l, v)
}