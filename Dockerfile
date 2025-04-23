## STEP 1 - BUILD

FROM golang:1.23.7-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api ./

##
## STEP 2 - DEPLOY
##

FROM alpine:latest

# Install bash, curl, and tar (for grpcurl)
RUN apk add --no-cache bash curl tar

# Install grpcurl
# RUN curl -sSL https://github.com/fullstorydev/grpcurl/releases/latest/download/grpcurl_1.9.3_linux_x86_64.tar.gz -o grpcurl.tar.gz && \
#     tar -xzf grpcurl.tar.gz && \
#     mv grpcurl /usr/local/bin/ && \
#     chmod +x /usr/local/bin/grpcurl && \
#     rm -f grpcurl.tar.gz

WORKDIR /

COPY --from=builder /app/api .
COPY --from=builder /app/cb.proto .

RUN chmod +x ./api

EXPOSE 50051

ENTRYPOINT ["./api"]
