// Steve Phillips / elimisteve
// 2012.09.26

package tent

type URL struct {
	Host, Port, Path string
}

func (url *URL) String() string {
	proto := "https://"
	if url.Port == "80" {
		proto = "http://"
	}
	return proto + url.Host + url.Path
}

type Environment struct {
	Method  string
	URL     *URL
	Body    string
	Headers map[string]string
}

func NewEnv(host, port, path, method string) *Environment {
	url := URL{
		Host: host,
		Port: port,
		Path: path,
	}
	env := Environment{
		Method: method,
		URL: &url,
		Body: "",
		Headers: make(map[string]string),
	}
	return &env
}
