package proxy

type application struct {

}

func (a *application)handleRequest(url, method string) (int, string)  {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User creaate"
	}
	return 404, "not found"
}
