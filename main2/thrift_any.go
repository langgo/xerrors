package main

import (
	"context"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/langgo/xerrors/main2/gen-go/base"
)

func MarshalAny(ctx context.Context, typeUrl string, msg thrift.TStruct) (*base.ThriftAny, error) {
	bs, err := thrift.NewTSerializer().Write(ctx, msg)
	if err != nil {
		return nil, nil
	}

	return &base.ThriftAny{
		TypeUrl: typeUrl,
		Value:   bs,
	}, nil
}

// typeUrl and msg should matched
func UnmarshalAny(ctx context.Context, any *base.ThriftAny, typeUrl string, msg thrift.TStruct) error {
	if typeUrl != any.TypeUrl {
		return fmt.Errorf("mismatched message type: got %q want %q", any.TypeUrl, typeUrl)
	}

	return thrift.NewTDeserializer().Read(msg, any.Value)
}
