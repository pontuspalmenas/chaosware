# chaosware

Chaosware is a lightweight Go http middleware that introduces configurable random issues such as `panic()`. Used for chaos testing.

### Installation
```
go get github.com/pontuspalmenas/chaosware
```

### Usage
```
import (
    "fmt"
    "github.com/pontuspalmenas/chaosware/chaosware"
    "net/http"
)

func main() {
    handler := http.HandlerFunc(defaultHandler)
    
    cw := chaosware.NewDefaultChaosMiddleware()
    chaosHandler := cw.ChaosHandler(handler)
    http.Handle("/", chaosHandler)
    
    fmt.Println("server listening on :8080")
    err := http.ListenAndServe(":8080", chaosHandler)
    if err != nil {
        panic(err)
    }
}
```

### Configuration

You may use environment variables or in-code configuration. If you use configuration in code, environment variables are not read.
Note that chaosware currently does not support config reload and only reads config at startup.

| Setting               | Values | Description                                                                                                                      |
|-----------------------|--------|----------------------------------------------------------------------------------------------------------------------------------|
| `CHAOSW_PANIC_CHANCE` | 0-100  | The likelyhood of a `panic()`, in percentage. `25` means on average every fourth request triggers `panic()`. `0` means disabled. |

#### Example environment variables:
```
CHAOSW_PANIC_ENABLED=true CHAOSW_PANIC_CHANCE=25 go run .    
```

#### Example settings
```
func main() {
    handler := http.HandlerFunc(defaultHandler)
    
    cw, err := chaosware.NewChaosMiddleware(&chaosware.Settings{
        PanicChance:  25,
        PanicEnabled: true,
    })    
    if err != nil {
        panic("failed to create middleware: " + err.Error())
    }
    
    ...
}

```