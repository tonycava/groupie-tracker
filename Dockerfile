FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o back.exe .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

CMD [ "./back.exe" ]