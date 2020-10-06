package middleware

import (
    "strings"
	"context"
	"net/url"

	macaron "gopkg.in/macaron.v1"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/setting"
)

func ValidateCanAccessDatasourceProxyApi() macaron.Handler {
    return func(c *models.ReqContext, ctx *macaron.Context) {
	   dashId := c.Req.Header.Get("x-dashboard-id")

	   if strings.HasPrefix(req.RequestURI, "/api/datasources/proxy/:id/api/v1") {
			
	      if  dashId == "" {
		     accessForbidden(c)
		     return
	      }
	   }
    return
    }
}
