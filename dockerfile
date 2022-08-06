FROM golang:1.16-buster AS build

WORKDIR /app

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./

# To get all the libraries without depending on local modules
RUN go get  -t -v ./...

RUN go build -o /docker-marvel

EXPOSE 8080

CMD [ "/docker-marvel" ]