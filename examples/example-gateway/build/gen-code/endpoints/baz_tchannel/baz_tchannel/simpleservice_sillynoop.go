// Code generated by thriftrw v1.3.0
// @generated

package baz_tchannel

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SimpleService_SillyNoop_Args struct{}

func (v *SimpleService_SillyNoop_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SimpleService_SillyNoop_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SimpleService_SillyNoop_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("SimpleService_SillyNoop_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SimpleService_SillyNoop_Args) Equals(rhs *SimpleService_SillyNoop_Args) bool {
	return true
}

func (v *SimpleService_SillyNoop_Args) MethodName() string {
	return "SillyNoop"
}

func (v *SimpleService_SillyNoop_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SimpleService_SillyNoop_Helper = struct {
	Args           func() *SimpleService_SillyNoop_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*SimpleService_SillyNoop_Result, error)
	UnwrapResponse func(*SimpleService_SillyNoop_Result) error
}{}

func init() {
	SimpleService_SillyNoop_Helper.Args = func() *SimpleService_SillyNoop_Args {
		return &SimpleService_SillyNoop_Args{}
	}
	SimpleService_SillyNoop_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		case *ServerErr:
			return true
		default:
			return false
		}
	}
	SimpleService_SillyNoop_Helper.WrapResponse = func(err error) (*SimpleService_SillyNoop_Result, error) {
		if err == nil {
			return &SimpleService_SillyNoop_Result{}, nil
		}
		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_SillyNoop_Result.AuthErr")
			}
			return &SimpleService_SillyNoop_Result{AuthErr: e}, nil
		case *ServerErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_SillyNoop_Result.ServerErr")
			}
			return &SimpleService_SillyNoop_Result{ServerErr: e}, nil
		}
		return nil, err
	}
	SimpleService_SillyNoop_Helper.UnwrapResponse = func(result *SimpleService_SillyNoop_Result) (err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
			return
		}
		if result.ServerErr != nil {
			err = result.ServerErr
			return
		}
		return
	}
}

type SimpleService_SillyNoop_Result struct {
	AuthErr   *AuthErr   `json:"authErr,omitempty"`
	ServerErr *ServerErr `json:"serverErr,omitempty"`
}

func (v *SimpleService_SillyNoop_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.AuthErr != nil {
		w, err = v.AuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServerErr != nil {
		w, err = v.ServerErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_SillyNoop_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ServerErr_Read(w wire.Value) (*ServerErr, error) {
	var v ServerErr
	err := v.FromWire(w)
	return &v, err
}

func (v *SimpleService_SillyNoop_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AuthErr, err = _AuthErr_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServerErr, err = _ServerErr_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.AuthErr != nil {
		count++
	}
	if v.ServerErr != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("SimpleService_SillyNoop_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *SimpleService_SillyNoop_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}
	if v.ServerErr != nil {
		fields[i] = fmt.Sprintf("ServerErr: %v", v.ServerErr)
		i++
	}
	return fmt.Sprintf("SimpleService_SillyNoop_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SimpleService_SillyNoop_Result) Equals(rhs *SimpleService_SillyNoop_Result) bool {
	if !((v.AuthErr == nil && rhs.AuthErr == nil) || (v.AuthErr != nil && rhs.AuthErr != nil && v.AuthErr.Equals(rhs.AuthErr))) {
		return false
	}
	if !((v.ServerErr == nil && rhs.ServerErr == nil) || (v.ServerErr != nil && rhs.ServerErr != nil && v.ServerErr.Equals(rhs.ServerErr))) {
		return false
	}
	return true
}

func (v *SimpleService_SillyNoop_Result) MethodName() string {
	return "SillyNoop"
}

func (v *SimpleService_SillyNoop_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
