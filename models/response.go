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

import "encoding/xml"

type Response struct {
	XMLName   xml.Name  `xml:"CiscoIPPhoneImageFile"`
	LocationX int       `xml:"LocationX"`
	LocationY int       `xml:"LocationY"`
	Width     int       `xml:"Width"`
	Height    int       `xml:"Height"`
	Title     string    `xml:"Title"`
	Prompt    string    `xml:"Prompt"`
	URL       string    `xml:"URL"`
	SoftKey   []SoftKey `xml:"SoftKeyItem"`
}

type SoftKey struct {
	Name     string `xml:"Name" mapstructure:"Name"`
	URL      string `xml:"URL" mapstructure:"URL"`
	Position int    `xml:"Position" mapstructure:"Position"`
}
