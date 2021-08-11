FROM golang:alpine AS builder

WORKDIR /build

COPY . . 
RUN go mod download
RUN go build -o main .


WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .
RUN cp /build/docker-compose.yml .

FROM scratch

COPY --from=builder /dist/main .
COPY --from=builder /dist/.env .
COPY --from=builder /dist/docker-compose.yml .

EXPOSE 4040
ENTRYPOINT ["/main"]