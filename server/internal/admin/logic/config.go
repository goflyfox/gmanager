package logic

import (
	"context"
	"errors"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/consts"
	"gmanager/internal/dao"
	"gmanager/internal/library/cache"
	"gmanager/internal/library/gftoken"
	"gmanager/internal/model/do"
	"gmanager/internal/model/entity"
	"gmanager/internal/model/input"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Config 配置服务
var Config = new(config)

type config struct{}

// List 获取配置列表
func (s *config) List(ctx context.Context, in *v1.ConfigListReq) (res *v1.ConfigListRes, err error) {
	if in == nil {
		return
	}
	m := dao.Config.Ctx(ctx)
	columns := dao.Config.Columns()
	res = &v1.ConfigListRes{}

	// where条件
	if in.Keywords != "" {
		m = m.Where(m.Builder().WhereLike(columns.Name, "%"+in.Keywords+"%").
			WhereOrLike(columns.Key, "%"+in.Keywords+"%"))
	}
	if in.DataType > 0 {
		m = m.Where(columns.DataType, in.DataType)
	}
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}

	res.Total, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据数量失败！")
		return
	}
	res.CurrentPage = in.PageNum
	if res.Total == 0 {
		return
	}

	if in.OrderBy != "" {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("sort desc,id desc")
	}
	var pageList []*entity.Config
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	res.List = pageList

	return
}

func (s *config) DictOptions(ctx context.Context, in *v1.ConfigDictOptionsReq) (res *v1.ConfigDictOptionsRes, err error) {
	cacheMap, err := s.GetCache(ctx)
	if err != nil {
		return
	}
	options := make([]*input.OptionVal, 0, len(cacheMap))
	for _, v := range cacheMap {
		var model *entity.Config
		err = gconv.Struct(v, &model)
		if err != nil {
			return
		}
		if model.DataType != consts.DataTypeDict {
			continue
		}
		t := &input.OptionVal{
			Value: model.Id,
			Label: model.Name,
		}
		options = append(options, t)
	}
	res = &options
	return
}

func (s *config) Value(ctx context.Context, in *v1.ConfigValueReq) (*v1.ConfigValueRes, error) {
	if in.Name == "" && in.Key == "" {
		return nil, gerror.New("参数错误")
	}
	var model *entity.Config
	m := dao.Config.Ctx(ctx).Where(dao.Config.Columns().Enable, consts.EnableYes)
	if in.Key != "" {
		m = m.Where(dao.Config.Columns().Key, in.Key)
	}
	if in.Name != "" {
		m = m.Where(dao.Config.Columns().Name, in.Name)
	}
	err := m.Scan(&model)
	if err != nil {
		return nil, err
	}
	return &v1.ConfigValueRes{
		Id:    model.Id,
		Key:   model.Key,
		Name:  model.Name,
		Value: model.Value,
		Code:  model.Code,
	}, nil
}

func (s *config) DictItems(ctx context.Context, in *v1.ConfigDictItemsReq) (res *v1.ConfigDictItemsRes, err error) {
	cacheMap, err := s.GetCache(ctx)
	if err != nil {
		return
	}
	options := make([]*input.OptionVal, 0, len(cacheMap))
	for _, v := range cacheMap {
		var model *entity.Config
		err = gconv.Struct(v, &model)
		if err != nil {
			return
		}
		if model.DataType != consts.DataTypeDictData {
			continue
		}
		if model.ParentKey != in.ParentKey {
			continue
		}
		t := &input.OptionVal{
			Value: gconv.Int64(model.Value),
			Label: model.Name,
		}
		options = append(options, t)
	}
	res = &options
	return
}

// Get 获取配置详情
func (s *config) Get(ctx context.Context, id int64) (res *v1.ConfigGetRes, err error) {
	err = dao.Config.Ctx(ctx).Where(dao.Config.Columns().Id, id).Scan(&res)
	return
}

// Save 保存配置
func (s *config) Save(ctx context.Context, in *v1.ConfigSaveReq) error {
	if (in.DataType == consts.DataTypeKv || in.DataType == consts.DataTypeDictData) && in.Value == "" {
		return errors.New("配置值不能为空")
	}
	if in.DataType == consts.DataTypeDictData && in.ParentId == 0 {
		return errors.New("所属字典不能为空")
	}
	var model do.Config
	err := gconv.Struct(in, &model)
	if err != nil {
		return errors.New("数据转换错误")
	}

	// 配置名称唯一性校验
	nameCount, err := dao.Config.Ctx(ctx).Where(dao.Config.Columns().Name, model.Name).WhereNot(dao.Config.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "检查配置名称唯一性失败")
	}
	if nameCount > 0 {
		return gerror.New("配置名称已存在")
	}

	// 配置键唯一性校验
	keyCount, err := dao.Config.Ctx(ctx).Where(dao.Config.Columns().Key, model.Key).WhereNot(dao.Config.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "检查配置键唯一性失败")
	}
	if keyCount > 0 {
		return gerror.New("配置键已存在")
	}

	if in.DataType == consts.DataTypeDictData {
		res, err := s.Get(ctx, in.ParentId)
		if err != nil {
			return err
		}
		model.ParentKey = res.Key
	}

	m := dao.Config.Ctx(ctx)
	columns := dao.Config.Columns()

	configId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = configId
	model.UpdateAt = gtime.Now()
	if in.Id > 0 {
		model.CreateId = configId
		model.CreateAt = gtime.Now()
		_, err := m.Where(columns.Id, model.Id).Update(model)
		if err != nil {
			return err
		}
		_ = Log.Save(ctx, model, consts.UPDATE)
	} else {
		model.CreateId = configId
		model.CreateAt = gtime.Now()
		modelId, err := m.InsertAndGetId(model)
		if err != nil {
			return err
		}
		model.Id = modelId
		_ = Log.Save(ctx, model, consts.INSERT)
	}

	_ = s.Refresh(ctx)
	return nil
}

// Delete 删除配置
func (s *config) Delete(ctx context.Context, ids []int) error {
	_, err := dao.Config.Ctx(ctx).WhereIn(dao.Config.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}
	// 删除日志
	for _, id := range ids {
		_ = Log.Save(ctx, do.Config{
			Id: gconv.Int64(id),
		}, consts.DELETE)
	}
	_ = s.Refresh(ctx)
	return nil
}

func (s *config) Refresh(ctx context.Context) error {
	cache.InitCache(ctx)

	res, err := s.List(ctx, &v1.ConfigListReq{Enable: consts.EnableYes})
	if err != nil {
		return err
	}
	// 设置缓存
	cacheMap := g.Map{}
	for _, e := range res.List {
		cacheMap[e.Key] = e
	}
	err = cache.Instance().Set(ctx, consts.CacheConfig, cacheMap)
	if err != nil {
		return err
	}
	return nil
}

func (s *config) GetCache(ctx context.Context) (g.Map, error) {
	return cache.Instance().Get(ctx, consts.CacheConfig)
}
