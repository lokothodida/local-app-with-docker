FROM golang:1.14-alpine

## The below is a mix of:
## https://vsupalov.com/docker-shared-permissions/
## https://stackoverflow.com/questions/49955097/how-do-i-add-a-user-when-im-using-alpine-as-a-base-image

ARG USER_ID
ARG USER

RUN adduser --disabled-password --gecos "" --uid $USER_ID $USER

USER $USER
