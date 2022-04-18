# Suni: Streaming Uniq
Suni is a command-line tool similar to [uniq(1)](https://man7.org/linux/man-pages/man1/uniq.1.html). It takes in a stream of input and outputs a stream of compact JSON objects containing all lines seen so far, with a count of duplicates.

```bash
➜ ./suni <<EOF
heredoc> one
heredoc> two
heredoc> three
heredoc> two
heredoc> one
heredoc> EOF
{"one":1}
{"one":1,"two":1}
{"one":1,"three":1,"two":1}
{"one":1,"three":1,"two":2}
{"one":2,"three":1,"two":2}
```

## Building & Running
```bash
$ make        # For a native build
$ make sunix  # For a cross-platform build to AMD64.
```


## Usage Ideas
I'm using this in a pipeline for generating a _very_ simple analytics file, of this form:

```json
{
  "bede.io": 3,
  "bede.io/cv": 2,
  "bede.io/projects": 1,
}
```

Here's a rough approximation of the sort of script I'm running.

```bash
tail -f -n +1 $LOGFILE \
| jq -r --unbuffered '
  .request.host+.request.uri
  | select(. | test(
      # pattern for various boring stuff like images/txt files
    ) == false)
  | rtrimstr("/")' \
| ./suni
| # write each new line to analytics file on disk
```
