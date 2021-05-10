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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type DF1UnprotectedReadResponse struct {
	Data   []uint8
	Parent *DF1Command
}

// The corresponding interface
type IDF1UnprotectedReadResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DF1UnprotectedReadResponse) CommandCode() uint8 {
	return 0x41
}

func (m *DF1UnprotectedReadResponse) InitializeParent(parent *DF1Command, status uint8, transactionCounter uint16) {
	m.Parent.Status = status
	m.Parent.TransactionCounter = transactionCounter
}

func NewDF1UnprotectedReadResponse(data []uint8, status uint8, transactionCounter uint16) *DF1Command {
	child := &DF1UnprotectedReadResponse{
		Data:   data,
		Parent: NewDF1Command(status, transactionCounter),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastDF1UnprotectedReadResponse(structType interface{}) *DF1UnprotectedReadResponse {
	castFunc := func(typ interface{}) *DF1UnprotectedReadResponse {
		if casted, ok := typ.(DF1UnprotectedReadResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1UnprotectedReadResponse); ok {
			return casted
		}
		if casted, ok := typ.(DF1Command); ok {
			return CastDF1UnprotectedReadResponse(casted.Child)
		}
		if casted, ok := typ.(*DF1Command); ok {
			return CastDF1UnprotectedReadResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1UnprotectedReadResponse) GetTypeName() string {
	return "DF1UnprotectedReadResponse"
}

func (m *DF1UnprotectedReadResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DF1UnprotectedReadResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Manual Array Field (data)
	data := m.Data
	lengthInBits += DF1UtilsDataLength(data) * 8

	return lengthInBits
}

func (m *DF1UnprotectedReadResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DF1UnprotectedReadResponseParse(readBuffer utils.ReadBuffer) (*DF1Command, error) {
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadResponse"); pullErr != nil {
		return nil, pullErr
	}
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Manual Array Field (data)
	// Terminated array
	_dataList := make([]uint8, 0)
	for !((bool)(DF1UtilsDataTerminate(readBuffer))) {
		_dataList = append(_dataList, ((uint8)(DF1UtilsReadData(readBuffer))))

	}
	data := make([]uint8, len(_dataList))
	for i := 0; i < len(_dataList); i++ {
		data[i] = uint8(_dataList[i])
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1UnprotectedReadResponse{
		Data:   data,
		Parent: &DF1Command{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *DF1UnprotectedReadResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadResponse"); pushErr != nil {
			return pushErr
		}

		// Manual Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, Element := range m.Data {
				DF1UtilsWriteData(writeBuffer, Element)
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *DF1UnprotectedReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}