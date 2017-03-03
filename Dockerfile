FROM alpine
RUN apk -U add ca-certificates
ADD cassandra_golang_example /
EXPOSE 8000
CMD ["/cassandra_golang_example"]