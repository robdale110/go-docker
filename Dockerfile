FROM --platform=${BUILDPLATFORM} golang:1.14.3-alpine AS build
WORKDIR /src
ENV SGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .

FROM scratch as bin
COPY --from=build /out/example /