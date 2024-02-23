FROM busybox:1.36-uclibc as busybox

FROM golang:1.22-bookworm AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY lib/ ./lib

RUN go build -o /app/main -ldflags '-s -w' main.go


FROM gcr.io/distroless/base-debian12

COPY --from=busybox /bin/busybox /bin/busybox
RUN ["/bin/busybox", "--install", "/bin"]

USER nonroot:nonroot
WORKDIR /app
COPY --from=build /app/main /app/main
# ENV SUUMO_URL='https://suumo.jp/'

ENTRYPOINT ["sh", "-c", "/app/main > /app/output/$(date +%Y%m%d-%H%M%S).json"]
