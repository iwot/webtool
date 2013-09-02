package webtool

import (
	"bytes"
	"errors"
	"regexp"
)

const (
	DefaultPattern = "[^/]+"
)

func NewPathMatch() *PathMatch {
	ptn := `(\\.|.)`
	re := regexp.MustCompile(ptn)
	return &PathMatch{
		make(map[string][]string),
		re,
		make(map[string]string),
		make(map[string]*regexp.Regexp),
	}
}

type PathMatch struct {
	cache        map[string][]string
	splitPattern *regexp.Regexp
	regstrCache  map[string]string
	regexpCache  map[string]*regexp.Regexp
}

// <CONTENT> -> CONTENT
func TagContent(str string) (tagContent string, ok bool) {
	ok = false
	tagContent = str
	if len(str) >= 3 && str[0] == '<' && str[len(str)-1] == '>' {
		ok = true
		tagContent = str[1 : len(str)-1]
	}
	return
}

// タグ（<name:.+>）内では'<','>'はエスケープして使用すること。
// `/a/b/c<name:[^\>\[abc\]]+>d<test1:[a-z0-9]+>/<doctype>_<docname:[a-z]+>.html`
// -> ["/a/b/c" "<name:[^\\>\\[abc\\]]+>" "d" "<test1:[a-z0-9]+>" "/" "<doctype>" "_" "<docname:[a-z]+>" ".html"]
func (p *PathMatch) parse(pathPattern string) (parsed []string, err error) {

	tagStart := false
	matches := p.splitPattern.FindAllString(pathPattern, -1)
	var patternTemp bytes.Buffer
	var stringTemp bytes.Buffer
	parsed = make([]string, 0)

	for _, r := range matches {
		if tagStart {
			patternTemp.WriteString(r)
			switch r {
			case ">":
				tagStart = false
				parsed = append(parsed, patternTemp.String())
				patternTemp.Reset()
				continue
			default:
			}
		} else if r == "<" {
			tagStart = true
			patternTemp.WriteString(r)
			parsed = append(parsed, stringTemp.String())
			stringTemp.Reset()
		} else {
			stringTemp.WriteString(r)
		}
	}
	if stringTemp.Len() > 0 {
		parsed = append(parsed, stringTemp.String())
	}
	if patternTemp.Len() > 0 {
		err = errors.New("Invalid path pattern")
	} else {
		p.cache[pathPattern] = parsed
	}
	return
}

// <name:pattern> -> (?P<name>pattern)
// <name> -> (?P<name>defaultPattern)
// otherString -> regexp.QuoteMeta(otherString)
func (p *PathMatch) GetRegexp(str string, defaults map[string]string) (string, error) {
	// check cache
	reg, ok := p.regstrCache[str]
	if ok {
		return reg, nil
	}
	// start development
	preg := ""
	if tagContent, ok := TagContent(str); ok {
		res := regexp.MustCompile(":").Split(tagContent, 2)
		name := res[0]
		ptn, ok := defaults[name]
		if !ok {
			ptn = DefaultPattern
		}
		if len(res) == 2 {
			ptn = res[1]
		}
		if len(name) > 0 {
			preg = "(?P<" + regexp.QuoteMeta(name) + ">" + ptn + ")"
		} else {
			preg = regexp.QuoteMeta(str)
		}
	} else {
		preg = regexp.QuoteMeta(str)
	}
	p.regstrCache[str] = preg
	return p.regstrCache[str], nil
}

func (p *PathMatch) Parse(pathPattern string, defaults map[string]string) error {
	parsed, err := p.parse(pathPattern)
	if err != nil {
		return err
	}
	var pattern bytes.Buffer
	for _, row := range parsed {
		s, err := p.GetRegexp(row, defaults)
		if err != nil {
			return err
		}
		pattern.WriteString(s)
	}
	p.regexpCache[pathPattern] = regexp.MustCompile(`^` + pattern.String() + `$`)
	return nil
}

// return pathPattern
func (p *PathMatch) Match(path string) (result string, matches map[string]string, ok bool) {
	matches = make(map[string]string)
	for pathPattern, re := range p.regexpCache {
		ms := re.FindStringSubmatch(path)
		if len(ms) > 0 {
			ok = true
			result = pathPattern
			for idx, name := range re.SubexpNames() {
				if idx > 0 {
					matches[name] = ms[idx]
				}
			}
			break
		}
	}
	return
}
