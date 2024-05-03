

FROM golang:1.22.1-alpine as modules 

# Set destination for COPY

# Download Go modules
COPY go.mod . go.sum  /modules/
WORKDIR /modules

RUN go mod download



#Step 2: Builder
FROM golang:1.22.1-alpine as builder 
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN mkdir tmp
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build   -o /bin/app .

# Step 3: Final
FROM scratch
COPY --from=builder /bin/app /app
COPY --from=builder /app/tmp /tmp
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8231
CMD ["/app"]