FROM golang

ADD . /go/src/github.com/programadriano/apicontador-go/
COPY . /go/src/github.com/programadriano/apicontador-go/

RUN go get github.com/gorilla/mux

RUN go install github.com/programadriano/apicontador-go

ENTRYPOINT /go/bin/apicontador-go