package zoodmall_sdk

import (
	"context"
	"fmt"

	"github.com/ybbus/jsonrpc/v3"
)

const (
	methodProductsFind   = "LocalProductRpc.finds"
	methodProductsCreate = "LocalProductRpc.create"
	methodProductsUpdate = "LocalProductRpc.change"

	methodVariantsCreate = "LocalProductRpc.addSku"
	methodVariantsUpdate = "LocalProductRpc.changeSku"
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

func (c *Client) ProductsFind(ctx context.Context, data ReqProductsFind) (ProductList, error) {
	response, err := c.rpc.Call(ctx, methodProductsFind, data)
	if err != nil {
		return nil, fmt.Errorf("ProductsFind: failed http request: %w", err)
	}

	var res *resFindProducts
	if err := parseResponse(response, res); err != nil {
		return nil, err
	}

	return res.Products, nil
}

func (c *Client) ProductsCreate(ctx context.Context, data []ReqProductCreate) (
	ProductIdentifierList,
	error,
) {
	response, err := c.rpc.Call(ctx, methodProductsCreate, data)
	if err != nil {
		return nil, fmt.Errorf("ProductsCreate: failed http request: %w", err)
	}

	var res ProductIdentifierList
	if err := parseResponse(response, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ProductsUpdate(ctx context.Context, data []ReqProductUpdate) (
	ProductIdentifierList,
	error,
) {
	response, err := c.rpc.Call(ctx, methodProductsUpdate, data)
	if err != nil {
		return nil, fmt.Errorf("ProductsUpdate: failed http request: %w", err)
	}

	var res ProductIdentifierList
	if err := parseResponse(response, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) VariantsCreate(ctx context.Context, data []ReqVariantCreate) (
	ProductIdentifierList,
	error,
) {
	response, err := c.rpc.Call(ctx, methodVariantsCreate, data)
	if err != nil {
		return nil, fmt.Errorf("VariantsCreate: failed http request: %w", err)
	}

	var res ProductIdentifierList
	if err := parseResponse(response, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) VariantsUpdate(ctx context.Context, data []ReqVariantUpdate) (
	ProductIdentifierList,
	error,
) {
	response, err := c.rpc.Call(ctx, methodVariantsUpdate, data)
	if err != nil {
		return nil, fmt.Errorf("VariantsUpdate: failed http request: %w", err)
	}

	var res ProductIdentifierList
	if err := parseResponse(response, &res); err != nil {
		return nil, err
	}

	return res, nil
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
