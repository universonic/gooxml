// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_VerticalAlignRun struct {
	// Subscript/Superscript Value
	ValAttr sharedTypes.ST_VerticalAlignRun
}

func NewCT_VerticalAlignRun() *CT_VerticalAlignRun {
	ret := &CT_VerticalAlignRun{}
	ret.ValAttr = sharedTypes.ST_VerticalAlignRun(1)
	return ret
}

func (m *CT_VerticalAlignRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VerticalAlignRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = sharedTypes.ST_VerticalAlignRun(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_VerticalAlignRun: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_VerticalAlignRun and its children
func (m *CT_VerticalAlignRun) Validate() error {
	return m.ValidateWithPath("CT_VerticalAlignRun")
}

// ValidateWithPath validates the CT_VerticalAlignRun and its children, prefixing error messages with path
func (m *CT_VerticalAlignRun) ValidateWithPath(path string) error {
	if m.ValAttr == sharedTypes.ST_VerticalAlignRunUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
