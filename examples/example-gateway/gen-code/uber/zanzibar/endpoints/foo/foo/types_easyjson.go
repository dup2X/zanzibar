// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package foo

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

func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo(in *jlexer.Lexer, out *FooStruct) {
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
		case "fooString":
			out.FooString = string(in.String())
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
					out.FooMap = make(map[string]*FooName)
				} else {
					out.FooMap = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 *FooName
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(FooName)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					(out.FooMap)[key] = v1
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
}
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo(out *jwriter.Writer, in FooStruct) {
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
		v2First := true
		for v2Name, v2Value := range in.FooMap {
			if !v2First {
				out.RawByte(',')
			}
			v2First = false
			out.String(string(v2Name))
			out.RawByte(':')
			if v2Value == nil {
				out.RawString("null")
			} else {
				(*v2Value).MarshalEasyJSON(out)
			}
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FooStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FooStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FooStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FooStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo1(in *jlexer.Lexer, out *FooName) {
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo1(out *jwriter.Writer, in FooName) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FooName) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FooName) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FooName) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FooName) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayGenCodeUberZanzibarEndpointsFooFoo1(l, v)
}
