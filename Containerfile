#Build stage
FROM golang:1.25.6 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-website

#Runtime stage
FROM scratch

WORKDIR /app

COPY --from=build /app/go-website .

COPY ./ui ./ui

COPY ./static ./static

EXPOSE 8080

USER 1000:1000

#This can be overridden from a systemd Quadlet
#Exec=/app/go-website --flag=/app/config.yaml
CMD [ "./go-website", "--contentdir", "./markdown"]
