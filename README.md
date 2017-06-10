# heads_or_tails

Flip a coin, how many times will it land with heads or tails X amount of times in a row?


# To run

## Command Line

To see the flags and descriptions run  `go install && heads_or_tails --help`

## Docker

In docker set the env vars in the compose file and run
```
docker-compose build && docker-compose up
```
or

```
docker build -t heads_or_tails .
docker run -e ENV_SET=1 \
    -e CONCURRENT_THREAD=1 \
    -e MAX_TOSSES=-1 \
    -e MIN_IN_ROW=15 \
    -e NUM_SIDES=2 \
    -e PRINT_EVERY=1000000 \
    -e VERBOSE=true \
    --name heads_or_tails heads_or_tails
```
