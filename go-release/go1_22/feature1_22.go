package go1_22

import (
	"fmt"
	"go/version"
	"net/http"
)

func IterateSliceInClose() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	for v := range s {
		defer func() {
			fmt.Print(v) // 976...0(go1.22+)  999...9 (before go1.22)
		}()
	}
}

func IsGoVersionOlder(cur, target string) bool {
	return version.Compare(cur, target) < 0
}

func RoutingPatternsServeMux() http.Handler {
	var mux http.ServeMux
	
	mux.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, _ = fmt.Fprint(w, id)
	})

	// 匹配所有剩余的片段
	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		paths := r.PathValue("path")
		_, _ = fmt.Fprint(w, paths)
	})

	return &mux
}
