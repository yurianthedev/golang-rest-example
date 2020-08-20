FROM golang:lastest

RUN make production
RUN make production-run

EXPOSE 80