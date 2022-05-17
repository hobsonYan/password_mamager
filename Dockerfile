FROM golang:alpine

ENV GO111MODULE = on \
    CGO_ENABLE = 0 \
    GOOS = linux \
    GOPATH = amd64 \
    GOPROXY = "https://goproxy.cn,direct"

COPY /bin/password_manager_linux .

CMD ["./password_manager_linux"]