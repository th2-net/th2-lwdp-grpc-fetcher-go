package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/th2-net/th2-common-go/pkg/factory"
// 	"github.com/th2-net/th2-common-go/pkg/common"
// 	"github.com/th2-net/th2-common-go/pkg/modules/grpc"
// 	grpc_common "github.com/th2-net/th2-grpc-common-go"
// 	"github.com/th2-net/th2-lwdp-grpc-fetcher-go/pkg/fetcher"
// )

func main() {
// 	newFactory, err := factory.NewFromConfig(factory.Config{
// 		ConfigurationsDir: "cfg",
// 		FileExtension:     ".json",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer newFactory.Close()

// 	if err := newFactory.Register(grpc.NewModule); err != nil {
// 		panic(err)
// 	}

// 	mod, err := grpc.ModuleID.GetModule(newFactory)
// 	if err != nil {
// 		panic(err)
// 	}
// 	router := mod.GetRouter()
// 	lwdp, err := fetcher.NewLwdpFetcher(router)
// 	if err != nil {
// 		panic(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000)*time.Millisecond)
// 	defer cancel()
// 	msg, err := lwdp.GetLastGroupedMessage(ctx, "th2_5269", "mysql_01", "mysql_01", grpc_common.Direction_FIRST, fetcher.LwdpBase64Format)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("msg %w\n", msg)
}
