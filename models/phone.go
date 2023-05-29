// Copyright (c) 2023 Zion Dials <me@ziondials.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"encoding/xml"
	"time"

	"github.com/ziondials/go-cer/tools"
	"gorm.io/gorm"
)

type Phone struct {
	MAC              string         `xml:"PhoneMACAddress" gorm:"not null;primaryKey"`
	SwitchHostName   string         `xml:"SwitchHostName"`
	SwitchIPAddress  string         `xml:"SwitchIPAddress"`
	ERLName          string         `xml:"ERLName"`
	ERL              ERL            `gorm:"foreignKey:ERLName;references:ERLName"`
	IfName           string         `xml:"IfName"`
	Location         string         `xml:"Location"`
	PhoneExtension   string         `xml:"PhoneExtension"`
	PhoneIPAddress   string         `xml:"PhoneIPAddress"`
	PhoneIPv6Address string         `xml:"PhoneIPv6Address"`
	PhoneType        string         `xml:"PhoneType"`
	PortDescription  string         `xml:"PortDescription"`
	PortIdentifier   string         `xml:"PortIdentifier"`
	CreatedAt        time.Time      `xml:"-" json:"-"`
	UpdatedAt        time.Time      `xml:"-" json:"-"`
	DeletedAt        gorm.DeletedAt `xml:"-" json:"-"`
}

type Switch struct {
	SwitchHostName  string   `xml:"SwitchHostName"`
	SwitchIPAddress string   `xml:"SwitchIPAddress"`
	Phone           []*Phone `xml:"portDetails"`
}

type PhoneResponse struct {
	XMLName xml.Name  `xml:"SwitchPort"`
	Status  string    `xml:"status"`
	Switch  []*Switch `xml:"switchip"`
}

func (pr *PhoneResponse) Parse(b []byte) ([]*Phone, error) {
	if err := xml.Unmarshal(b, &pr); err != nil {
		return nil, err
	}

	var phones []*Phone
	for _, sw := range pr.Switch {
		for _, ph := range sw.Phone {
			if ph.ERLName == "" || ph.MAC == "" {
				continue
			} else {
				ph.SwitchHostName = sw.SwitchHostName
				ph.SwitchIPAddress = sw.SwitchIPAddress
				ph.PhoneIPAddress = tools.SpaceMap(ph.PhoneIPAddress)
				phones = append(phones, ph)
			}
		}
	}

	return phones, nil
}
