wasm:
	GOARCH=wasm GOOS=js go build -o public/client.wasm ./client/.

build-server:
	go build -o server-app ./server/.

run: wasm build-server
	./server-app
