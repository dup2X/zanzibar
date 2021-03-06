// Code generated by thriftrw v1.3.0
// @generated

package googlenow

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type GoogleNow_AddCredentials_Args struct {
	AuthCode string `json:"authCode,required"`
}

func (v *GoogleNow_AddCredentials_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.AuthCode), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GoogleNow_AddCredentials_Args) FromWire(w wire.Value) error {
	var err error
	authCodeIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.AuthCode, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				authCodeIsSet = true
			}
		}
	}
	if !authCodeIsSet {
		return errors.New("field AuthCode of GoogleNow_AddCredentials_Args is required")
	}
	return nil
}

func (v *GoogleNow_AddCredentials_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("AuthCode: %v", v.AuthCode)
	i++
	return fmt.Sprintf("GoogleNow_AddCredentials_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *GoogleNow_AddCredentials_Args) Equals(rhs *GoogleNow_AddCredentials_Args) bool {
	if !(v.AuthCode == rhs.AuthCode) {
		return false
	}
	return true
}

func (v *GoogleNow_AddCredentials_Args) MethodName() string {
	return "addCredentials"
}

func (v *GoogleNow_AddCredentials_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var GoogleNow_AddCredentials_Helper = struct {
	Args           func(authCode string) *GoogleNow_AddCredentials_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*GoogleNow_AddCredentials_Result, error)
	UnwrapResponse func(*GoogleNow_AddCredentials_Result) error
}{}

func init() {
	GoogleNow_AddCredentials_Helper.Args = func(authCode string) *GoogleNow_AddCredentials_Args {
		return &GoogleNow_AddCredentials_Args{AuthCode: authCode}
	}
	GoogleNow_AddCredentials_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	GoogleNow_AddCredentials_Helper.WrapResponse = func(err error) (*GoogleNow_AddCredentials_Result, error) {
		if err == nil {
			return &GoogleNow_AddCredentials_Result{}, nil
		}
		return nil, err
	}
	GoogleNow_AddCredentials_Helper.UnwrapResponse = func(result *GoogleNow_AddCredentials_Result) (err error) {
		return
	}
}

type GoogleNow_AddCredentials_Result struct{}

func (v *GoogleNow_AddCredentials_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GoogleNow_AddCredentials_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *GoogleNow_AddCredentials_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("GoogleNow_AddCredentials_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *GoogleNow_AddCredentials_Result) Equals(rhs *GoogleNow_AddCredentials_Result) bool {
	return true
}

func (v *GoogleNow_AddCredentials_Result) MethodName() string {
	return "addCredentials"
}

func (v *GoogleNow_AddCredentials_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
