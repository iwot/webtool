package articles

import (
	"../../../../webtool"
	"../../libs"
	"net/http"
	"strconv"
)

func TopPage(transaction TransactionInterface) (message string, err webtool.ActionError) {
	//fmt.Fprintf(w, "Accept, %q\n", html.EscapeString(r.URL.Path))
	data := make(map[string]string)
	data["title"] = "タイトル"
	data["message"] = "Hello, world."
	view := libs.Viewer("templates/articles/top.html")
	if view != nil {
		view(transaction.Writer(), data)
	} else {
		err = webtool.NewActionError("404", "View Not Found")
	}
	message = "TopPage:" + r.URL.Path
	return
}

func ArticlePage(transaction TransactionInterface) (message string, err webtool.ActionError) {
	params := transaction.DefaultParams()
	message = "TopPage:" + r.URL.Path
	data := make(map[string]interface{})
	data["title"] = "タイトル"
	_, ok := params["id"]
	if ok {
		id, e := strconv.ParseInt(params["id"], 10, 0)
		if e != nil {
			err = webtool.NewActionError("404", "id mismatch")
			return
		}
		data["id"] = id
		article, e := Get(id)
		if e != nil {
			err = webtool.NewActionError("404", "article not found")
			return
		}
		data["article"] = article
	} else {
		err = webtool.NewActionError("404", "article not found")
		return
	}

	view := libs.Viewer("templates/articles/article.html")
	if view != nil {
		view(transaction.Writer(), data)
	} else {
		err = webtool.NewActionError("404", "view not found")
	}
	return
}
