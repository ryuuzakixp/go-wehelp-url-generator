FROM golang:1.21.3-bullseye

RUN apt-get update && apt-get install -y \
    git \
    vim \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app-go-wehelp-url

RUN apt-get update 

COPY . .

CMD ["tail", "-f", "/dev/null"]
