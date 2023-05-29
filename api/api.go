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
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/ziondials/go-cer/config"
	"github.com/ziondials/go-cer/cron"
	"github.com/ziondials/go-cer/db"
	"github.com/ziondials/go-cer/logger"
)

type APIController struct {
	DataService *db.DataService
	Config      *config.GlobalConfig
}

func InitAPI() *APIController {
	conf := config.GetGlobalConfig()
	dataService := db.InitDB(conf.Database)

	cron.RunCronJobs(dataService)

	return &APIController{
		DataService: dataService,
		Config:      conf,
	}
}

func Serve() {

	config.SetDefaults()
	logger.InitLogger()

	apiController := InitAPI()

	conf := config.GetGlobalConfig()
	if conf.Logging.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(requestid.New())

	router.GET("/root.xml", apiController.serveRoot)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(conf.Website.Port),
		Handler: router,
	}

	go func() {
		// service connections
		logger.Info("Listening on http://localhost:" + strconv.Itoa(conf.Website.Port) + "/root.xml")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panic(fmt.Sprintf("listen: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.Info("timeout of 1 seconds.")
	}
	logger.Info("Server exiting")
}
