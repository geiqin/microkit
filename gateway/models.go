package gateway

//授权频道
type GrantChanel struct {
	Id     int    `json:"id"`                   //id
	Name   string `json:"name"  gorm:"size:50"` //频道名称
	Routes []*GrantRoute
}

//授权路由
type GrantRoute struct {
	Id            int    `json:"id"`                              //id
	GrantChanelId int    `json:"grant_channel_id"`                //id
	Name          string `json:"name" gorm:"size:50"`             //名称
	Path          string `json:"path" gorm:"size:100"`            //路径
	Method        string `json:"method" gorm:"size:10"`           //请求方式 ALL GET POST DELETE
	HasStore      bool   `json:"has_store" gorm:"default:0"`      //是否店铺API
	AllowUser     bool   `json:"allow_user" gorm:"default:0"`     //授权用户
	AllowCustomer bool   `json:"allow_customer" gorm:"default:0"` //授权客户
	Unlimited     bool   `json:"unlimited" gorm:"default:1"`      //无限制（任何人都可以访问）
}
