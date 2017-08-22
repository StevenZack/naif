# Naif - Build Android App using HTML/JS/CSS
***
## Get Started
<a href="https://github.com/StevenZack/naif/blob/master/README-zh.md">中文文档</a><br><br><br>
 1. Download Release <a href="https://github.com/StevenZack/naif/releases">here</a><br>
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
 6. Done ! Now you can build & run this project.<br>
***

## How it works?
It's just a webview visiting a local web server(written in Go) listened on http://127.0.0.1:10246/

## Some APIs
#### 1. CacheFile 
> You can download file into local server via JavaScript
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
