package session

import (
	"backend/internal/service"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

func (s *sSession) GenerateSessionId(r *ghttp.Request) string {
	var (
		address = r.RemoteAddr
		header  = fmt.Sprintf("%v", r.Header)
	)
	return guid.S([]byte(address), []byte(header))
}
