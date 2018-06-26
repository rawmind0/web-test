package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	webHead = `<html>
<head>
	<title>Rancher</title>
	<style>
	body {
		background-color: white;
		text-align: center;
		padding: 50px;
		font-family: "Open Sans","Helvetica Neue",Helvetica,Arial,sans-serif;
	}

	#logo {
		margin-bottom: 40px;
	}
	</style>
</head>
<body>
	<img id="logo" src="img/rancher-logo.svg" alt="Rancher logo" width=400 />
	<h1>Hello world!</h1>`

	webTail = `</body>
</html>`
)

func getServices() map[string]string {
	k8s_services := make(map[string]string)

	for _, evar := range os.Environ() {
		show := strings.Split(evar, "=")
		if strings.HasSuffix(show[0], "_PORT") {
			k8s_services[strings.TrimSuffix(show[0], "_PORT")] = show[1]
		}
	}

	return k8s_services
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	k8s_services := getServices()

	fmt.Fprintln(w, webHead)
	fmt.Fprintln(w, "<h3>My hostname is ", hostname, "</h3>")
	if len(k8s_services) > 0 {
		fmt.Fprintln(w, "<h3>K8s services found</h3>")
		for k, v := range k8s_services {
			fmt.Fprintln(w, "<b>", k, "</b> ", v, "<br />")
		}
	}

	fmt.Fprintln(w, "<h3>Request info</h3>")
	fmt.Fprintln(w, "<b>Host:</b> ", r.Host, "<br />")
	fmt.Fprintln(w, "<b>Pod:</b> ", hostname, "</b><br />")
	for k, v := range r.Header {
		if strings.HasPrefix(k, "X-") {
			fmt.Fprintln(w, "<b>", k, "</b> ", v, "<br />")
		}
	}
	fmt.Fprintln(w, webTail)
}

func main() {
	web_port := os.Getenv("WEB_PORT")
	if web_port == "" {
		web_port = "8080"
	}

	fmt.Println("Running web-test service at", web_port, "port")
	http.HandleFunc("/", handler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(os.Getenv("PWD")))))
	http.ListenAndServe(":" + web_port, nil)
}
