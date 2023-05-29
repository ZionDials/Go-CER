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

package template

import (
	"github.com/spf13/viper"
	"github.com/ziondials/go-cer/config"
	"github.com/ziondials/go-cer/models"
)

const AddressTemplate = `{{ .ERLName }}.`

func LocationStructBuilder(erl *models.ERL, conf *config.WebsiteConfig) *models.Response {

	var SoftKeys []models.SoftKey

	viper.UnmarshalKey("website.softkeys", &SoftKeys)

	if erl == nil {
		return &models.Response{
			LocationX: 0,
			LocationY: 0,
			Width:     800,
			Height:    480,
			Title:     conf.Title,
			URL:       conf.FailureURL,
			Prompt:    "Not Provisioned for CER",
			SoftKey:   SoftKeys,
		}
	}

	location := &models.Response{
		LocationX: 0,
		LocationY: 0,
		Width:     800,
		Height:    480,
		Title:     conf.Title,
		Prompt:    erl.ERLName,
		URL:       conf.SucessURL,
		SoftKey:   SoftKeys,
	}

	return location
}
