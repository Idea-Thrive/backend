# build stage
FROM golang:alpine AS build-env
RUN apk add --no-cache make git
COPY . src/ideaThrive
WORKDIR src/ideaThrive
RUN make build

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/ideaThrive /app/
COPY --from=build-env /go/src/ideaThrive/config /app/config
EXPOSE 8080
ENTRYPOINT["./ideaThrive"]