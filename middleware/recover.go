/*
 * @Author: qiuling
 * @Date: 2019-08-22 18:26:18
 * @Last Modified by: qiuling
 * @Last Modified time: 2020-05-21 23:00:02
 */
package middleware

import (
	"errors"
	"net/http"

	"github.com/wlxpkg/base/log"
	. "github.com/wlxpkg/tiktok"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// var buf [4096]byte
				// n := runtime.Stack(buf[:], false)
				// fmt.Printf("==> %s\n", string(buf[:n]))

				logData := make(map[string]string)
				bodyCopy := ""
				if m, ok := c.Get("bodyCopy"); ok && m != nil {
					bodyCopy = m.(string)
				}
				path := c.Request.URL.Path
				method := c.Request.Method

				logData["path"] = method + "@" + path
				logData["request"] = bodyCopy

				log.Info(logData)
				log.Err(r)
				var err error

				switch x := r.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = errors.New("内部系统错误")
				}

				errors, ok := Errs[err.Error()]
				if !ok {
					errors = Errors{Code: 0, Message: "内部系统错误"}
				}

				c.JSON(http.StatusOK, gin.H{
					"code":    errors.Code,
					"message": errors.Message,
					"data":    "",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
