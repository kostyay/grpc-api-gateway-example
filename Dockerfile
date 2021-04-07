FROM alpine

ARG image_name
ENV SVC_ENV=$image_name
WORKDIR /app
COPY ./out /app/

ENTRYPOINT ./$SVC_ENV