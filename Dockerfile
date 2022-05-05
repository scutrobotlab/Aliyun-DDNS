FROM golang
COPY . .
RUN go build

FROM scratch
COPY --from=0 ./aliyun-ddns ./
CMD ["./aliyun-ddns"]