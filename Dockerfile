FROM golang:1.21-alpine

WORKDIR src/
ADD . ./
# install make cause alpine doesn't ship with it
RUN apk add --no-cache make
RUN go mod download
# compile application
RUN make build
# run application
CMD make run_prod
