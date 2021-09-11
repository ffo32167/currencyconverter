package rate_on_date

import "net/http"

type Rate struct {
	hello string
}

func NewRate() Rate {
	return Rate{hello: "hello, rate!"}
}

func (h Rate) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte(h.hello) // slice of bytes
	res.Write(data)
}
