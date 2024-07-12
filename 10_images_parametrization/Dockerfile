ARG NG_VERSION=latest

FROM nginx:${NG_VERSION}

ARG ARG_FILE=default_file_name

RUN apt-get update && \
    apt-get install -y -q nginx

RUN touch /opt/${ARG_FILE}
RUN echo "some content for my file" > /opt/${ARG_FILE}

ENV NG_VERSION=$NG_VERSION

CMD ["nginx", "-g", "daemon off;"]
