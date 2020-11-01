FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go mod vendor
# CGO_ENABLED=0 is required for some reason in this
RUN CGO_ENABLED=0 go build -o /out/backend .
RUN ls -la /out/backend
FROM scratch AS backend
COPY --from=build /out/backend /
EXPOSE 8081
ENTRYPOINT ["/backend"]