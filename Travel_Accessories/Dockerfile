FROM golang:1.21

WORKDIR /Travel_Accessories

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /Travel_Accessories/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /greenlight

EXPOSE 4000
CMD ["/greenlight"]