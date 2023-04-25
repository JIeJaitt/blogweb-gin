在项目中，需要涉及到MVC的三个层面，包括view，controller，还有model。

进入main.go,初始化路由，以及端口号。
根据浏览器输入的URL地址，在router路由器中找到对应的路由函数方法。
根据路由中URL后指定的函数，在Controller中找到对应的方法函数。
Controller中调用models关于数据库方面的方法函数；渲染html页面，执行js,css等效果

