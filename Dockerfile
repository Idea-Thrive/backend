# build stage
#FROM golang:alpine AS build-env
#WORKDIR /app
#COPY . /app
#RUN go mod download
#RUN CGO_ENABLED=0 go build -o /backend .

# final stage
FROM alpine
WORKDIR /app
#COPY --from=build-env /backend /app/
#COPY --from=build-env /backend /app/config
COPY main /app/main
COPY config.yml /app/config.yml
EXPOSE 8080
ENTRYPOINT ["./main"]
CMD ["serve"]