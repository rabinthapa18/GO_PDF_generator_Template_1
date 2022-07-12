FROM golang:1.8

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get -v github.com/aws/aws-sdk-go-v2

RUN cd /build/grrow_pdf && go build

# RUN go mod init grrow_pdf

# RUN cd / && go get -u

# RUN go run main.go

# COPY . .

EXPOSE 3000

CMD [ "go","run", "main.go" ]