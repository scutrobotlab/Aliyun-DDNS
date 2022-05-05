FROM golang
WORKDIR /build
COPY . .
RUN go build

FROM scratch
COPY --from=0 /build/aliyun-ddns ./
ENTRYPOINT ["./aliyun-ddns"]