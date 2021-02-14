go-appendable
=============

`go-appendable` provides File struct which supports append operation.

## Usage

```go
package main

import (
	"fmt"

	"github.com/moutend/go-appendable/pkg/appendable"
)

func main() {
	for i := 0; i < 5; i++ {
		write(i)
	}
}

func write(index int) {
	file, err := appendable.NewFile("output.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Fprintf(file, "Hello %d\n", index+1)
}
```

After executing the code above, you'll got:

```console
$ cat output.txt
Hello 1
Hello 2
Hello 3
Hello 4
Hello 5
```

## append.NewFile vs os.OpenFile

The `append.NewFile()` returns no error even if the file is not found. Usually, you should use the `os.OpenFile()` for handling such error.

## LICENSE

MIT

## Author

[Yoshiyuki Koyanagi <moutend@gmail.com>](https://github.com/moutend)
