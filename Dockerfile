# 1st stage (build golang)
FROM golang:1.13.0-alpine as buildENV

LABEL Name="Fajar"
LABEL Version="0.5"

RUN apk update && apk add --no-cache git g++ ca-certificates


WORKDIR /src

#COPY go.mod go.sum ./

COPY . ./
#RUN make
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GO11MODULE=on TZ=Asia/Jakarta go build -a -installsuffix cgo -o /src/app

# 2nd stage (put app into container)

FROM alpine

RUN apk update && apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
RUN echo "Asia/Jakarta" >  /etc/timezone

WORKDIR /app

COPY --from=buildENV /src/app /app/
COPY .env /app/

EXPOSE 8008

ENTRYPOINT ["./app"]
