FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/backend .
FROM scratch AS backend
COPY --from=build /out/backend /
ENTRYPOINT ["/backend"]