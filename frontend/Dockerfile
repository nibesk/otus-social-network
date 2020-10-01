FROM node:12.4

RUN mkdir frontend
WORKDIR /frontend

ADD ./entrypoint.sh ./entrypoint.sh

ENTRYPOINT ["/frontend/entrypoint.sh"]
