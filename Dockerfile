FROM golang:alpine as build
WORKDIR /app/src
RUN apk --no-cache add alpine-sdk opus-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN addgroup -S goffredo && adduser -S goffredo -G goffredo
RUN apk --no-cache add ffmpeg
WORKDIR /app
COPY --from=build /app/src/app ./
COPY views/ views/
COPY conf/ conf/
USER goffredo
CMD ["./app"]
