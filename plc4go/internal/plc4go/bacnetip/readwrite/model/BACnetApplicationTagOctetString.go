/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetApplicationTagOctetString struct {
	*BACnetTag
	Value             string
	ActualLengthInBit uint16
}

// The corresponding interface
type IBACnetApplicationTagOctetString interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagOctetString) TagClass() TagClass {
	return TagClass_APPLICATION_TAGS
}

func (m *BACnetApplicationTagOctetString) InitializeParent(parent *BACnetTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isBoolean bool, isConstructed bool, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.TagNumber = tagNumber
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetApplicationTagOctetString(value string, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetTag {
	child := &BACnetApplicationTagOctetString{
		Value:     value,
		BACnetTag: NewBACnetTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Child = child
	return child.BACnetTag
}

func CastBACnetApplicationTagOctetString(structType interface{}) *BACnetApplicationTagOctetString {
	castFunc := func(typ interface{}) *BACnetApplicationTagOctetString {
		if casted, ok := typ.(BACnetApplicationTagOctetString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagOctetString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetApplicationTagOctetString(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetApplicationTagOctetString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagOctetString) GetTypeName() string {
	return "BACnetApplicationTagOctetString"
}

func (m *BACnetApplicationTagOctetString) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagOctetString) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.ActualLengthInBit)

	return lengthInBits
}

func (m *BACnetApplicationTagOctetString) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagOctetStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagOctetString"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_actualLengthInBit := uint16(actualLength) * uint16(uint16(8))
	actualLengthInBit := uint16(_actualLengthInBit)

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadString("value", uint32(actualLengthInBit))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagOctetString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagOctetString{
		Value:             value,
		ActualLengthInBit: actualLengthInBit,
		BACnetTag:         &BACnetTag{},
	}
	_child.BACnetTag.Child = _child
	return _child.BACnetTag, nil
}

func (m *BACnetApplicationTagOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagOctetString"); pushErr != nil {
			return pushErr
		}
		// Virtual field
		if _actualLengthInBitErr := writeBuffer.WriteVirtual("actualLengthInBit", m.ActualLengthInBit); _actualLengthInBitErr != nil {
			return errors.Wrap(_actualLengthInBitErr, "Error serializing 'actualLengthInBit' field")
		}

		// Simple Field (value)
		value := string(m.Value)
		_valueErr := writeBuffer.WriteString("value", uint32(m.ActualLengthInBit), "ASCII", (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagOctetString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
