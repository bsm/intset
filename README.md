# Intset [![Build Status](https://travis-ci.org/bsm/intset.png?branch=master)](https://travis-ci.org/bsm/intset)

Simplest-possible intset implementation - uses sorted slices.

### Example

```go
package main

import (
  "fmt"

  "github.com/bsm/intset"
)

func main() {
  // Create a new set
  set := intset.New(3)
  set.Add(2) // true
  set.Add(1) // true
  set.Add(3) // true
  set.Add(1) // false

  fmt.Println(set.Slice()) // [1, 2, 3]

  set.Exists(1) // true
  set.Exists(4) // false

  set.Remove(1) // true
  set.Remove(4) // false
  fmt.Println(set.Slice()) // [2, 3]
}
```

### Licence

```
Copyright (c) 2015 Black Square Media

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
