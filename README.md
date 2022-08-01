# crazy

An attempt to try to re-create the [`crazy`](https://en.wikipedia.org/wiki/Malbolge#Crazy_operation)
operation from [Malbolge](https://en.wikipedia.org/wiki/Malbolge) using a lookup
method with the truth table.

Your best bet for complete accuracy is `CrazyString` because `CrazyInt` takes in
integers (`int64`) which some of the ternary bits could get cut off by integer
parsers if they were zero.

## Why?

It's summer, it's basically hot as hell and Malbolge was named after the eighth
circle of hell. It fits.
