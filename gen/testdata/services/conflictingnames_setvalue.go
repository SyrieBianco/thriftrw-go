// Code generated by thriftrw v1.1.0
// @generated

package services

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ConflictingNames_SetValue_Args struct {
	Request *ConflictingNamesSetValueArgs `json:"request,omitempty"`
}

func (v *ConflictingNames_SetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ConflictingNamesSetValueArgs_Read(w wire.Value) (*ConflictingNamesSetValueArgs, error) {
	var v ConflictingNamesSetValueArgs
	err := v.FromWire(w)
	return &v, err
}

func (v *ConflictingNames_SetValue_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _ConflictingNamesSetValueArgs_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ConflictingNames_SetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("ConflictingNames_SetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ConflictingNames_SetValue_Args) Equals(rhs *ConflictingNames_SetValue_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}
	return true
}

func (v *ConflictingNames_SetValue_Args) MethodName() string {
	return "setValue"
}

func (v *ConflictingNames_SetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ConflictingNames_SetValue_Helper = struct {
	Args           func(request *ConflictingNamesSetValueArgs) *ConflictingNames_SetValue_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*ConflictingNames_SetValue_Result, error)
	UnwrapResponse func(*ConflictingNames_SetValue_Result) error
}{}

func init() {
	ConflictingNames_SetValue_Helper.Args = func(request *ConflictingNamesSetValueArgs) *ConflictingNames_SetValue_Args {
		return &ConflictingNames_SetValue_Args{Request: request}
	}
	ConflictingNames_SetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	ConflictingNames_SetValue_Helper.WrapResponse = func(err error) (*ConflictingNames_SetValue_Result, error) {
		if err == nil {
			return &ConflictingNames_SetValue_Result{}, nil
		}
		return nil, err
	}
	ConflictingNames_SetValue_Helper.UnwrapResponse = func(result *ConflictingNames_SetValue_Result) (err error) {
		return
	}
}

type ConflictingNames_SetValue_Result struct{}

func (v *ConflictingNames_SetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ConflictingNames_SetValue_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *ConflictingNames_SetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("ConflictingNames_SetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ConflictingNames_SetValue_Result) Equals(rhs *ConflictingNames_SetValue_Result) bool {
	return true
}

func (v *ConflictingNames_SetValue_Result) MethodName() string {
	return "setValue"
}

func (v *ConflictingNames_SetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
