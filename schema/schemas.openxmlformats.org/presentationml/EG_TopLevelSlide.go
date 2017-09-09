// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type EG_TopLevelSlide struct {
	// Color Scheme Map
	ClrMap *drawingml.CT_ColorMapping
}

func NewEG_TopLevelSlide() *EG_TopLevelSlide {
	ret := &EG_TopLevelSlide{}
	ret.ClrMap = drawingml.NewCT_ColorMapping()
	return ret
}

func (m *EG_TopLevelSlide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	seclrMap := xml.StartElement{Name: xml.Name{Local: "p:clrMap"}}
	e.EncodeElement(m.ClrMap, seclrMap)
	return nil
}

func (m *EG_TopLevelSlide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ClrMap = drawingml.NewCT_ColorMapping()
lEG_TopLevelSlide:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "clrMap":
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_TopLevelSlide %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TopLevelSlide
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TopLevelSlide and its children
func (m *EG_TopLevelSlide) Validate() error {
	return m.ValidateWithPath("EG_TopLevelSlide")
}

// ValidateWithPath validates the EG_TopLevelSlide and its children, prefixing error messages with path
func (m *EG_TopLevelSlide) ValidateWithPath(path string) error {
	if err := m.ClrMap.ValidateWithPath(path + "/ClrMap"); err != nil {
		return err
	}
	return nil
}
