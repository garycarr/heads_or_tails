# heads_or_tails

Flip a coin, how likely is it to land with heads or tails X amount of times in a row?

Either run go test, or to run in a docker container -

```
docker build -t heads_or_tails . && docker run --name heads_or_tails -d heads_or_tails && docker logs -f heads_or_tails
```
