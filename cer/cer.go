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

package cer

import (
	"github.com/ziondials/go-cer/config"
	"github.com/ziondials/go-cer/db"
)

type CERService struct {
	DataService *db.DataService
	Hostname    string
	Username    string
	Password    string
}

func NewCERService(dataService *db.DataService, config *config.CERConfig) *CERService {

	return &CERService{DataService: dataService, Hostname: config.Host, Username: config.Username, Password: config.Password}
}
