FROM golang:1.16.0-buster
WORKDIR /app
RUN apt update && apt install -y gcc-mingw-w64-x86-64 g++-mingw-w64-x86-64 gcc
COPY . .
ENV CGO_ENABLED=1
ENV GOOS=windows
ENV GOARCH=amd64
ENV CC=x86_64-w64-mingw32-gcc
RUN go build
CMD /bin/bash
