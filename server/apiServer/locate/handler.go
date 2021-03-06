package locate

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Handler 查询文件所在的数据节点,向客户端返回数据节点的地址
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//传入hash值,获取存有该文件的数据节点的地址
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}
