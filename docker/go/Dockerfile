FROM golang:1.13

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apt-get update && apt-get -y upgrade && \
    apt-get install -y build-essential \
