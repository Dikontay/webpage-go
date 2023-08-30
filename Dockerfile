FROM golang:latest
WORKDIR /app
COPY ./ ./
RUN go build -o ascii-art-web .
CMD ["./webpage-go"]
EXPOSE 4000