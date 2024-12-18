FROM golang:1.23.4 AS build-stage

WORKDIR /app
COPY ./go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .
#COPY *.go ./
#COPY ./storage/*.go ./storage/

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-app

# run test
FROM build-stage AS run-test-stage
RUN go test -v ./..

# deploy
FROM gcr.io/distroless/static-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /golang-app /golang-app

EXPOSE 8088

USER nonroot:nonroot

ENTRYPOINT ["/golang-app"]