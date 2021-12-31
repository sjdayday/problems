# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS build

WORKDIR /app

# https://stackoverflow.com/questions/64462922/docker-multi-stage-build-go-image-x509-certificate-signed-by-unknown-authorit
RUN apk --update add --no-cache ca-certificates openssl git tzdata && update-ca-certificates

# to run as non root, consider https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

#ADD kpathsmetrics ./kpathsmetrics
ADD main.go ./main.go
ADD go.mod ./go.mod
ADD go.sum ./go.sum

#WORKDIR /app/kpathsmetrics

RUN go mod download

RUN CGO_ENABLED=0 go build -o /problem-stats 


FROM scratch 

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /problem-stats /problem-stats

EXPOSE 8080

ENTRYPOINT [ "/problem-stats" ] 
