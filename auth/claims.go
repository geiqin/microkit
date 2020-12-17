package auth

import "github.com/dgrijalva/jwt-go"

// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type AccessClaims struct {
	Mode  string       //模式: user , store_user , store_customer , store_app
	User  *GrantUser   //访问用户 当 mode 为 store_app时，该值允许为空
	Store *GrantStore  //当 mode为 store_user 或 store_customer 或 store_app 时，该店铺必须有值
	Limit *AccessLimit //访问限制
	// 使用标准的 payload
	jwt.StandardClaims
}

//授权店铺信息
type GrantStore struct {
	Id       int64  //店铺ID
	Name     string //店铺名称
	StoreKey string //店铺Key
}

//授权用户信息
type GrantUser struct {
	Id       int64             //用户ID
	Name     string            //用户账号
	Type     string            //类型: user ，customer
	MetaData map[string]string //附加信息
}

//访问限制
type AccessLimit struct {
	SessionKey string //会话key
	ClientIp   string //访客IP
	CreatedAt  string //Token创建时间
}
