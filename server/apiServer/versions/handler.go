package versions

import (
	"distributed_storage_system/utils/elasticSearch"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

//name不为空:搜这个name的所有版本的元数据信息, name为空,搜所有元数据信息
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("apiServer get a version")
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	from := 0
	size := 1000
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	//每次尝试取一千个文件信息,直到取得所有相关文件信息
	for {
		metas, e := elasticSearch.SearchAllVersions(name, from, size)
		log.Printf("get %d file", len(metas))
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for i := range metas {
			b, _ := json.Marshal(metas[i])
			w.Write(b)
			w.Write([]byte("\n"))
		}
		if len(metas) != size {
			return
		}
		from += size
	}
	log.Println("apiServer versions end")
}
