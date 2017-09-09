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

type CT_PresentationProperties struct {
	// HTML Publishing Properties
	HtmlPubPr *CT_HtmlPublishProperties
	// Web Properties
	WebPr *CT_WebProperties
	// Printing Properties
	PrnPr *CT_PrintProperties
	// Presentation-wide Show Properties
	ShowPr *CT_ShowProperties
	// Color MRU
	ClrMru *drawingml.CT_ColorMRU
	ExtLst *CT_ExtensionList
}

func NewCT_PresentationProperties() *CT_PresentationProperties {
	ret := &CT_PresentationProperties{}
	return ret
}

func (m *CT_PresentationProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.HtmlPubPr != nil {
		sehtmlPubPr := xml.StartElement{Name: xml.Name{Local: "p:htmlPubPr"}}
		e.EncodeElement(m.HtmlPubPr, sehtmlPubPr)
	}
	if m.WebPr != nil {
		sewebPr := xml.StartElement{Name: xml.Name{Local: "p:webPr"}}
		e.EncodeElement(m.WebPr, sewebPr)
	}
	if m.PrnPr != nil {
		seprnPr := xml.StartElement{Name: xml.Name{Local: "p:prnPr"}}
		e.EncodeElement(m.PrnPr, seprnPr)
	}
	if m.ShowPr != nil {
		seshowPr := xml.StartElement{Name: xml.Name{Local: "p:showPr"}}
		e.EncodeElement(m.ShowPr, seshowPr)
	}
	if m.ClrMru != nil {
		seclrMru := xml.StartElement{Name: xml.Name{Local: "p:clrMru"}}
		e.EncodeElement(m.ClrMru, seclrMru)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PresentationProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PresentationProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "htmlPubPr":
				m.HtmlPubPr = NewCT_HtmlPublishProperties()
				if err := d.DecodeElement(m.HtmlPubPr, &el); err != nil {
					return err
				}
			case "webPr":
				m.WebPr = NewCT_WebProperties()
				if err := d.DecodeElement(m.WebPr, &el); err != nil {
					return err
				}
			case "prnPr":
				m.PrnPr = NewCT_PrintProperties()
				if err := d.DecodeElement(m.PrnPr, &el); err != nil {
					return err
				}
			case "showPr":
				m.ShowPr = NewCT_ShowProperties()
				if err := d.DecodeElement(m.ShowPr, &el); err != nil {
					return err
				}
			case "clrMru":
				m.ClrMru = drawingml.NewCT_ColorMRU()
				if err := d.DecodeElement(m.ClrMru, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PresentationProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PresentationProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PresentationProperties and its children
func (m *CT_PresentationProperties) Validate() error {
	return m.ValidateWithPath("CT_PresentationProperties")
}

// ValidateWithPath validates the CT_PresentationProperties and its children, prefixing error messages with path
func (m *CT_PresentationProperties) ValidateWithPath(path string) error {
	if m.HtmlPubPr != nil {
		if err := m.HtmlPubPr.ValidateWithPath(path + "/HtmlPubPr"); err != nil {
			return err
		}
	}
	if m.WebPr != nil {
		if err := m.WebPr.ValidateWithPath(path + "/WebPr"); err != nil {
			return err
		}
	}
	if m.PrnPr != nil {
		if err := m.PrnPr.ValidateWithPath(path + "/PrnPr"); err != nil {
			return err
		}
	}
	if m.ShowPr != nil {
		if err := m.ShowPr.ValidateWithPath(path + "/ShowPr"); err != nil {
			return err
		}
	}
	if m.ClrMru != nil {
		if err := m.ClrMru.ValidateWithPath(path + "/ClrMru"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
