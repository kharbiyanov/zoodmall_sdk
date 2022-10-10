package zoodmall_sdk

import (
	"context"
	"fmt"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/billz-uz/zoodmall_sdk/pkg/structures"
)

const (
	methodCreateProduct = "LocalProductRpc.create"
	methodFindProducts  = "LocalProductRpc.finds"
)

type Client struct {
	rpc jsonrpc.RPCClient
}

func New(host, username, password string) *Client {
	rpc := jsonrpc.NewClientWithOpts(host, &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"X-RPC-Auth-Username": username,
			"X-Rpc-Auth-Password": password,
		},
	})

	return &Client{
		rpc: rpc,
	}
}

func (c *Client) CreateProduct(ctx context.Context, data structures.ReqCreateProduct) (
	id string, sku string, err error,
) {
	response, err := c.rpc.Call(ctx, methodCreateProduct, data)
	if err != nil {
		return "", "", fmt.Errorf("CreateProduct: failed http request: %w", err)
	}

	var res *structures.ResCreateProduct
	if err := parseResponse(response, res); err != nil {
		return "", "", err
	}

	return res.ID, res.SKU, nil
}

func (c *Client) FindProducts(ctx context.Context, data structures.ReqFindProducts) (*structures.ProductList, error) {
	response, err := c.rpc.Call(ctx, methodFindProducts, data)
	if err != nil {
		return nil, fmt.Errorf("FindProducts: failed http request: %w", err)
	}

	var res *structures.ResFindProducts
	if err := parseResponse(response, res); err != nil {
		return nil, err
	}

	return &res.Products, nil
}

func parseResponse(res *jsonrpc.RPCResponse, data interface{}) error {
	if res.Error != nil {
		return fmt.Errorf("RPCError: failed calling method: %w", res.Error)
	}

	if err := res.GetObject(&data); err != nil {
		return fmt.Errorf("GetObject: failed unmarshall: %w", res.Error)
	}

	return nil
}
