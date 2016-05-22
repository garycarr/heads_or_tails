# heads_or_tails

Flip a coin, how many times will it land with heads or tails X amount of times in a row?

Flags
```
	--concurrent-threads  # Number of concurrent threads
	--max-tosses  # Maximum number of tosses. Input -1 for no maximum
	--min-count  # Min in a row to start counting
	--num-sides  # How many sides. Example - 2 for a coin or 6 for a dice
	--print-every  # How many tosses to print at
	--verbose # Prints out at every printEvery

go install && heads_or_tails --print-every 500000000 --max-tosses=-1 --min-count=19 --concurrent-threads=1
```

To run in a docker container -
```
docker build -t heads_or_tails .
docker run -e ENV_SET=1 \
	-e CONCURRENT_THREAD=1 \
	-e MAX_TOSSES=100000000 \
	-e MIN_IN_ROW=15 \
	-e NUM_SIDES=2 \
	-e PRINT_EVERY=1000000 \
	-e VERBOSE=true \
	--name heads_or_tails -d heads_or_tails
docker logs -f heads_or_tails
```
