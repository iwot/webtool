package webtool

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	router := NewHttpRouter()
	router.SetBasePath("/test")
	router.SetRoute("/sample/", func(params map[string]string, w http.ResponseWriter, r *http.Request) (result string, err ActionError) {
		//fmt.Fprintf(w, "Accept, %q\n", html.EscapeString(r.URL.Path))
		result = "sample"
		return
	})
	router.SetRoute("/sample/<doc><year:[0-9]{4}>.html", func(params map[string]string, w http.ResponseWriter, r *http.Request) (result string, err ActionError) {
		//fmt.Fprintf(w, "Accept, %q\n", html.EscapeString(r.URL.Path))
		doc := params["doc"]
		year := params["year"]
		result = fmt.Sprintf("sample[%s][%s]", doc, year)
		return
	})

	res, _ := router.GetAction("/test/sample/")
	result, actionerr := res.Exec(nil, nil)
	assert.False(t, IsActionError(actionerr))
	assert.Equal(t, "sample", result)

	res, err := router.GetAction("/test/sample/abc1970.html")
	assert.Equal(t, nil, err)
	result, actionerr = res.Exec(nil, nil)
	assert.Equal(t, "sample[abc][1970]", result)

	res, err = router.GetAction("/test/SAMPLE/abc1970.html")
	assert.NotEqual(t, nil, err)
}
