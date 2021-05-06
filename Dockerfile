FROM golang AS builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -a -o /app/build/ginstructure .

FROM alpine

WORKDIR /app
COPY --from=builder /app/build/ginstructure .
COPY --from=builder /app/config ./config

CMD [ "./ginstructure"]