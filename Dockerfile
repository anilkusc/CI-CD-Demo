FROM golang:1.15 as BUILD
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app .

FROM alpine:3.8
RUN adduser -D hepsiburada
RUN mkdir /app/
WORKDIR /app
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
COPY --from=BUILD /bin/app .
RUN chown -R hepsiburada:hepsiburada /app
USER hepsiburada
ENTRYPOINT ["sh","/app/entrypoint.sh"]
CMD [""]
