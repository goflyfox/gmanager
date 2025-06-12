package consts

const (
	UserTypeAdmin  = 1
	UserTypeNormal = 2

	RoleAdmin = "ADMIN" // 管理员角色

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
	CacheConfig = "config"
)
