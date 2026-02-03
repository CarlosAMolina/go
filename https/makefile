# CGO_ENABLED = 0 to avoid these errors in the VPS when running the binary:
# ```
# /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.34' not found
# /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.32' not found
# ```
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o https .

run:
	go run . -cert server.cert -key server.key
