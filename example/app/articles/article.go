package articles

import (
	"errors"
)

type Article struct {
	Title    string
	Body     string
	Datetime string
}

func Get(id int64) (article Article, err error) {
	switch id {
	case 1:
		article = Article{
			Title:    "記事１のタイトル",
			Body:     "記事１の内容",
			Datetime: "2013/07/27",
		}
	case 2:
		article = Article{
			Title:    "記事2のタイトル",
			Body:     "記事2の内容",
			Datetime: "2013/07/28",
		}
	default:
		err = errors.New("not found")
	}
	return
}
