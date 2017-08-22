# Naif - 用最简单的方式写Android
***
## 开始使用
<a href="https://github.com/StevenZack/naif/blob/master/README.md">English</a><br><br><br>
 1. 在<a href="https://github.com/StevenZack/naif/releases">这里</a><br>下载工程压缩包
 2. 解压<br>
 3. 用<a href="https://developer.android.com/studio/index.html">Android Studio</a><br>打开
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
***

## 原理
在本地http://127.0.0.1:10246/上运行了一个小型Web服务器，App启动时会通过WebView来访问

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
