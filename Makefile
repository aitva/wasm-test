wasm:
	GOARCH=wasm GOOS=js go build -o test.wasm .

build-server:
	go build -o server-app ./server/.

run: build-server wasm
	./server-app
