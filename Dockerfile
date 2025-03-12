FROM 374142386962.dkr.ecr.us-east-1.amazonaws.com/docker-hub/library/golang:1.24 AS build
LABEL authors="Brian Tajuddin"

RUN mkdir /code

COPY . /code

WORKDIR /code

ENV GOBIN /code/build
RUN mkdir build && go install


FROM scratch
COPY --from=build --chmod=555 /code/build/git-credential-github-app-sm /