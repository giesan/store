# Store - Read and write data structures

Store provides the ability to write the data structures to a file and read from a file in the [Go](https://go.dev/) programming language.

This module was inspired by [Perl Storable](https://perldoc.perl.org/Storable). The module provides possibility of the data structures in a JSON format or in a binary format. During the writing or reading process the file is locked.

Reading/writing in JSON format were implemented with the standard library [encoding/json](https://pkg.go.dev/encoding/json).

The module [encoding/gob](https://pkg.go.dev/encoding/gob) is used for reading/writing in binary format.

# Install

```
go get -u github.com/giesan/store
```

# Usage

## Writing in JSON format to a file:

```golang
package main

import (
	"log"
	"time"

	"github.com/giesan/store"
)

type message struct {
	Name      string
	Timestamp time.Time
	Payload   []byte
	Ssid      []uint32
}

func main() {
	msg := &message{
		Name:      "John",
		Timestamp: time.Now(),
		Payload:   []byte("hi"),
		Ssid:      []uint32{1, 2, 3},
	}

	if _, err := store.WriteJSON("./temp.json", msg); err != nil {
		log.Fatalln(err)
	}
}
```

## Read from a file written in JSON format:

```golang
package main

import (
	"log"
	"time"

	"github.com/giesan/store"
)

type message struct {
	Name      string
	Timestamp time.Time
	Payload   []byte
	Ssid      []uint32
}

func main() {
	var msg message

	if err := store.ReadJSON("./temp.json", &msg); err != nil {
		log.Fatalln(err)
	}

	log.Println(msg.Name)
}
```

## Writing in binary format to a file:

```golang
package main

import (
	"log"
	"time"

	"github.com/giesan/store"
)

type message struct {
	Name      string
	Timestamp time.Time
	Payload   []byte
	Ssid      []uint32
}

func main() {
	msg := &message{
		Name:      "John",
		Timestamp: time.Now(),
		Payload:   []byte("hi"),
		Ssid:      []uint32{1, 2, 3},
	}

	if _, err := store.WriteBinary("./temp.bin", msg); err != nil {
		log.Fatalln(err)
	}
}
```

## Read from a file written in binary format:

```golang
package main

import (
	"log"
	"time"

	"github.com/giesan/store"
)

type message struct {
	Name      string
	Timestamp time.Time
	Payload   []byte
	Ssid      []uint32
}

func main() {
	var msg message

	if err := store.ReadBinary("./temp.bin", &msg); err != nil {
		log.Fatalln(err)
	}

	log.Println(msg.Name)
}
```

# Author

Andrej Giesbrecht

# LICENSE

glg released under MIT license, refer [LICENSE](https://github.com/giesan/store/blob/main/LICENSE) file.