//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type SysexCommandPinStateResponse struct {
	Pin      uint8
	PinMode  uint8
	PinState uint8
	Parent   *SysexCommand
}

// The corresponding interface
type ISysexCommandPinStateResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandPinStateResponse) CommandType() uint8 {
	return 0x6E
}

func (m *SysexCommandPinStateResponse) Response() bool {
	return false
}

func (m *SysexCommandPinStateResponse) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandPinStateResponse(pin uint8, pinMode uint8, pinState uint8) *SysexCommand {
	child := &SysexCommandPinStateResponse{
		Pin:      pin,
		PinMode:  pinMode,
		PinState: pinState,
		Parent:   NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandPinStateResponse(structType interface{}) *SysexCommandPinStateResponse {
	castFunc := func(typ interface{}) *SysexCommandPinStateResponse {
		if casted, ok := typ.(SysexCommandPinStateResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandPinStateResponse); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandPinStateResponse(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandPinStateResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandPinStateResponse) GetTypeName() string {
	return "SysexCommandPinStateResponse"
}

func (m *SysexCommandPinStateResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandPinStateResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Simple field (pinMode)
	lengthInBits += 8

	// Simple field (pinState)
	lengthInBits += 8

	return lengthInBits
}

func (m *SysexCommandPinStateResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandPinStateResponseParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandPinStateResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pin)
	pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}

	// Simple Field (pinMode)
	pinMode, _pinModeErr := readBuffer.ReadUint8("pinMode", 8)
	if _pinModeErr != nil {
		return nil, errors.Wrap(_pinModeErr, "Error parsing 'pinMode' field")
	}

	// Simple Field (pinState)
	pinState, _pinStateErr := readBuffer.ReadUint8("pinState", 8)
	if _pinStateErr != nil {
		return nil, errors.Wrap(_pinStateErr, "Error parsing 'pinState' field")
	}

	if closeErr := readBuffer.CloseContext("SysexCommandPinStateResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandPinStateResponse{
		Pin:      pin,
		PinMode:  pinMode,
		PinState: pinState,
		Parent:   &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandPinStateResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandPinStateResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Simple Field (pinMode)
		pinMode := uint8(m.PinMode)
		_pinModeErr := writeBuffer.WriteUint8("pinMode", 8, (pinMode))
		if _pinModeErr != nil {
			return errors.Wrap(_pinModeErr, "Error serializing 'pinMode' field")
		}

		// Simple Field (pinState)
		pinState := uint8(m.PinState)
		_pinStateErr := writeBuffer.WriteUint8("pinState", 8, (pinState))
		if _pinStateErr != nil {
			return errors.Wrap(_pinStateErr, "Error serializing 'pinState' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandPinStateResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandPinStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}