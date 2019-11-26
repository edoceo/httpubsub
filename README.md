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

## Inspiration

A user @anderspitman posted an awesome creation called [PatchBay](https://patchbay.pub/) on [Hacker News](https://news.ycombinator.com/item?id=21639066).
Then someone asked for source.
Then after seeing the service I thought: "it's just channels"
And I posted that I expected to see a clone in Go real soon.
Then an associate called me out to post it.
