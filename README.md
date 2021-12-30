# Store - Read and write data structures

Store provides the ability to write the data structures to a file and read from a file in the [Go](https://go.dev/) programming language.

The module provides possibility of the data structures in a JSON format or in a binary format. During the writing or reading process the file is locked.

Das Lesen/Schreiben im JSON-Format wurden mit der Standard-Bibliothek [encoding/json](https://pkg.go.dev/encoding/json) umgesetzt.

Reading/writing in JSON format were implemented with the standard library [encoding/json](https://pkg.go.dev/encoding/json).

For reading/writing in binary format the module is [encoding/gob](https://pkg.go.dev/encoding/gob) used.


# Usage

## Writing in JSON format to a file:

```
type message struct {
	Name      string
	Timestamp time.Time
	Payload   []byte
	Ssid      []uint32
}

msg := &message{
    Name:      "John",
    Timestamp: time.Now(),
    Payload:   []byte("hi"),
    Ssid:      []uint32{1, 2, 3},
}

if err := WriteJSON("./temp.json", msg); err != nil {
    log.Fatalln(err)
}
```

## Read from a file written in JSON format:

```
var msg message
if err := ReadJSON("./temp.json", &msg); err != nil {
    log.Fatalln(err)
}

fmt.Println(msg.Name) // Output John
```

## Writing in binary format to a file:

```
type message struct {
	Name      string
	Timestamp time.Time
	Payload   []byte
	Ssid      []uint32
}

msg := &message{
    Name:      "John",
    Timestamp: time.Now(),
    Payload:   []byte("hi"),
    Ssid:      []uint32{1, 2, 3},
}

if err := WriteBinary("./temp.bin", msg); err != nil {
    log.Fatalln(err)
}
```

## Read from a file written in binary format:

```
var msg message
if err := ReadBinary("./temp.bin", &msg); err != nil {
    log.Fatalln(err)
}
fmt.Println(msg.Name) // Output John
```

# Copyright

Copyright (c) 2021 Andrej Giesbrecht
