# Naif - Build cross-platform GUI using HTML/JS/CSS
***
## Get Started
<a href="https://github.com/StevenZack/naif/blob/master/README-zh.md">中文文档</a><br><br><br>
- <a href="#android">Android</a>
- <a href="windows">Windows</a>
- <a href="linux">Linux</a>
<br><br>
### <a name="android">Android</a>

 1. Download Release <a href="https://github.com/StevenZack/naif/releases/download/latest/Naif-android.7z">here</a><br>
 2. Uncompress it<br>
 3. Open it width <a href="https://developer.android.com/studio/index.html">Android Studio</a><br>
 4. Create a index.html file , and write this on it:<br>
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
 5. Put index.html file into the html root directory:<b>Naif/app/src/main/assets/dir/</b> (if index.html exists,it's a demo file ,overwrite it)<br>
 6. Done ! Simple as hell ,now you can build & run this project.<br>
<br><br>
### <a name="windows">Windows</a>
1. Install <a href="http://golang.org/">Go</a>
2. Execute those command in CMD:<pre>go get github.com/lxn/walk
go get github.com/akavel/rsrc
go get github.com/StevenZack/naif
</pre>
3. Download project file <a href="https://github.com/StevenZack/naif/releases/download/latest/Naif-Windows-x86.7z">here</a>, and uncompress it.
4. Cd into the project folder , and execute:<pre>go build -ldflags="-H windowsgui"</pre>
5. Done ! you can run the output .exe file now .
<br><br>
### <a name="linux">Linux</a>
1. Download linux release <a href="https://github.com/StevenZack/naif/releases/download/latest/Naif-Linux-amd64.run">here</a>
2. Cd into the download folder , <pre>mkdir views</pre>
3. Create a index.html at the 'views' folder we just created with blow content:
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
4. Then we can run it:
<pre>
chmod +x Naif-Linux-amd64.run
./Naif-Linux-amd64.run
</pre>
> Unfortunately , we can't make Linux app native-like yet ( Due to the Linux Webview compatibility )


<br><br><br>
## How it works?
It's just a webview visiting a local web server(written in Go) listened on http://127.0.0.1:10246/

## Some APIs
#### 1. CacheFile 
> To solve the shortage of storaging,we offered a JavaScript Function to help you store remote file into local server
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
> 'http://101.200.54.63/' is the download link address of remote file<br>
'/img/webmtv.html' is the local path you want download into.It's a relative path, means you can visit the file on local server via url '/img/webmtv.html'.<br>
