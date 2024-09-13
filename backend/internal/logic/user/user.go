package user

import (
	auth "backend/api/auth/v1"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
	"encoding/hex"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) GenerateUid(ctx context.Context) (uid string) {
	uid = grand.Digits(10)
	check, _ := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Uid, uid).Count()
	for check > 0 {
		uid = grand.Digits(10)
		check, _ = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Uid, uid).Count()
	}
	return
}

func (s *sUser) CreateUser(ctx context.Context, data *auth.SignupReq) (user *entity.Users, err error) {
	r := g.RequestFromCtx(ctx)
	if err := s.checkUserEmail(r.Context(), data.Email); err != nil {
		return nil, err
	}
	if err := s.checkUsername(r.Context(), data.Username); err != nil {
		return nil, err
	}
	if err := gconv.Struct(data, &user); err != nil {
		return nil, err
	}
	uid := s.GenerateUid(ctx)
	user.Uid = uid
	encrypted, err := gaes.Encrypt([]byte(data.Password), []byte(consts.AESKey))
	if err != nil {
		return nil, err
	}
	user.Password = hex.EncodeToString(encrypted)
	user.Nickname = user.Username
	user.AvatarUrl = "uri"
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err = gerror.NewCodef(gcode.CodeDbOperationError, "User create failed")
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = dao.Users.Ctx(ctx).TX(tx).Data(user).OmitEmpty().Save()
	return
}

// This function is used to find user from ctx after passing authenticator middleware
func (s *sUser) GetUserFromCtx(ctx context.Context) *gdb.Model {
	key := dao.Users.Columns().Uid
	uid := g.RequestFromCtx(ctx).GetParam(key)
	return dao.Users.Ctx(ctx).WithAll().Where(key, uid)
}

func (s *sUser) GetUserFromUid(ctx context.Context, uid string) (user *entity.Users, err error) {
	key := dao.Users.Columns().Uid
	err = dao.Users.Ctx(ctx).Where(key, uid).Scan(&user)
	return
}

func (s *sUser) checkUsername(ctx context.Context, username string) error {
	n, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, username).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`username" %s "has been used`, username)
	}
	return nil
}

func (s *sUser) checkUserEmail(ctx context.Context, email string) error {
	n, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Email, email).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`Email has been used`)
	}
	return nil
}
