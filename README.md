# th2-lwdp-grpc-fetcher-go

This project provide ability to fetch data from lw-data-provider via gRPC

## Light Weight Data Provider Fetcher (LwdpFetcher)

This structure holds gRPC client made by [lw_data_provider.proto](https://github.com/th2-net/th2-lw-data-provider/blob/dev-version-2/grpc/src/main/proto/th2_grpc_lw_data_provider/lw_data_provider.proto) schema

### Constructor

`NewLwdpFetcher` uses gRPC router from gRPC common module to create instance of LwdpFetcher structure.

```go
package main

import (
    "github.com/th2-net/th2-common-go/pkg/factory"
    "github.com/th2-net/th2-common-go/pkg/modules/grpc"
    "github.com/th2-net/th2-lwdp-grpc-fetcher-go/pkg/fetcher"
)

func main() {
    newFactory := factory.New()
    defer newFactory.Close()

    if err := newFactory.Register(grpc.NewModule); err != nil {
        panic(err)
    }

    mod, err := grpc.ModuleID.GetModule(newFactory)
    if err != nil {
        panic(err)
    }
    router := mod.GetRouter()
    lwdp, err := fetcher.NewLwdpFetcher(router)
    if err != nil {
        panic(err)
    }
    // ... other code ...
}
```

### Methods

#### GetLastGroupedMessage 

Makes `DataProvider.SearchMessageGroups` gRPC request by parameters to get the last `*th2-grpc-lw-data-provider-go.MessageGroupResponse` message 
```go
import (
    // ... other imports ...
    "github.com/th2-net/th2-common-go/pkg/factory"
    "github.com/th2-net/th2-common-go/pkg/modules/grpc"
    "github.com/th2-net/th2-lwdp-grpc-fetcher-go/pkg/fetcher"
)

func main() {
    // ... other code ...
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000)*time.Millisecond)
    defer cancel()
    msg, err := lwdp.GetLastGroupedMessage(ctx, "book", "group", "alias", grpc_common.Direction_FIRST, fetcher.LwdpBase64Format)
    if err != nil {
        panic(err)
    }
    // ... other code ...
}
```