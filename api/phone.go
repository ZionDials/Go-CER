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

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziondials/go-cer/logger"
	"github.com/ziondials/go-cer/template"
)

func (a *APIController) serveRoot(c *gin.Context) {

	requestIP := c.ClientIP()
	c.Header("Content-Type", "text/xml")

	phoneLocation, err := a.DataService.GetPhoneByIPAddress(requestIP)
	nilLocation := template.LocationStructBuilder(nil, a.Config.Website)

	if err != nil || phoneLocation == nil {
		logger.Info("No location found for IP: %s", requestIP)
		c.XML(http.StatusOK, nilLocation)
		return
	}

	erl, err := a.DataService.GetERLByName(phoneLocation.ERLName)
	if err != nil || erl == nil {
		logger.Info("No ERL found for ERLName: %s", phoneLocation.ERLName)
		c.XML(http.StatusOK, nilLocation)
		return
	}

	location := template.LocationStructBuilder(erl, a.Config.Website)
	logger.Info("Serving location for IP: %s, ERLName: %s", requestIP, phoneLocation.ERLName)

	c.XML(http.StatusOK, location)
}
