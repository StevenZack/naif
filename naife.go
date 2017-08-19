package naif

import (
	"fmt"
	// download "github.com/joeybloggs/go-download"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	Domain    string = "127.0.0.1:10246"
	IsRunning bool   = false
	Pkg       string
)

func Start(str string) {
	Pkg = str
	if strings.LastIndex(str, "/") != len(str)-1 {
		Pkg = str + "/"
	}
	if !IsRunning {
		go Run()
	}
}
func Run() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(Pkg))))
	http.HandleFunc("/cacheFile", cacheFile)
	http.HandleFunc("/naif.js", naif)
	IsRunning = true
	err := http.ListenAndServe(Domain, nil)
	if err != nil {
		fmt.Println(err)
	}
}
func cacheFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	url := r.FormValue("url")
	path := r.FormValue("path")
	if url == "" || path == "" {
		fmt.Fprint(w, "ERR:url or path empty")
		return
	}
	fmt.Fprint(w, downloadFile(url, path))
}
func downloadFile(url string, filename string) string {
	// f, err := download.Open(url, nil)
	rp, err := http.Get(url)
	if err != nil {
		return "ERR:open err , " + err.Error()
	}
	defer rp.Body.Close()
	checkDir(Pkg + filename)
	mf, err := os.Create(Pkg + filename)
	if err != nil {
		return "ERR:create, " + err.Error()
	}
	defer mf.Close()
	_, err = io.Copy(mf, rp.Body)
	if err != nil {
		return "ERR:copy," + err.Error()
	}
	return "OK"
}
func checkDir(str string) {
	fmt.Println("checkDir: str=", str)
	n := strings.LastIndex(str, "/")
	if n == len(str)-1 {
		fmt.Println("checkDir:2")

		err := os.MkdirAll(str, 0755)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	fmt.Println("checkDir:1")
	err := os.MkdirAll(str[:n+1], 0755)
	if err != nil {
		fmt.Println(err)
	}
}
func naif(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `function cacheFile(url,path) {
	var xhr=new XMLHttpRequest()
	xhr.open("GET", "http://127.0.0.1:10246/cacheFile?url="+encodeURIComponent(url)+"&path="+encodeURIComponent(path))
 	xhr.setRequestHeader("Content-Type", "Application/x-www-form-urlencoded")
	xhr.send(null)
}`)
}
