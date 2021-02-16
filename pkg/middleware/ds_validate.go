package middleware

import (
    "strings"
    "net/http"

    macaron "gopkg.in/macaron.v1"

    "github.com/grafana/grafana/pkg/models"
)

func ValidateCanAccessDatasourceProxyApi() macaron.Handler {
    return func(c *models.ReqContext, req *http.Request, ctx *macaron.Context) {
       dashId := c.Req.Header.Get("X-Grafana-Org-Id")

       if strings.HasPrefix(req.RequestURI, "/api/datasources/proxy") {

          if !c.IsGrafanaAdmin && !c.SignedInUser.HasRole(models.ROLE_EDITOR) && dashId == "" {
             accessForbidden(c)
             return
          }
       }
    return
    }
}
