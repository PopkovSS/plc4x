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
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetConfirmedServiceACKReadProperty struct {
    BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKReadProperty interface {
    IBACnetConfirmedServiceACK
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceACKReadProperty) ServiceChoice() uint8 {
    return 0x0C
}

func (m BACnetConfirmedServiceACKReadProperty) initialize() spi.Message {
    return m
}

func NewBACnetConfirmedServiceACKReadProperty() BACnetConfirmedServiceACKInitializer {
    return &BACnetConfirmedServiceACKReadProperty{}
}

func CastIBACnetConfirmedServiceACKReadProperty(structType interface{}) IBACnetConfirmedServiceACKReadProperty {
    castFunc := func(typ interface{}) IBACnetConfirmedServiceACKReadProperty {
        if iBACnetConfirmedServiceACKReadProperty, ok := typ.(IBACnetConfirmedServiceACKReadProperty); ok {
            return iBACnetConfirmedServiceACKReadProperty
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetConfirmedServiceACKReadProperty(structType interface{}) BACnetConfirmedServiceACKReadProperty {
    castFunc := func(typ interface{}) BACnetConfirmedServiceACKReadProperty {
        if sBACnetConfirmedServiceACKReadProperty, ok := typ.(BACnetConfirmedServiceACKReadProperty); ok {
            return sBACnetConfirmedServiceACKReadProperty
        }
        if sBACnetConfirmedServiceACKReadProperty, ok := typ.(*BACnetConfirmedServiceACKReadProperty); ok {
            return *sBACnetConfirmedServiceACKReadProperty
        }
        return BACnetConfirmedServiceACKReadProperty{}
    }
    return castFunc(structType)
}

func (m BACnetConfirmedServiceACKReadProperty) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetConfirmedServiceACK.LengthInBits()

    return lengthInBits
}

func (m BACnetConfirmedServiceACKReadProperty) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKReadPropertyParse(io *utils.ReadBuffer) (BACnetConfirmedServiceACKInitializer, error) {

    // Create the instance
    return NewBACnetConfirmedServiceACKReadProperty(), nil
}

func (m BACnetConfirmedServiceACKReadProperty) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetConfirmedServiceACKSerialize(io, m.BACnetConfirmedServiceACK, CastIBACnetConfirmedServiceACK(m), ser)
}

func (m *BACnetConfirmedServiceACKReadProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
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
            }
        }
    }
}

func (m BACnetConfirmedServiceACKReadProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKReadProperty"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
