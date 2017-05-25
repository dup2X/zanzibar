// Code generated by zanzibar
// @generated
// Checksum : Ykdd2glAqGrExuLmVSjafw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	foo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/foo"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(in *jlexer.Lexer, out *Bar_TooManyArgs_Result) {
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
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, &*out.Success)
			}
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(in, &*out.BarException)
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
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(out *jwriter.Writer, in Bar_TooManyArgs_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.Success)
		}
	}
	if in.BarException != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"barException\":")
		if in.BarException == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(out, *in.BarException)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_TooManyArgs_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_TooManyArgs_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_TooManyArgs_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_TooManyArgs_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs(l, v)
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(in *jlexer.Lexer, out *BarException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar1(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[string]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 int32
					v1 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 int32
					v2 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithRange\":")
	out.Int32(int32(in.IntWithRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithoutRange\":")
	out.Int32(int32(in.IntWithoutRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithRange\":")
	if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in.MapIntWithRange {
			if !v3First {
				out.RawByte(',')
			}
			v3First = false
			out.String(string(v3Name))
			out.RawByte(':')
			out.Int32(int32(v3Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithoutRange\":")
	if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.MapIntWithoutRange {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			out.Int32(int32(v4Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(in *jlexer.Lexer, out *Bar_TooManyArgs_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
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
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(BarRequest)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(in, &*out.Request)
			}
			RequestSet = true
		case "foo":
			if in.IsNull() {
				in.Skip()
				out.Foo = nil
			} else {
				if out.Foo == nil {
					out.Foo = new(foo.FooStruct)
				}
				easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(in, &*out.Foo)
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
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(out *jwriter.Writer, in Bar_TooManyArgs_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	if in.Request == nil {
		out.RawString("null")
	} else {
		easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(out, *in.Request)
	}
	if in.Foo != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"foo\":")
		if in.Foo == nil {
			out.RawString("null")
		} else {
			easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(out, *in.Foo)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_TooManyArgs_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_TooManyArgs_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_TooManyArgs_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_TooManyArgs_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarTooManyArgs1(l, v)
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(in *jlexer.Lexer, out *foo.FooStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var FooStringSet bool
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
		case "fooString":
			out.FooString = string(in.String())
			FooStringSet = true
		case "fooI32":
			if in.IsNull() {
				in.Skip()
				out.FooI32 = nil
			} else {
				if out.FooI32 == nil {
					out.FooI32 = new(int32)
				}
				*out.FooI32 = int32(in.Int32())
			}
		case "fooI16":
			if in.IsNull() {
				in.Skip()
				out.FooI16 = nil
			} else {
				if out.FooI16 == nil {
					out.FooI16 = new(int16)
				}
				*out.FooI16 = int16(in.Int16())
			}
		case "fooDouble":
			if in.IsNull() {
				in.Skip()
				out.FooDouble = nil
			} else {
				if out.FooDouble == nil {
					out.FooDouble = new(float64)
				}
				*out.FooDouble = float64(in.Float64())
			}
		case "fooBool":
			if in.IsNull() {
				in.Skip()
				out.FooBool = nil
			} else {
				if out.FooBool == nil {
					out.FooBool = new(bool)
				}
				*out.FooBool = bool(in.Bool())
			}
		case "fooMap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.FooMap = make(map[string]string)
				} else {
					out.FooMap = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 string
					v5 = string(in.String())
					(out.FooMap)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
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
	if !FooStringSet {
		in.AddError(fmt.Errorf("key 'fooString' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(out *jwriter.Writer, in foo.FooStruct) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fooString\":")
	out.String(string(in.FooString))
	if in.FooI32 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooI32\":")
		if in.FooI32 == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.FooI32))
		}
	}
	if in.FooI16 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooI16\":")
		if in.FooI16 == nil {
			out.RawString("null")
		} else {
			out.Int16(int16(*in.FooI16))
		}
	}
	if in.FooDouble != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooDouble\":")
		if in.FooDouble == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.FooDouble))
		}
	}
	if in.FooBool != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooBool\":")
		if in.FooBool == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.FooBool))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fooMap\":")
	if in.FooMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v6First := true
		for v6Name, v6Value := range in.FooMap {
			if !v6First {
				out.RawByte(',')
			}
			v6First = false
			out.String(string(v6Name))
			out.RawByte(':')
			out.String(string(v6Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}
func easyjson87e68f88DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(in *jlexer.Lexer, out *BarRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var BoolFieldSet bool
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
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "boolField":
			out.BoolField = bool(in.Bool())
			BoolFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !BoolFieldSet {
		in.AddError(fmt.Errorf("key 'boolField' is required"))
	}
}
func easyjson87e68f88EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar2(out *jwriter.Writer, in BarRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"boolField\":")
	out.Bool(bool(in.BoolField))
	out.RawByte('}')
}
