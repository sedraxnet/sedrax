ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as kaspad
FROM ${KASPAMINER_IMAGE} as sedraxminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/sedraxnet/sedrax

WORKDIR /go/src/github.com/sedraxnet/sedrax

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=kaspad /app/ /app/
COPY --from=sedraxminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/sedraxnet/sedrax/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]