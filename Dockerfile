FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./

RUN go mod download
RUN go get github.com/sirupsen/logrus
RUN go get github.com/gorilla/mux
#RUN go get github.com/prometheus/common/version

RUN go mod vendor
RUN go mod tidy

RUN go build -o /powerball-exporter

EXPOSE 8080

CMD [ "/powerball-exporter" ]
