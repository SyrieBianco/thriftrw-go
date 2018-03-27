// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package api

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Plugin_Handshake_Args represents the arguments for the Plugin.handshake function.
//
// The arguments for handshake are sent and received over the wire as this struct.
type Plugin_Handshake_Args struct {
	Request *HandshakeRequest `json:"request,omitempty"`
}

// ToWire translates a Plugin_Handshake_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Plugin_Handshake_Args) ToWire() (wire.Value, error) {
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

func _HandshakeRequest_Read(w wire.Value) (*HandshakeRequest, error) {
	var v HandshakeRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Plugin_Handshake_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Plugin_Handshake_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Plugin_Handshake_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Plugin_Handshake_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _HandshakeRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Plugin_Handshake_Args
// struct.
func (v *Plugin_Handshake_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Plugin_Handshake_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Plugin_Handshake_Args match the
// provided Plugin_Handshake_Args.
//
// This function performs a deep comparison.
func (v *Plugin_Handshake_Args) Equals(rhs *Plugin_Handshake_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "handshake" for this struct.
func (v *Plugin_Handshake_Args) MethodName() string {
	return "handshake"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Plugin_Handshake_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Plugin_Handshake_Helper provides functions that aid in handling the
// parameters and return values of the Plugin.handshake
// function.
var Plugin_Handshake_Helper = struct {
	// Args accepts the parameters of handshake in-order and returns
	// the arguments struct for the function.
	Args func(
		request *HandshakeRequest,
	) *Plugin_Handshake_Args

	// IsException returns true if the given error can be thrown
	// by handshake.
	//
	// An error can be thrown by handshake only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for handshake
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// handshake into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by handshake
	//
	//   value, err := handshake(args)
	//   result, err := Plugin_Handshake_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from handshake: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*HandshakeResponse, error) (*Plugin_Handshake_Result, error)

	// UnwrapResponse takes the result struct for handshake
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if handshake threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Plugin_Handshake_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Plugin_Handshake_Result) (*HandshakeResponse, error)
}{}

func init() {
	Plugin_Handshake_Helper.Args = func(
		request *HandshakeRequest,
	) *Plugin_Handshake_Args {
		return &Plugin_Handshake_Args{
			Request: request,
		}
	}

	Plugin_Handshake_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Plugin_Handshake_Helper.WrapResponse = func(success *HandshakeResponse, err error) (*Plugin_Handshake_Result, error) {
		if err == nil {
			return &Plugin_Handshake_Result{Success: success}, nil
		}

		return nil, err
	}
	Plugin_Handshake_Helper.UnwrapResponse = func(result *Plugin_Handshake_Result) (success *HandshakeResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Plugin_Handshake_Result represents the result of a Plugin.handshake function call.
//
// The result of a handshake execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Plugin_Handshake_Result struct {
	// Value returned by handshake after a successful execution.
	Success *HandshakeResponse `json:"success,omitempty"`
}

// ToWire translates a Plugin_Handshake_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Plugin_Handshake_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Plugin_Handshake_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _HandshakeResponse_Read(w wire.Value) (*HandshakeResponse, error) {
	var v HandshakeResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Plugin_Handshake_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Plugin_Handshake_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Plugin_Handshake_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Plugin_Handshake_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _HandshakeResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Plugin_Handshake_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Plugin_Handshake_Result
// struct.
func (v *Plugin_Handshake_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Plugin_Handshake_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Plugin_Handshake_Result match the
// provided Plugin_Handshake_Result.
//
// This function performs a deep comparison.
func (v *Plugin_Handshake_Result) Equals(rhs *Plugin_Handshake_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "handshake" for this struct.
func (v *Plugin_Handshake_Result) MethodName() string {
	return "handshake"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Plugin_Handshake_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
