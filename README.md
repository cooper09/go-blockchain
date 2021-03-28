# GoLang Simple BlockChain
> Simple Example of a Golang Blockchain, essentially a linked list using hashed variables a keys. 

### Install golang (Good Luck with That )



Inspiration came from [npm-golang](https://github.com/chemdrew/npm-golang).

However, npm is not required to run this simple app

### Okay you got me, how do I use it?

```
$ go run main.go
```

```
A new blockchain with be created, 3 blocks are added with string data to displa using the previous hash address and creating a new hash address for the current block. 

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}//end Genesis
```

### See [example/](README.md) for more!
