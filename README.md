# irange

[![Go Report Card](https://goreportcard.com/badge/github.com/spellr/irange)](https://goreportcard.com/report/github.com/spellr/irange)
[![GoDoc](https://godoc.org/github.com/spellr/irange?status.svg)](https://godoc.org/github.com/spellr/irange)

Produce a sequence of integers. Simple, but helpful

## Installation
`go get github.com/spellr/irange`

## Example

```golang
package main

import (
	"fmt"
	"github.com/spellr/irange"
)

func main() {
	fmt.Println(irange.Range(5))
}
```
```
[0 1 2 3 4]
```

## How to use

`Range` function takes at least 1, and up to 3 parameters (the rest are ignored):

If a single paramter is given:
```
Range(end)
```
`Range` will return the number from 0 to end, exclusive.

If more than 1 parameter is given:
```
Range(start, end, step?)
```
`Range` will return the numbers from `start` (inclusive) to `end` (exclusive), with `step` as the increment (or decrement) between every 2 numbers. `step` defaults to 1.

## More examples
```golang
Range(3, 5) // [3, 4]
Range(3, -2, -2) // [3, 1, -1]
```

## Contributing

Contributions are welcome.

When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other method with the owners of this repository before making a change.
