/*
Simple Go Get Alias
Copyright (C) 2021 Daniel A. Hawton (daniel@hawton.org)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package server

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/dhawton/hawton.io/internal/middleware"
	"github.com/dhawton/log4g"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	engine *gin.Engine
}

var log = log4g.Category("server")
var server *Server

func Run(port int) {
	intro := figure.NewFigure("hawton.dev", "", false).Slicify()
	for i := 0; i < len(intro); i++ {
		log.Info(intro[i])
	}

	log.Info("Checking for .env, loading if exists")
	if _, err := os.Stat(".env"); err == nil {
		log.Info("Found .env, loading")
		err := godotenv.Load()
		if err != nil {
			log.Error("Error loading .env file")
		}
	}

	log.Info("Configuring server")
	server = NewServer()
	SetupRoutes()

	log.Info("Starting server on :%d", port)
	server.engine.Run(fmt.Sprintf(":%d", port))
}

func NewServer() *Server {
	log.Debug("Setting gin mode to release mode")
	gin.SetMode(gin.ReleaseMode)

	server := Server{}
	engine := gin.New()

	log.Debug("Loading Recovery middleware")
	engine.Use(gin.Recovery())

	log.Debug("Loading CORS middleware")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	engine.Use(cors.New(corsConfig))

	log.Debug("Loading Logger middleware")
	engine.Use(middleware.Logger)

	log.Debug("Loading HTML globs")
	server.engine = engine
	engine.LoadHTMLGlob("static/*")

	return &server
}
