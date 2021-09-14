# syntax=docker/dockerfile:1
FROM alpine:latest

# update alpine and extra tools
RUN apk add --no-cache git make musl-dev go nmap

# configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:PATH
ENV PATH /usr/bin/go:PATH

RUN /bin/mkdir -p ${GOPATH}/src ${GOPATH}/bin

# get ffuf
RUN /usr/bin/go get -u github.com/ffuf/ffuf

# copy the project to the container
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN /usr/bin/go mod download

COPY main.go  ./
COPY api/**/*.go ./
COPY internal/**/*.go ./
COPY log/**/*.go ./

RUN /usr/bin/go build -o /code-runner

EXPOSE 5000

CMD [ "/code-runner" ]
