FROM golang:latest

WORKDIR /app

# RUN export GO111MODULE=on
# RUN go get -v github.com/aws/aws-sdk-go-v2

# RUN cd /build/grrow_pdf && go build

# RUN go mod init grrow_pdf

# RUN cd / && go get -u

# RUN go run main.go

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3000

# ENV PORT

# ENV AWS_ACCESS_KEY_ID

# ENV AWS_SECRET_ACCESS_KEY


RUN go build

CMD [ "./grrow_pdf" ]