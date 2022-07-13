FROM golang:latest

WORKDIR /app
RUN apt-get update

RUN apt-get install ghostscript -y 

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3000

# ENV AWS_ACCESS_KEY_ID

# ENV AWS_SECRET_ACCESS_KEY


RUN go build

CMD [ "./grrow_pdf" ]