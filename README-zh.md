# Naif - 用最简单的方式编写跨平台GUI
***
## 开始使用
<a href="https://github.com/StevenZack/naif/blob/master/README.md">English</a><br><br><br>
<br>
- <a href="#android">Android</a>
- <a href="windows">Windows</a>
- <a href="linux">Linux</a>
<br><br>
### <a name="android">Android</a>

  1. 下载.aar文件 <a href="https://github.com/StevenZack/naif/releases/download/latest/naif-android.aar">下载</a> <br>
 2. 打开Android Studio ,新建一个项目<br>
 3. 在 AndroidManifest.xml里面添加联网权限
 ```xml
    <uses-permission android:name="android.permission.INTERNET"/>
```
<br>
4.在Android Studio里面，依次打开菜单"File>New>New Module>import .jar/.aar package>" ,  选择我们刚刚下载好的.aar文件 <br>
5.然后在build.gradle里面添加一行如下代码，然后sync一下:
<pre> compile project(':naif')</pre>
 6. 在"YourProjectFolder/app/src/main/assets/dir"目录下创建一个index.html文件，并写入基本的hello world内容<br>
<br>
 7. 在 res/values/styles.xml 文件中添加一个新的style,名为“AppLauch":
<pre>
    <style name="AppLauch">
        <item name="android:windowBackground">@mipmap/ic_launcher</item>
    </style>
</pre>
并在AndroidManifest.xml中使用这个style:<br>
<pre>
        android:theme="@style/AppLauch">
</pre>
<br>
 8. 把你的activity_main.layout改成这种：<a href="https://github.com/StevenZack/naif-android-example/blob/master/app/src/main/res/layout/activity_main.xml" target="_blank">activity_main.xml</a><br>
 9. MainActivity.java 改成类似这样的 <a href="https://github.com/StevenZack/naif-android-example/blob/master/app/src/main/java/io/github/naife/stevenzack/myapplication/MainActivity.java" target="_blank">MainActivity.java</a>
<br>
10.OK !你现在可以运行该项目了 . 如有不懂，请看示例 <a href="https://github.com/StevenZack/naif-android-example">example project</a><br>

### <a name="windows">Windows</a>
1. 安装<a href="http://golang.org/">Go</a>
2. 在CMD中执行以下命令:<pre>go get github.com/lxn/walk
go get github.com/akavel/rsrc
go get github.com/StevenZack/naif
</pre>
3. 下载文件并解压：<a href="https://github.com/StevenZack/naif/releases/download/latest/Naif-Windows-x86.7z">点我下载</a>
4. 进入解压出来的目录，执行以下命令：<pre>go build -ldflags="-H windowsgui"</pre>
5. OK! 现在可以运行.exe后缀的文件了！
<br><br>

### <a name="linux">Linux</a>
1. 安装 <a href="http://golang.org/">Go</a>
2. 执行以下命令:
<pre>go get github.com/stevenzack/openurl
go get github.com/StevenZack/naif
</pre>
3.创建main.go文件 YourGoProject/main.go :
```go
package main

import (
	"fmt"
	"github.com/stevenzack/naif"
	"github.com/stevenzack/openurl"
)

func main() {
	port := naif.Start("./views/")
	openurl.Open("http://127.0.0.1:" + fmt.Sprintf("%v", port) + "/")
	fmt.Print("[q]Quit , [o]Open again\n")
	for {
		fmt.Print(">")
		s := ""
		fmt.Scanf("%s", &s)
		if s == "o" {
			openurl.Open("http://127.0.0.1:" + fmt.Sprintf("%v", port) + "/")
		} else if s == "q" {
			return
		}
	}
}
```
3. 在YourGoProject/views/文件夹里面创建index.html文件，并写入以下内容:
```html
<!DOCTYPE html>
<html>
<head>
	<title></title>
</head>
<body>
Hello
</body>
</html>
```
4. 然后运行查看效果:
<pre>
go run YourGoProject/main.go
</pre>
> 由于Linux Webview 的兼容性问题，我们暂时无法让程序像webview一样运行


## 原理
在本地127.0.0.1加随机端口 上运行了一个小型Web服务器，App启动时会通过WebView来访问

## APIs
#### 1. CacheFile 
> 为了解决Web程序不能存文件的问题，我们提供了一个JavaScript函数方便你进行文件存储
``` html
<!DOCTYPE html>
<html>
<head>
	<title></title>
	<script src="/naif.js"></script>
</head>
<body>
<span>hello</span>
<a href="/img/webmtv.html">webmtv</a>
<script type="text/javascript">
	cacheFile('http://101.200.54.63/','/img/webmtv.html')
</script>
</body>
</html>
```
> 'http://101.200.54.63/' 就是你要下载的文件的下载链接<br>
'/img/webmtv.html' 就是你要存储的路径，这是一个相对路径，意味着下载完成之后，你可以通过'/img/webmtv.html'链接来访问你存到本地服务器的文件.<br>
