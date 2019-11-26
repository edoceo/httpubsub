# HTTP Pub Sub

It's an HTTPS server that has a blocking read-write pipe based on URL paths.
Effectively, it's just paths connected via channel.

## Installation

    git clone
    go build
    go build -ldflags="-s -w"

## Running

You can run it using some of these options.
To run in the background use `nohup` and some redirection

    ./httpubsub --port=8080
    ./httpubsub --port=8080 --cert=[some file] --cert-key=[some other file]
    nohup ./httpubsub 2>&1 >httpubsub.log &
