FROM golang:1.15.2

RUN mkdir backend
WORKDIR /backend

ADD ./entrypoint.sh ./entrypoint.sh

EXPOSE 3001

ENTRYPOINT ["/backend/entrypoint.sh"]
