# FROM golang:latest

# WORKDIR /app

# COPY ./go.mod .
# COPY ./go.sum .

# RUN go mod download

# COPY . .
# CMD ["go", "run", "main.go"]


# FROM golang:alpine
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN go build -o main . 
# CMD ["/app/main"]


FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

RUN go build -o main . 
CMD ["/app/main"]
