function cacheFile(url,path)txt {
	var xhr=new XMLHttpRequest()
	xhr.onreadystatechange=function(evt) {
		if (xhr.readyState==4) {
			return xhr.responseText
		}
	}
	xhr.open("GET", "http://127.0.0.1:10246/cacheFile?url="+encodeURIComponent(url)+"&path="+encodeURIComponent(path))
 	xhr.setRequestHeader("Content-Type", "Application/x-www-form-urlencoded")
	xhr.send(null)
}