FROM golang:1.22.1 as builder

WORKDIR /dist

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o image_shifter

FROM scratch

COPY --from=builder /dist/image_shifter .
COPY --from=builder /dist/image.jpg .

CMD [ "./image_shifter" ]