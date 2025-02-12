/*
 * Copyright 2025 Exactpro (Exactpro Systems Limited)
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package fetcher

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/th2-net/th2-common-go/pkg/grpc"
	"github.com/th2-net/th2-common-go/pkg/log"
	common "github.com/th2-net/th2-grpc-common-go"
	lwdp "github.com/th2-net/th2-grpc-lw-data-provider-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	LwdpBase64Format     = "BASE_64"
	LwdpJsonParsedFormat = "JSON_PARSED"
	lwdpService          = "DataProviderService"
)

var (
	logger = log.ForComponent("lwdp-fetcher")
)

type LwdpFetcher struct {
	dpClient lwdp.DataProviderClient
}

func NewLwdpFetcher(router grpc.Router) (*LwdpFetcher, error) {
	conn, err := router.GetConnection(lwdpService)
	if err != nil {
		return nil, fmt.Errorf("getting connection for '%s' service failure: %w", lwdpService, err)
	}
	client := lwdp.NewDataProviderClient(conn)
	return &LwdpFetcher{dpClient: client}, nil
}

func (f *LwdpFetcher) GetLastGroupedMessage(ctx context.Context, book string, group string, alias string, direction common.Direction, format string) (*lwdp.MessageGroupResponse, error) {
	if book == "" {
		return nil, errors.New("book can't be empty")
	}
	if group == "" {
		return nil, errors.New("group can't be empty")
	}
	if alias == "" {
		alias = group
	}
	if format != LwdpBase64Format && format != LwdpJsonParsedFormat {
		return nil, fmt.Errorf("unknown format '%s'. known values ['%s','%s']", format, LwdpBase64Format, LwdpJsonParsedFormat)
	}

	req := &lwdp.MessageGroupsSearchRequest{
		StartTimestamp:  timestamppb.Now(),
		EndTimestamp:    &timestamppb.Timestamp{Seconds: 0, Nanos: 0},
		BookId:          &lwdp.BookId{Name: book},
		MessageGroup:    []*lwdp.MessageGroupsSearchRequest_Group{&lwdp.MessageGroupsSearchRequest_Group{Name: group}},
		Stream:          []*lwdp.MessageStream{&lwdp.MessageStream{Name: alias, Direction: direction}},
		SearchDirection: lwdp.TimeRelation_PREVIOUS,
		ResponseFormats: []string{format},
	}
	stream, err := f.dpClient.SearchMessageGroups(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("searching message groups by [book=%s, group=%s, alias=%s, direction=%d] parameters failure: %w", book, group, alias, direction, err)
	}
	for {
		res, err := stream.Recv()
		switch err {
		case nil:
			logger.Trace().Str("operation", "GetLastGroupedMessage").Any("data", res.Data).Str("book", book).Str("group", group).Str("alias", alias).Any("direction", direction).Msg("received data")
		case io.EOF:
			logger.Trace().Str("operation", "GetLastGroupedMessage").Str("book", book).Str("group", group).Str("alias", alias).Any("direction", direction).Msg("end of stream")
			return nil, nil
		default:
			return nil, fmt.Errorf("receiving message groups by [book=%s, group=%s, alias=%s, direction=%d] parameters failure: %w", book, group, alias, direction, err)
		}
		msg := res.GetMessage()
		if msg != nil {
			logger.Info().Str("operation", "GetLastGroupedMessage").Any("message-id", msg.MessageId).Msg("received message")
			return res.GetMessage(), nil
		}
	}
}
