FROM golang:1.15
RUN mkdir -p /api
WORKDIR /api
COPY . .
RUN go mod download
RUN go build -o main 
CMD ["./main"]
