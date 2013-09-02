package webtool

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPathParse(t *testing.T) {
	m := NewPathMatch()

	pattern := `/a/b/c<name:[^\>\[abc\](]+>d<test1:[a-z0-9]+>/<doctype>_<docname:[a-z]+>.html`
	// -> ["/a/b/c" "<name:[^\\>\\[abc\\]]+>" "d" "<test1:[a-z0-9]+>" "/" "<doctype>" "_" "<docname:[a-z]+>" ".html"]
	parsed, err := m.parse(pattern)
	assert.Equal(t, nil, err)
	assert.Equal(t, 9, len(parsed))
	assert.Equal(t, "/a/b/c", parsed[0])
	assert.Equal(t, "<name:[^\\>\\[abc\\](]+>", parsed[1])
	tagContent, ok := TagContent(parsed[1])
	assert.True(t, ok)
	assert.Equal(t, "name:[^\\>\\[abc\\](]+", tagContent)
	assert.Equal(t, "d", parsed[2])
	assert.Equal(t, "<test1:[a-z0-9]+>", parsed[3])
	assert.Equal(t, "/", parsed[4])
	assert.Equal(t, "<doctype>", parsed[5])
	assert.Equal(t, "_", parsed[6])
	assert.Equal(t, "<docname:[a-z]+>", parsed[7])
	assert.Equal(t, ".html", parsed[8])
}

func TestPathMatch(t *testing.T) {
	m := NewPathMatch()

	pattern1 := `/a/b/c<name:[^\>\[abc\](]+>d<test1:[a-z0-9]+>/<doctype>_<docname:[a-z]+>.html`
	pattern2 := `/say/hello/<name>`
	pattern3 := `/member/<id:[0-9]{8}>/<page>`
	// -> ["/a/b/c" "<name:[^\\>\\[abc\\]]+>" "d" "<test1:[a-z0-9]+>" "/" "<doctype>" "_" "<docname:[a-z]+>" ".html"]
	defaults := make(map[string]string)
	err := m.Parse(pattern1, defaults)
	assert.Equal(t, nil, err)
	err = m.Parse(pattern2, defaults)
	assert.Equal(t, nil, err)
	err = m.Parse(pattern3, defaults)
	assert.Equal(t, nil, err)

	path := `/a/b/c0123fd1a2b/info_new.html`
	pathPattern, matches, ok := m.Match(path)
	assert.True(t, ok)
	assert.Equal(t, pattern1, pathPattern)
	assert.Equal(t, "0123f", matches["name"])
	assert.Equal(t, "1a2b", matches["test1"])
	assert.Equal(t, "info", matches["doctype"])
	assert.Equal(t, "new", matches["docname"])

	path = `/a/b/c0123fD1a2b/info_new.html`
	pathPattern, matches, ok = m.Match(path)
	assert.False(t, ok)

	path = `/say/hello/mister`
	pathPattern, matches, ok = m.Match(path)
	assert.True(t, ok)
	assert.Equal(t, pattern2, pathPattern)
	assert.Equal(t, "mister", matches["name"])

	path = `/say/hallo/mister`
	pathPattern, matches, ok = m.Match(path)
	assert.False(t, ok)

	path = `/member/01234567/news`
	pathPattern, matches, ok = m.Match(path)
	assert.True(t, ok)
	assert.Equal(t, pattern3, pathPattern)
	assert.Equal(t, "01234567", matches["id"])
	assert.Equal(t, "news", matches["page"])
}
