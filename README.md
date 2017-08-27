# Naif - Build cross-platform GUI using HTML/JS/CSS
***
## Get Started
<a href="https://github.com/StevenZack/naif/blob/master/README-zh.md">中文文档</a><br><br><br>
- <a href="#android">Android</a>
- <a href="windows">Windows</a>
- <a href="linux">Linux</a>
<br><br>
### <a name="android">Android</a>

 1. Download aar file <a href="https://github.com/StevenZack/naif/releases/download/latest/naif-android.aar">here</a> ( 中国镜像下载链接<a href="https://github.com/StevenZack/naif/releases/download/latest/naif-android.aar">下载 )</a><br>
 2. Open Android Studio , create a new project<br>
 3. Add internet permission on AndroidManifest.xml
 ```xml
    <uses-permission android:name="android.permission.INTERNET"/>
```
<br>
 4. Create a index.html in YourProjectFolder/app/src/main/assets/dir/ folder with baisc content<br>
<br>
 5. Add a new style (named "AppLauch") in your project res/values/styles.xml file
<pre>
    <style name="AppLauch">
        <item name="android:windowBackground">@mipmap/ic_launcher</item>
    </style>
</pre>
And use this style as your application theme ,  in your AndroidManifest.xml theme:<br>
<pre>
        android:theme="@style/AppLauch">
</pre>
<br>
 6. Make your activity_main.layout like this <a href="https://github.com/StevenZack/naif-android-example/blob/master/app/src/main/res/layout/activity_main.xml">activity_main.xml</a><br>
 7. MainActivity.java like this <a href="https://github.com/StevenZack/naif-android-example/blob/master/app/src/main/java/io/github/naife/stevenzack/myapplication/MainActivity.java">MainActivity.java</a>
<br>
8.On Android Studio, open "File>New>New Module>import .jar/.aar package>" ,  choose the .aar file we just downloaded  <br>
9.Then add this line to your build.gradle file, and then sync project:
<pre> compile project(':naif')</pre>
10.Done ! Now you can run your project<br>

### <a name="windows">Windows</a>
1. Install <a href="http://golang.org/">Go</a>
2. Execute those command in CMD:<pre>go get github.com/lxn/walk
go get github.com/akavel/rsrc
go get github.com/StevenZack/naif
</pre>
3. Download project file <a href="https://github.com/StevenZack/naif/releases/download/latest/Naif-Windows-x86.7z">here</a>, and uncompress it.<br>
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
