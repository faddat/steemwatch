package home

import (
	"net/http"

	"github.com/tchap/steemwatch/server/context"
	"github.com/tchap/steemwatch/server/views"

	"github.com/labstack/echo"
)

type Handler struct {
	ctx *context.Context
}

func NewHandlerFunc(serverCtx *context.Context) echo.HandlerFunc {
	handler := &Handler{serverCtx}
	return handler.HandlerFunc
}

func (handler *Handler) HandlerFunc(ctx echo.Context) error {
	profile, err := handler.ctx.SessionManager.GetProfile(ctx)
	if err != nil {
		return err
	}

	var (
		templateName string
		templateCtx  = &views.PageContext{CanonicalURL: handler.ctx.CanonicalURL}
	)
	if profile == nil {
		templateName = "welcome.html"
	} else {
		templateName = "app.html"
		templateCtx.Environment = string(handler.ctx.Env)
		templateCtx.UserId = profile.Id
		templateCtx.UserEmail = profile.Email
		templateCtx.UserDisplayName = profile.Email

		for k, v := range profile.SocialLinks {
			templateCtx.UserDisplayName = v.UserName + "@" + k
			break
		}
	}

	return ctx.Render(http.StatusOK, templateName, templateCtx)
}
