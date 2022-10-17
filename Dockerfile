FROM golang
WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine

WORKDIR /dist
COPY --from=0 /build/main /dist/main


RUN apk add --no-cache texlive

ENV LISTEN_HOST=0.0.0.0
ENV LISTEN_PORT=80
ENV ENVIRONMENT=production
ENV GIN_MODE=release

EXPOSE 80
CMD ["/dist/main"]