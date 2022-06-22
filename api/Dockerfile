FROM golang:1.17-alpine as build

# Set destination for COPY
WORKDIR /dartgame

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
RUN apk add build-base
COPY . ./

# Build
RUN go build -o ./dart-api

FROM alpine

WORKDIR /app

COPY --from=build /dartgame/dart-api ./

COPY --from=build /dartgame/ ./

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8585

RUN cp .env.docker .env

# Run
CMD [ "./dart-api" ]
