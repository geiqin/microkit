package gateway

//授权频道
type GrantChanel struct {
	Id     int    `json:"id"`   //id
	Name   string `json:"name"` //频道名称
	Routes []*GrantRoute
}

//授权路由
type GrantRoute struct {
	Id          int    `json:"id"`                            //id
	Name        string `json:"name"`                          //名称
	Path        string `json:"path"`                          //路径
	Method      string `json:"method"`                        //请求方式 ALL GET POST DELETE
	HasStore    bool   `json:"has_store" gorm:"default:0"`    //授权店铺
	HasUser     bool   `json:"has_user" gorm:"default:0"`     //授权用户
	HasCustomer bool   `json:"has_customer" gorm:"default:0"` //授权客户
	Unlimited   bool   `json:"unlimited" gorm:"default:1"`    //无限制（任何人都可以访问）
}
