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
	"fmt"
	"log"
)

type CT_TLAnimateEffectBehavior struct {
	// Transition
	TransitionAttr ST_TLAnimateEffectTransition
	// Filter
	FilterAttr *string
	// Property List
	PrLstAttr *string
	CBhvr     *CT_TLCommonBehaviorData
	// Progress
	Progress *CT_TLAnimVariant
}

func NewCT_TLAnimateEffectBehavior() *CT_TLAnimateEffectBehavior {
	ret := &CT_TLAnimateEffectBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLAnimateEffectBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TransitionAttr != ST_TLAnimateEffectTransitionUnset {
		attr, err := m.TransitionAttr.MarshalXMLAttr(xml.Name{Local: "transition"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filter"},
			Value: fmt.Sprintf("%v", *m.FilterAttr)})
	}
	if m.PrLstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "prLst"},
			Value: fmt.Sprintf("%v", *m.PrLstAttr)})
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.Progress != nil {
		seprogress := xml.StartElement{Name: xml.Name{Local: "p:progress"}}
		e.EncodeElement(m.Progress, seprogress)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimateEffectBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "transition" {
			m.TransitionAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "filter" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FilterAttr = &parsed
		}
		if attr.Name.Local == "prLst" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PrLstAttr = &parsed
		}
	}
lCT_TLAnimateEffectBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cBhvr":
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case "progress":
				m.Progress = NewCT_TLAnimVariant()
				if err := d.DecodeElement(m.Progress, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLAnimateEffectBehavior %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateEffectBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLAnimateEffectBehavior and its children
func (m *CT_TLAnimateEffectBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateEffectBehavior")
}

// ValidateWithPath validates the CT_TLAnimateEffectBehavior and its children, prefixing error messages with path
func (m *CT_TLAnimateEffectBehavior) ValidateWithPath(path string) error {
	if err := m.TransitionAttr.ValidateWithPath(path + "/TransitionAttr"); err != nil {
		return err
	}
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.Progress != nil {
		if err := m.Progress.ValidateWithPath(path + "/Progress"); err != nil {
			return err
		}
	}
	return nil
}
