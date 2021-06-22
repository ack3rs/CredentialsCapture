package controllers

import "net/http"

func ServeLanding(w http.ResponseWriter, r *http.Request) {

	html := []byte(`<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Credentials Capture</title>
	  </head>
	  <body>
		<h1>Credentials Capture</h1>
	  </body>
	</html>`)

	w.Write(html)

}
