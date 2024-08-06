package middleware

import (
	"backend/api"
	"backend/internal/service"
	"net/http"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/genv"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// Auth validates the request to allow only signed-in users visit.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	ok, _ := r.Session.Contains("uid")
	if ok {
		r.Middleware.Next()
	} else {
		r.Response.WriteJson(g.Map{
			"code":    403,
			"message": "Login timeout",
			"data":    nil,
		})
	}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// Set allow all cors when dev or staging (only in docker)
	env := genv.Get("APP_ENV", gvar.New("prod"))
	if env.String() == "prod" {
		corsOptions.AllowDomain = []string{"unclechair.vip", "silicious.net"}
		r.Response.CORS(corsOptions)
	} else {
		r.Response.CORSDefault()
	}
	r.Middleware.Next()
}

// i18n language setting.
func (s *sMiddleware) Language(r *ghttp.Request) {
	lang := r.GetQuery("lang", "zh-CN").String()
	ctx := gi18n.WithLanguage(r.Context(), lang)
	r.SetCtx(ctx)
	r.Middleware.Next()
}

// Prevent context canceled error.
func (s *sMiddleware) NeverDoneCtx(r *ghttp.Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}

// ResponseHandler is the middleware handling response object and its error.
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status >= 300 || r.Response.Status < 200 {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}
	r.Response.WriteJson(api.CommonRes{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
