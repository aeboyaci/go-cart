FROM golang:1.19

ARG DB_URL
ARG JWT_SECRET

ENV DB_URL=$DB_URL
ENV JWT_SECRET=$JWT_SECRET

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-cart ./cmd/

EXPOSE 1323

CMD ["/go-cart"]