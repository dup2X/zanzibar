// Code generated by zanzibar
// @generated
// Checksum : RBLto9narUj0U8KnSd4Gjw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package googlenow

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

func easyjson50713e80DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials(in *jlexer.Lexer, out *GoogleNow_CheckCredentials_Result) {
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
func easyjson50713e80EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials(out *jwriter.Writer, in GoogleNow_CheckCredentials_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNow_CheckCredentials_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson50713e80EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNow_CheckCredentials_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson50713e80EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNow_CheckCredentials_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson50713e80DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNow_CheckCredentials_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson50713e80DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials(l, v)
}
func easyjson50713e80DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials1(in *jlexer.Lexer, out *GoogleNow_CheckCredentials_Args) {
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
func easyjson50713e80EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials1(out *jwriter.Writer, in GoogleNow_CheckCredentials_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GoogleNow_CheckCredentials_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson50713e80EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GoogleNow_CheckCredentials_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson50713e80EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GoogleNow_CheckCredentials_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson50713e80DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GoogleNow_CheckCredentials_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson50713e80DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsGooglenowGooglenowGoogleNowCheckCredentials1(l, v)
}
