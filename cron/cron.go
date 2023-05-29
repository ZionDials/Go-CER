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

package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/ziondials/go-cer/cer"
	"github.com/ziondials/go-cer/config"
	"github.com/ziondials/go-cer/db"
)

func RunCronJobs(ds *db.DataService) {
	cerConfig := config.GetCERFromGlobalConfig()
	cer := cer.NewCERService(ds, cerConfig)

	s := gocron.NewScheduler(time.UTC)

	s.Every(cerConfig.Refresh).Minutes().Do(func() {
		cer.GetERLs()
		cer.GetPhones()
	})

	s.StartAsync()
}
