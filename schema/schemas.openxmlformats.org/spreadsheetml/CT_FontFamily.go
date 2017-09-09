// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_FontFamily struct {
	ValAttr int64
}

func NewCT_FontFamily() *CT_FontFamily {
	ret := &CT_FontFamily{}
	ret.ValAttr = 0
	return ret
}

func (m *CT_FontFamily) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FontFamily) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = 0
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FontFamily: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FontFamily and its children
func (m *CT_FontFamily) Validate() error {
	return m.ValidateWithPath("CT_FontFamily")
}

// ValidateWithPath validates the CT_FontFamily and its children, prefixing error messages with path
func (m *CT_FontFamily) ValidateWithPath(path string) error {
	if m.ValAttr < 0 {
		return fmt.Errorf("%s/m.ValAttr must be >= 0 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr > 14 {
		return fmt.Errorf("%s/m.ValAttr must be <= 14 (have %v)", path, m.ValAttr)
	}
	return nil
}
