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
    "encoding/xml"
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type GroupObjectDescriptorRealisationTypeB struct {
    UpdateEnable bool
    TransmitEnable bool
    SegmentSelectorEnable bool
    WriteEnable bool
    ReadEnable bool
    CommunicationEnable bool
    Priority CEMIPriority
    ValueType ComObjectValueType
    IGroupObjectDescriptorRealisationTypeB
}

// The corresponding interface
type IGroupObjectDescriptorRealisationTypeB interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewGroupObjectDescriptorRealisationTypeB(updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *GroupObjectDescriptorRealisationTypeB {
    return &GroupObjectDescriptorRealisationTypeB{UpdateEnable: updateEnable, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

func CastGroupObjectDescriptorRealisationTypeB(structType interface{}) *GroupObjectDescriptorRealisationTypeB {
    castFunc := func(typ interface{}) *GroupObjectDescriptorRealisationTypeB {
        if casted, ok := typ.(GroupObjectDescriptorRealisationTypeB); ok {
            return &casted
        }
        if casted, ok := typ.(*GroupObjectDescriptorRealisationTypeB); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *GroupObjectDescriptorRealisationTypeB) GetTypeName() string {
    return "GroupObjectDescriptorRealisationTypeB"
}

func (m *GroupObjectDescriptorRealisationTypeB) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (updateEnable)
    lengthInBits += 1

    // Simple field (transmitEnable)
    lengthInBits += 1

    // Simple field (segmentSelectorEnable)
    lengthInBits += 1

    // Simple field (writeEnable)
    lengthInBits += 1

    // Simple field (readEnable)
    lengthInBits += 1

    // Simple field (communicationEnable)
    lengthInBits += 1

    // Simple field (priority)
    lengthInBits += 2

    // Simple field (valueType)
    lengthInBits += 8

    return lengthInBits
}

func (m *GroupObjectDescriptorRealisationTypeB) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func GroupObjectDescriptorRealisationTypeBParse(io *utils.ReadBuffer) (*GroupObjectDescriptorRealisationTypeB, error) {

    // Simple Field (updateEnable)
    updateEnable, _updateEnableErr := io.ReadBit()
    if _updateEnableErr != nil {
        return nil, errors.New("Error parsing 'updateEnable' field " + _updateEnableErr.Error())
    }

    // Simple Field (transmitEnable)
    transmitEnable, _transmitEnableErr := io.ReadBit()
    if _transmitEnableErr != nil {
        return nil, errors.New("Error parsing 'transmitEnable' field " + _transmitEnableErr.Error())
    }

    // Simple Field (segmentSelectorEnable)
    segmentSelectorEnable, _segmentSelectorEnableErr := io.ReadBit()
    if _segmentSelectorEnableErr != nil {
        return nil, errors.New("Error parsing 'segmentSelectorEnable' field " + _segmentSelectorEnableErr.Error())
    }

    // Simple Field (writeEnable)
    writeEnable, _writeEnableErr := io.ReadBit()
    if _writeEnableErr != nil {
        return nil, errors.New("Error parsing 'writeEnable' field " + _writeEnableErr.Error())
    }

    // Simple Field (readEnable)
    readEnable, _readEnableErr := io.ReadBit()
    if _readEnableErr != nil {
        return nil, errors.New("Error parsing 'readEnable' field " + _readEnableErr.Error())
    }

    // Simple Field (communicationEnable)
    communicationEnable, _communicationEnableErr := io.ReadBit()
    if _communicationEnableErr != nil {
        return nil, errors.New("Error parsing 'communicationEnable' field " + _communicationEnableErr.Error())
    }

    // Simple Field (priority)
    priority, _priorityErr := CEMIPriorityParse(io)
    if _priorityErr != nil {
        return nil, errors.New("Error parsing 'priority' field " + _priorityErr.Error())
    }

    // Simple Field (valueType)
    valueType, _valueTypeErr := ComObjectValueTypeParse(io)
    if _valueTypeErr != nil {
        return nil, errors.New("Error parsing 'valueType' field " + _valueTypeErr.Error())
    }

    // Create the instance
    return NewGroupObjectDescriptorRealisationTypeB(updateEnable, transmitEnable, segmentSelectorEnable, writeEnable, readEnable, communicationEnable, priority, valueType), nil
}

func (m *GroupObjectDescriptorRealisationTypeB) Serialize(io utils.WriteBuffer) error {

    // Simple Field (updateEnable)
    updateEnable := bool(m.UpdateEnable)
    _updateEnableErr := io.WriteBit((updateEnable))
    if _updateEnableErr != nil {
        return errors.New("Error serializing 'updateEnable' field " + _updateEnableErr.Error())
    }

    // Simple Field (transmitEnable)
    transmitEnable := bool(m.TransmitEnable)
    _transmitEnableErr := io.WriteBit((transmitEnable))
    if _transmitEnableErr != nil {
        return errors.New("Error serializing 'transmitEnable' field " + _transmitEnableErr.Error())
    }

    // Simple Field (segmentSelectorEnable)
    segmentSelectorEnable := bool(m.SegmentSelectorEnable)
    _segmentSelectorEnableErr := io.WriteBit((segmentSelectorEnable))
    if _segmentSelectorEnableErr != nil {
        return errors.New("Error serializing 'segmentSelectorEnable' field " + _segmentSelectorEnableErr.Error())
    }

    // Simple Field (writeEnable)
    writeEnable := bool(m.WriteEnable)
    _writeEnableErr := io.WriteBit((writeEnable))
    if _writeEnableErr != nil {
        return errors.New("Error serializing 'writeEnable' field " + _writeEnableErr.Error())
    }

    // Simple Field (readEnable)
    readEnable := bool(m.ReadEnable)
    _readEnableErr := io.WriteBit((readEnable))
    if _readEnableErr != nil {
        return errors.New("Error serializing 'readEnable' field " + _readEnableErr.Error())
    }

    // Simple Field (communicationEnable)
    communicationEnable := bool(m.CommunicationEnable)
    _communicationEnableErr := io.WriteBit((communicationEnable))
    if _communicationEnableErr != nil {
        return errors.New("Error serializing 'communicationEnable' field " + _communicationEnableErr.Error())
    }

    // Simple Field (priority)
    _priorityErr := m.Priority.Serialize(io)
    if _priorityErr != nil {
        return errors.New("Error serializing 'priority' field " + _priorityErr.Error())
    }

    // Simple Field (valueType)
    _valueTypeErr := m.ValueType.Serialize(io)
    if _valueTypeErr != nil {
        return errors.New("Error serializing 'valueType' field " + _valueTypeErr.Error())
    }

    return nil
}

func (m *GroupObjectDescriptorRealisationTypeB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    for {
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "updateEnable":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.UpdateEnable = data
            case "transmitEnable":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.TransmitEnable = data
            case "segmentSelectorEnable":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SegmentSelectorEnable = data
            case "writeEnable":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.WriteEnable = data
            case "readEnable":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ReadEnable = data
            case "communicationEnable":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.CommunicationEnable = data
            case "priority":
                var data CEMIPriority
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Priority = data
            case "valueType":
                var data ComObjectValueType
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ValueType = data
            }
        }
    }
}

func (m *GroupObjectDescriptorRealisationTypeB) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.GroupObjectDescriptorRealisationTypeB"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.UpdateEnable, xml.StartElement{Name: xml.Name{Local: "updateEnable"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.TransmitEnable, xml.StartElement{Name: xml.Name{Local: "transmitEnable"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SegmentSelectorEnable, xml.StartElement{Name: xml.Name{Local: "segmentSelectorEnable"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.WriteEnable, xml.StartElement{Name: xml.Name{Local: "writeEnable"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ReadEnable, xml.StartElement{Name: xml.Name{Local: "readEnable"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.CommunicationEnable, xml.StartElement{Name: xml.Name{Local: "communicationEnable"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ValueType, xml.StartElement{Name: xml.Name{Local: "valueType"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
