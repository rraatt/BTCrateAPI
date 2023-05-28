FROM golang:1.19


WORKDIR /main


COPY ./src/main/go.mod ./src/main/go.sum ./
RUN go mod download


COPY ./src/main/ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main


EXPOSE 8080

# Run
CMD [ "./main" ]