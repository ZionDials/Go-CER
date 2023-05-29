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

	"gorm.io/gorm"
)

type ERL struct {
	ERLName   string         `xml:"ERLName" gorm:"not null;primaryKey"`
	ALI       ALI            `xml:"ALIInfo" gorm:"embedded"`
	CreatedAt time.Time      `xml:"-" json:"-"`
	UpdatedAt time.Time      `xml:"-" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" xml:"-" json:"-"`
}

type ALI struct {
	HouseNumber       string `xml:"HouseNumber"`
	HouseNumberSuffix string `xml:"HouseNumberSuffix"`
	StreetName        string `xml:"StreetName"`
	PrefixDirectional string `xml:"PrefixDirectional"`
	StreetSuffix      string `xml:"StreetSuffix"`
	PostDirectional   string `xml:"PostDirectional"`
	CommunityName     string `xml:"CommunityName"`
	State             string `xml:"State"`
	MainNPA           string `xml:"MainNPA"`
	CustomerName      string `xml:"CustomerName"`
	COS               string `xml:"COS"`
	TOS               string `xml:"TOS"`
	Exchange          string `xml:"Exchange"`
	MainTelNum        string `xml:"MainTelNum"`
	OrderNumber       string `xml:"OrderNumber"`
	CountyID          string `xml:"CountyID"`
	CompanyID         string `xml:"CompanyID"`
	ZipCode           string `xml:"ZipCode"`
	ZipCodeExt        string `xml:"ZipCodeExt"`
	CustomerCode      string `xml:"CustomerCode"`
	Comments          string `xml:"Comments"`
	Longitude         string `xml:"Longitude"`
	Latitude          string `xml:"Latitude"`
	Elevation         string `xml:"Elevation"`
	TARCode           string `xml:"TARCode"`
	Location          string `xml:"Location"`
	ProvReserved      string `xml:"ProvReserved"`
	ERLType           string `xml:"ERLType"`
}

type ERLResponse struct {
	XMLName    xml.Name `xml:"ConventionalERL"`
	Status     string   `xml:"status"`
	ERLDetails []*ERL   `xml:"ERLDetails"`
}
