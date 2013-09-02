package webtool

type Executor struct {
	params map[string]string
	action ActionFunc
}

func (e *Executor) Exec(w http.ResponseWriter, r *http.Request) (string, ActionError) {
	return e.action(e.params, w, r)
}
