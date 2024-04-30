FROM golang:alpine as build-go
WORKDIR /opt/build
COPY . /opt/build
RUN go build -o cmd

FROM golang:alpine
WORKDIR /opt/cmd
COPY --from=build-go /opt/build/cmd /opt/cmd/
CMD ["./cmd"]dock