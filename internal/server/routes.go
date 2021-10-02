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
	"github.com/gin-gonic/gin"
	"github.com/vzau/common/utils"
)

func SetupRoutes() {
	server.engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "https://github.com/dhawton")
	})

	server.engine.GET("/:pkg", func(c *gin.Context) {
		if val, exists := c.GetQuery("go-get"); exists && val == "1" {
			c.HTML(200, "go-get.tmpl", gin.H{
				"alias_base": utils.Getenv("ALIAS_BASE", "hawton.dev"),
				"git_base":   utils.Getenv("BASE_GO_GET_URL", "https://github.com/dhawton"),
				"pkg":        c.Param("pkg"),
			})
		} else {
			if c.Param("pkg") == "ping" {
				c.String(200, "PONG")
			} else {
				c.Redirect(302, "https://github.com/dhawton")
			}
		}
	})
}
