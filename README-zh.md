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

 1. 在<a href="https://github.com/StevenZack/naif/releases">这里</a>下载工程压缩包<br>
 2. 解压<br>
 3. 用<a href="https://developer.android.com/studio/index.html">Android Studio</a>打开<br>
 4. 创建一个index.html文件，写入一下内容:<br>
``` html
<html>
  <head>
  	<title></title>
  </head>
  <body>
    <span>hello</span>
  </body>
</html>
```
 5. 把这个文件放入html根目录:<b>Naif/app/src/main/assets/dir/</b> (里面有一个我写好的样例文件，删掉它就好)<br>
 6. OK ! 就这么简单，你现在可以导出或者运行这个工程<br>
<br><br>

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
1. 下载程序： <a href="https://github.com/StevenZack/naif/releases/download/latest/Naif-Linux-amd64.run">下载</a>
2. 进入程序所在目录 , 创建views文件夹<pre>mkdir views</pre>
3. 在views文件夹里面创建index.html文件，并写入以下内容:
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
4. 然后运行Naif-Linux-amd64.run查看效果:
<pre>
chmod +x Naif-Linux-amd64.run
./Naif-Linux-amd64.run
</pre>
> 由于Linux Webview 的兼容性问题，我们暂时无法让程序像本地应用一样运行


## 原理
在本地http://127.0.0.1:10246/ 上运行了一个小型Web服务器，App启动时会通过WebView来访问

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
