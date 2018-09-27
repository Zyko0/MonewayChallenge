FROM golang
ADD . /go/src/github.com/Zyko0/MonewayChallenge
WORKDIR /go/src/github.com/Zyko0/MonewayChallenge
RUN go build -o moneway-api .
CMD ["./moneway-api"]