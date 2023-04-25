在项目中，需要涉及到MVC的三个层面，包括view，controller，还有model。

进入main.go,初始化路由，以及端口号。
根据浏览器输入的URL地址，在router路由器中找到对应的路由函数方法。
根据路由中URL后指定的函数，在Controller中找到对应的方法函数。
Controller中调用models关于数据库方面的方法函数；渲染html页面，执行js,css等效果

如果您想使用curl命令访问 http://127.0.0.1:8080/register?username=JIeJaitt&password=12345678&repassword=12345678 ，您可以在终端中输入以下命令：
```bash
curl -X POST "http://127.0.0.1:8080/register" -d "username=JIeJaitt&password=12345678&repassword=12345678"
```
这个命令使用POST方法将数据作为表单参数发送到 http://127.0.0.1:8080/register 的端点。其中，-d选项指定了要发送的数据，即用户名、密码和确认密码。请注意，这个命令仅适用于在您的本地主机上运行的应用程序，如果您要访问的是互联网上的应用程序，那么您需要提供正确的URL和端点。