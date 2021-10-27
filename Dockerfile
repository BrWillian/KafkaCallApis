FROM golang:1.17.2-alpine3.13 AS build

RUN apk add alpine-sdk

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN GOOS=linux GOARCH=amd64 go build -o /KafkaCallApis -tags musl

FROM alpine:3.13

WORKDIR /

RUN apk --no-cache add libaio libnsl libc6-compat curl && \
cd /tmp && \
curl -o instantclient-basiclite.zip https://download.oracle.com/otn_software/linux/instantclient/instantclient-basiclite-linuxx64.zip -SL && \
unzip instantclient-basiclite.zip && \
mv instantclient*/ /usr/lib/instantclient && \
rm instantclient-basiclite.zip && \
ln -s /usr/lib/instantclient/libclntsh.so.19.1 /usr/lib/libclntsh.so && \
ln -s /usr/lib/instantclient/libocci.so.19.1 /usr/lib/libocci.so && \
ln -s /usr/lib/instantclient/libociicus.so /usr/lib/libociicus.so && \
ln -s /usr/lib/instantclient/libnnz19.so /usr/lib/libnnz19.so && \
ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
ln -s /lib/libc.so.6 /usr/lib/libresolv.so.2 && \
ln -s /lib64/ld-linux-x86-64.so.2 /usr/lib/ld-linux-x86-64.so.2

ENV ORACLE_BASE /usr/lib/instantclient
ENV LD_LIBRARY_PATH /usr/lib/instantclient
ENV TNS_ADMIN /usr/lib/instantclient
ENV ORACLE_HOME /usr/lib/instantclient

COPY --from=build /KafkaCallApis /KafkaCallApis

EXPOSE 10000

ENV DATABASE_SOURCE user="ANTT_OCORRENCIA" password="anttocorrencia" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=DTF-LBDEXP-DEV.datatraffic.com.br)(PORT=1521))(CONNECT_DATA=(Service_name=xe)))"
ENV KAFKA_BOOTSTRAP_SERVERS 192.168.250.11:9092
ENV APIOCR_URL http://192.168.250.10:5001/api/ocr
ENV APICAPACETE_URL http://192.168.250.10:5001/api/ocr
ENV APICLASSIFICADOR_URL http://192.168.250.10:5001/api/ocr

ENTRYPOINT ["/KafkaCallApis"]