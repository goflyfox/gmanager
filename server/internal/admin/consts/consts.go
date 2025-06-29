package consts

const (
	CodeOK = 0

	UserTypeAdmin  = 1
	UserTypeNormal = 2

	DefaultAvatar = "/images/title1.png"
	RoleAdmin     = "ADMIN" // 管理员角色

	EnableYes = 1
	EnableNo  = 2

	MenuTypeMenu    = 1 //菜单
	MenuTypeCatalog = 2 //目录
	MenuTypeExtlink = 3 //外链
	MenuTypeButton  = 4 //按钮

	DefaultPassword = "123456"

	DataTypeKv       = 1 // 键值对
	DataTypeDict     = 2 // 字典列表
	DataTypeDictData = 3 // 字典数据
)

// log type
const (
	LOGIN      = "登录"
	LOGOUT     = "登出"
	INSERT     = "插入"
	UPDATE     = "更新"
	DELETE     = "删除"
	TypeEdit   = 2
	TypeSystem = 1
)

// 缓存 cache
const (
	CacheData     = "data"
	CacheConfig   = "config"      // 配置缓存
	CacheUserPerm = "userPerm_%d" // 用户按钮权限缓存
	CacheUserMenu = "userMenu_%d" // 用户菜单权限缓存
)
