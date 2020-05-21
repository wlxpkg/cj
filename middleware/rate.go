/*
 * 限流中间件
 * @Author: qiuling
 * @Date: 2019-09-18 16:16:34
 * @Last Modified by: qiuling
 * @Last Modified time: 2020-05-21 23:00:02
 */

package middleware

import (
	"github.com/wlxpkg/base/model"
	. "github.com/wlxpkg/tiktok"
	"strings"

	"github.com/gin-gonic/gin"
)

func Rate() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("authorization")
		jwt := strings.TrimPrefix(authorization, "Bearer ")
		method := c.Request.Method
		path := c.Request.URL.Path
		// R(path, "path")
		if jwt != "" {
			check := model.RateCheck(jwt, method, path)
			if !check {
				err := Excp("ERR_TOO_MANY_REQUEST")
				Abort(c, err)
				return
			}
		}
	}
}
