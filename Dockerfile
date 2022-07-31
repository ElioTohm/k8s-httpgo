FROM golang:1.18-buster AS build

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN go test ./... -v
RUN go build -o /k8shttpgo

FROM gcr.io/distroless/base-debian10

COPY --from=build /k8shttpgo /k8shttpgo

EXPOSE 8080

ENTRYPOINT ["/k8shttpgo"]
