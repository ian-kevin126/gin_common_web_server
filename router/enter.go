package router

import "ewa_admin_server/router/system"

/*
这段代码定义了一个结构体类型 `RouterGroup` 和一个全局变量 `RouterGroupApp`，是为了更好地对路由进行管理和组织在一个项目中。
通过创建这个结构体和全局变量，可以在应用级别上为同一类路由设置公共属性和方法，并便于在其他组件中复用。

例如，在一个 Web 应用程序中，通常会有许多不同的路由需要被注册、管理和处理。将这些路由按照业务逻辑或功能特点进行分组，
可以让代码更加清晰易懂，同时也方便进行统一的权限控制、请求过滤等操作。使用结构体类型和全局变量来管理路由可以帮助开发者更好地组织代码，
减少重复代码的出现，提高可维护性和可扩展性。
*/

type RouterGroup struct {
	System system.RouterGroup
}

/*
声明了一个名为 RouterGroupApp 的全局变量，并将其初始化为 new(RouterGroup)，这意味着 RouterGroupApp 是一个指向 RouterGroup 类型的指针，并且它的值为 nil。
*/

var RouterGroupApp = new(RouterGroup)
