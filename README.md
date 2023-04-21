# go-api-demo

## Get Started
### Create
```
>> go mod init go-api-demo
```

### Test
```
>> cd concurrency
>> go test
>> go test -v
```

```
>> go test ./concurrency
>> go test -v ./concurrency
```

or

```
>> go test ./...
>> go test -v ./...
```

## Details
1\.  
In Go, files do not have a concrete significance as they do in other programming languages like Java and Python. Instead, they are simply used to structure and arrange your code in a way that makes sense to you.

Variables are accessible throughout the entire package, which means that if there are two instances in different files under the same package, then both instances of the variable have package-wide visibility. As a result, the compiler will generate an error message if there are two global variables with the same name.