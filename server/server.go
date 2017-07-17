package server

import (
	"fmt"
	"net/http"
	"github.com/tyzual/tisq/util"
)

/*
HandleAddComment 添加评论Handler
*/
func HandleAddComment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("ERROR: USE POST"))
		return
	}
	util.Log(fmt.Sprintf("%v", r.Method))
}

/*
HandleCommentList 获取评论列表Handler
*/
func HandleCommentList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("ERROR: USE POST"))
		return
	}
	util.Log(fmt.Sprintf("%v", r.Method))
}
