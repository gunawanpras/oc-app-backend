FROM golang:1.24-alpine AS builder

ARG ENVIRONMENT=ENVIRONMENT

ENV ENV=${ENVIRONMENT}
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /oc-app-backend

COPY . .

COPY config.yaml config.yaml

COPY .env.development .env.application

# COPY tnsnames.ora tnsnames.ora

RUN apk update && apk add --no-cache gcc libc-dev && \
    go version && go mod download && go mod verify && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=1.0.0 -X main.buildTime=$(date +%Y-%m-%d) -s -w" -o ./oc-app-backend ./cmd

#
FROM alpine:latest

ENV TZ=Asia/Jakarta
ENV TNS_ADMIN=/opt/oracle/instantclient_21_10/network/admin
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_21_10:$LD_LIBRARY_PATH
ENV PKG_CONFIG_PATH=/opt/oracle/instantclient_21_10
ENV ORACLE_SID=XE
ENV GODROR_DRIVER_MODE=thick

WORKDIR /oc-app-backend

RUN mkdir -p /etc/ld.so.conf.d

RUN apk update && apk add --no-cache libaio libc6-compat gcompat tzdata && \
    wget https://download.oracle.com/otn_software/linux/instantclient/2110000/instantclient-basic-linux.x64-21.10.0.0.0dbru.zip && \
    mkdir /opt/oracle && \
    unzip instantclient-basic-linux.x64-21.10.0.0.0dbru.zip -d /opt/oracle && rm -f instantclient-basic-linux.x64-21.10.0.0.0dbru.zip && \     
# echo /opt/oracle/instantclient_21_10 > /etc/ld.so.conf.d/oracle-instantclient.conf && ldconfig && \
    ln -s /lib/libgcompat.so.0 /usr/lib/libresolv.so.2

COPY --from=builder /oc-app-backend/oc-app-backend .
COPY --from=builder /oc-app-backend/.env.application .env
COPY --from=builder /oc-app-backend/config.yaml .
# COPY --from=builder /oc-app-backend/tnsnames.ora /opt/oracle/instantclient_21_10/network/admin/tnsnames.ora
 
RUN chmod +x ./oc-app-backend

CMD ["./oc-app-backend"]