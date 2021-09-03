FROM golang:1.17-rc-alpine

RUN apk upgrade --no-cache

RUN apk add --no-cache gcc git make musl-dev curl vim 
