# WASM test

This repository contains a Go application for the browser. The app is built with
`GOOS=js GOARCH=wasm` and can be run on Edge, Chrome & Firefox. It is able to update
the HTML using Go template library, enabling Go developper to share logic between
frontend and the backend. :tada:

There is also a server app, serving the files from disk, nothing fancy. The client
and the server can be run using `make run`.
