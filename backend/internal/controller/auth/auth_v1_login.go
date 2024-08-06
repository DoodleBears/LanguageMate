package auth

import (
	"context"
	"encoding/hex"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "backend/api/auth/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	r := g.RequestFromCtx(ctx)
	var user *gdb.Model
	identity := req.Username
	if identity == "" {
		identity = req.Email
	}
	user = dao.Users.Ctx(ctx).Where(g.Map{
		dao.Users.Columns().Username: identity,
	})
	exist, _ := user.Count()
	if exist == 0 {
		return nil, AuthError.NoUser
	}
	result, err := user.Fields(
		dao.Users.Columns().Password,
		dao.Users.Columns().Uid,
		dao.Users.Columns().Lock,
		dao.Users.Columns().LoginAttempts).One()
	if err != nil {
		return nil, AuthError.InternalError
	}

	pass := result[dao.Users.Columns().Password]
	uid := result[dao.Users.Columns().Uid]
	lock := result[dao.Users.Columns().Lock]
	loginAttempts := result[dao.Users.Columns().LoginAttempts]

	if lock.Bool() {
		return nil, AuthError.TooManyAttempts
	}
	password, _ := hex.DecodeString(pass.String())
	decrypt, err := gaes.Decrypt(password, []byte(consts.AESKey))
	if err != nil {
		return nil, err
	}
	if gbinary.DecodeToString(decrypt) != req.Password {
		if loginAttempts.Int() > 3 {
			user.Data(g.Map{
				dao.Users.Columns().LastLogin: gtime.Now(),
				dao.Users.Columns().Lock:      true,
			}).Update()
		} else {
			user.Data(g.Map{
				dao.Users.Columns().LastLogin:     gtime.Now(),
				dao.Users.Columns().LoginAttempts: loginAttempts.Int() + 1,
			}).Update()
		}
		return nil, AuthError.WrongCredential
	} else {
		user.Data(g.Map{
			dao.Users.Columns().LastLogin:     gtime.Now(),
			dao.Users.Columns().LoginAttempts: 0,
		}).Update()

		SessionId := service.Session().GenerateSessionId(r)
		r.Session.Set("uid", uid)
		r.Cookie.SetSessionId(SessionId)
		return
	}
}

var AuthError = struct {
	InternalError   error
	NoUser          error
	WrongCredential error
	TooManyAttempts error
	Banned          error
}{
	InternalError:   gerror.New("Something wrong with server"),
	NoUser:          gerror.New("No related users"),
	WrongCredential: gerror.New("Wrong username or password"),
	Banned:          gerror.New("Your account has been banned, please contact admin"),
}
