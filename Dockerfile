#Build Stage
FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run Stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 4000
CMD [ "/app/main" ]