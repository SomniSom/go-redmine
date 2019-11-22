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

func easyjson384e462aDecodeGithubComSomniSomGoRedmine(in *jlexer.Lexer, out *issueStatusesResult) {
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
		case "issue_statuses":
			if in.IsNull() {
				in.Skip()
				out.IssueStatuses = nil
			} else {
				in.Delim('[')
				if out.IssueStatuses == nil {
					if !in.IsDelim(']') {
						out.IssueStatuses = make([]IssueStatus, 0, 2)
					} else {
						out.IssueStatuses = []IssueStatus{}
					}
				} else {
					out.IssueStatuses = (out.IssueStatuses)[:0]
				}
				for !in.IsDelim(']') {
					var v1 IssueStatus
					(v1).UnmarshalEasyJSON(in)
					out.IssueStatuses = append(out.IssueStatuses, v1)
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
func easyjson384e462aEncodeGithubComSomniSomGoRedmine(out *jwriter.Writer, in issueStatusesResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"issue_statuses\":"
		out.RawString(prefix[1:])
		if in.IssueStatuses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.IssueStatuses {
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
func (v issueStatusesResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson384e462aEncodeGithubComSomniSomGoRedmine(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v issueStatusesResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson384e462aEncodeGithubComSomniSomGoRedmine(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *issueStatusesResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson384e462aDecodeGithubComSomniSomGoRedmine(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *issueStatusesResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson384e462aDecodeGithubComSomniSomGoRedmine(l, v)
}
func easyjson384e462aDecodeGithubComSomniSomGoRedmine1(in *jlexer.Lexer, out *IssueStatus) {
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
		case "is_default":
			out.IsDefault = bool(in.Bool())
		case "is_closed":
			out.IsClosed = bool(in.Bool())
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
func easyjson384e462aEncodeGithubComSomniSomGoRedmine1(out *jwriter.Writer, in IssueStatus) {
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
		const prefix string = ",\"is_default\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsDefault))
	}
	{
		const prefix string = ",\"is_closed\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsClosed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IssueStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson384e462aEncodeGithubComSomniSomGoRedmine1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IssueStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson384e462aEncodeGithubComSomniSomGoRedmine1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IssueStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson384e462aDecodeGithubComSomniSomGoRedmine1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IssueStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson384e462aDecodeGithubComSomniSomGoRedmine1(l, v)
}