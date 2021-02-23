FROM golang:1.16.0-buster
WORKDIR /app
# RUN apt update && apt install -y gcc libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libx11-dev
RUN apt update && apt install -y gcc libgl1-mesa-dev xorg-dev
COPY . .
RUN go build
CMD /bin/bash
