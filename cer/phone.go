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
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/ziondials/go-cer/logger"
	"github.com/ziondials/go-cer/models"
	"github.com/ziondials/go-cer/version"
)

func (cer *CERService) GetPhones() {

	logger.Info("Getting Phones from CER...")
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{Timeout: 30 * time.Second, Jar: jar}

	PhoneRequestURL := "https://" + cer.Hostname + "/cerappservices/export/switchport/info"

	req, err := http.NewRequest(http.MethodGet, PhoneRequestURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", version.FullAppName)
	req.Header.Set("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.SetBasicAuth(cer.Username, cer.Password)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusUnauthorized {
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var p models.PhoneResponse

		phones, err := p.Parse(resBody)
		if err != nil {
			log.Fatal(err)
		}
		logger.Info("Phones received from CER: %d", len(phones))
		cer.DataService.CreatePhones(phones)
	}
}
