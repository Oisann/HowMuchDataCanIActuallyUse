FROM golang:1.11-alpine3.8 AS BUILDER
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=BUILDER /app/main /

# Download speed in Kbps
ENV download=10000

# Data cap and data used in MB
ENV cap=1000000 used=0

# Renew date (1st of every month)
ENV renew=1

# Percentage used before we actually care
ENV percentage=50

CMD ["/main"]
