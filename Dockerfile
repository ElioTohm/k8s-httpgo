FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .

RUN go mod vendor

RUN go test ./... -v
RUN go build -o /k8shttpgo

FROM gcr.io/distroless/base-debian10

COPY --from=build /k8shttpgo /k8shttpgo

EXPOSE 8080

ENTRYPOINT ["/k8shttpgo"]
