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
type BACnetPropertyValueRelinquishDefault struct {
	*BACnetPropertyValue
	Value *BACnetApplicationTagReal
}

// The corresponding interface
type IBACnetPropertyValueRelinquishDefault interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetPropertyValueRelinquishDefault) Identifier() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

func (m *BACnetPropertyValueRelinquishDefault) InitializeParent(parent *BACnetPropertyValue) {
}

func NewBACnetPropertyValueRelinquishDefault(value *BACnetApplicationTagReal) *BACnetPropertyValue {
	child := &BACnetPropertyValueRelinquishDefault{
		Value:               value,
		BACnetPropertyValue: NewBACnetPropertyValue(),
	}
	child.Child = child
	return child.BACnetPropertyValue
}

func CastBACnetPropertyValueRelinquishDefault(structType interface{}) *BACnetPropertyValueRelinquishDefault {
	castFunc := func(typ interface{}) *BACnetPropertyValueRelinquishDefault {
		if casted, ok := typ.(BACnetPropertyValueRelinquishDefault); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetPropertyValueRelinquishDefault); ok {
			return casted
		}
		if casted, ok := typ.(BACnetPropertyValue); ok {
			return CastBACnetPropertyValueRelinquishDefault(casted.Child)
		}
		if casted, ok := typ.(*BACnetPropertyValue); ok {
			return CastBACnetPropertyValueRelinquishDefault(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetPropertyValueRelinquishDefault) GetTypeName() string {
	return "BACnetPropertyValueRelinquishDefault"
}

func (m *BACnetPropertyValueRelinquishDefault) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetPropertyValueRelinquishDefault) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += m.Value.LengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyValueRelinquishDefault) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetPropertyValueRelinquishDefaultParse(readBuffer utils.ReadBuffer, identifier BACnetPropertyIdentifier, actualLength uint32) (*BACnetPropertyValue, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyValueRelinquishDefault"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
	_value, _valueErr := BACnetTagParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := CastBACnetApplicationTagReal(_value)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyValueRelinquishDefault"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyValueRelinquishDefault{
		Value:               CastBACnetApplicationTagReal(value),
		BACnetPropertyValue: &BACnetPropertyValue{},
	}
	_child.BACnetPropertyValue.Child = _child
	return _child.BACnetPropertyValue, nil
}

func (m *BACnetPropertyValueRelinquishDefault) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyValueRelinquishDefault"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return pushErr
		}
		_valueErr := m.Value.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return popErr
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyValueRelinquishDefault"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
