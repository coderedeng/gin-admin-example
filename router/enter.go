package router

type RouterGroup struct {
	ApiRouter
	UserRouter
	JwtRouter
	MenuRouter
	CasbinRouter
	AuthorityRouter
	AuthorityBtnRouter
}

var RouterGroupApp = new(RouterGroup)
