// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package core

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

func easyjson9280440fDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(in *jlexer.Lexer, out *Admin) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "mail":
			out.Mail = string(in.String())
		case "password":
			out.PasswordHash = string(in.String())
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
func easyjson9280440fEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(out *jwriter.Writer, in Admin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"mail\":"
		out.RawString(prefix)
		out.String(string(in.Mail))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.PasswordHash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Admin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9280440fEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Admin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9280440fEncodeGithubComGoParkMailRu20232UmlautInternalModelCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Admin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9280440fDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Admin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9280440fDecodeGithubComGoParkMailRu20232UmlautInternalModelCore(l, v)
}
