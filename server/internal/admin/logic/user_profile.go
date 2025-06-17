package logic

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	dao2 "gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	"gmanager/internal/library/gftoken"
)

func (s *user) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	model, err := s.Get(ctx, gftoken.GetUserId(ctx))
	if err != nil {
		return
	}
	if model == nil {
		return nil, nil
	}

	deptMap, err := Dept.DeptMap(ctx)
	if err != nil {
		return nil, err
	}

	val, err := dao2.Role.Ctx(ctx).Fields("GROUP_CONCAT(name)").WhereIn(dao2.Role.Columns().Id, model.RoleIds).Value()
	if err != nil {
		return nil, err
	}

	res = &v1.UserProfileRes{
		Id:        model.Id,
		UserName:  model.UserName,
		NickName:  model.NickName,
		Mobile:    model.Mobile,
		Avatar:    model.Avatar,
		Email:     model.Email,
		Gender:    model.Gender,
		CreateAt:  model.CreateAt,
		DeptName:  deptMap[model.DeptId].Name,
		RoleNames: val.String(),
	}
	return
}

func (s *user) SaveProfile(ctx context.Context, in *v1.UserSaveProfileReq) error {
	userId := gftoken.GetUserId(ctx)
	var model do.User
	err := gconv.Struct(in, &model)
	if err != nil {
		return gerror.Wrap(err, "数据转换错误")
	}

	model.Id = userId
	model.UpdateId = userId
	model.UpdateAt = gtime.Now()

	_, err = dao2.User.Ctx(ctx).OmitEmpty().Where(dao2.User.Columns().Id, model.Id).Update(model)
	if err != nil {
		return err
	}
	_ = Log.Save(ctx, model, consts.UPDATE)
	return err
}

// ChangePassword 修改密码
func (s *user) ChangePassword(ctx context.Context, in *v1.UserChangePasswordReq) error {
	userId := gftoken.GetUserId(ctx)
	res, err := s.Get(ctx, userId)
	if err != nil {
		return err
	}
	if res == nil {
		return nil
	}
	oldPassword, err := gmd5.Encrypt(in.OldPassword + res.Salt)
	if err != nil {
		return err
	}
	if oldPassword != res.Password {
		return gerror.New("原密码错误")
	}

	password, err := gmd5.Encrypt(in.NewPassword + res.Salt)
	if err != nil {
		return err
	}
	columns := dao2.User.Columns()
	_, err = dao2.User.Ctx(ctx).Where(columns.Id, userId).Update(do.User{
		UpdateId: userId,
		UpdateAt: gtime.Now(),
		Password: password,
	})
	return nil
}

func (s *user) SendMobileCode(ctx context.Context, req *v1.UserSendMobileCodeReq) error {
	return nil
}

func (s *user) SaveMobile(ctx context.Context, req *v1.UserSaveMobileReq) error {
	// 验证码校验
	userId := gftoken.GetUserId(ctx)
	return s.SaveProfile(ctx, &v1.UserSaveProfileReq{Id: userId, Mobile: req.Mobile})
}

func (s *user) SendEmailCode(ctx context.Context, req *v1.UserSendEmailCodeReq) error {
	return nil
}

func (s *user) SaveEmail(ctx context.Context, req *v1.UserSaveEmailReq) error {
	// 验证码校验
	userId := gftoken.GetUserId(ctx)
	return s.SaveProfile(ctx, &v1.UserSaveProfileReq{Id: userId, Email: req.Email})
}
