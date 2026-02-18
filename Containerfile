#Build stage
FROM golang:1.25.6 AS build

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-website

#Runtime stage
FROM scratch

WORKDIR /

COPY --from=build /go-website .

COPY ./ui ./ui

COPY ./static ./static

COPY ./markdown ./markdown

EXPOSE 8080

USER 1000:1000

CMD [ "./go-website" ]
