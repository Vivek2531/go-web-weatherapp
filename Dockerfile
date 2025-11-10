FROM golang:1.25.4-alpine3.22 AS base

WORKDIR /app

COPY go.mod .

RUN go mod download 

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base

COPY --from=base /app/main .

COPY  --from=base /app/static ./static

EXPOSE 8000

CMD [ "./main" ]
