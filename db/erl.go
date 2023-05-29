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

package db

import (
	"github.com/ziondials/go-cer/models"
	"gorm.io/gorm/clause"
)

func (ds DataService) CreateERLs(erls []*models.ERL) error {

	if rsp := ds.Session.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&erls, 50); rsp.Error != nil {
		return rsp.Error
	}

	return nil
}

func (ds DataService) GetERLByName(name string) (*models.ERL, error) {
	var erl models.ERL

	queryParams := models.ERL{ERLName: name}
	if rsp := ds.Session.First(&erl, queryParams); rsp.Error != nil {
		return nil, rsp.Error
	}

	return &erl, nil
}
