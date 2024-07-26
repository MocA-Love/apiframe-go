# Apiframe-Go
Go library for [APIFRAME.PRO](https://apiframe.pro) (Midjourney API)

API Documentation is available at https://docs.apiframe.pro

## Install
```
go get github.com/MocA-Love/apiframe-go
```


## Usage
```go
apikey := "YOUR API KEY HERE"

client, err := apiframe.NewApiframeClient(apikey, true)
if err != nil {
    fmt.Println(err)
    return
}


task, err := client.Imagine("a nice day near a non-active volcano, photorealism, high details, high quality", "1:1", "fast", "", "")
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(task)
```
