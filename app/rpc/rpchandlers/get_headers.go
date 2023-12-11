package rpchandlers

import (
	"github.com/sedraxnet/sedrax/app/appmessage"
	"github.com/sedraxnet/sedrax/app/rpc/rpccontext"
	"github.com/sedraxnet/sedrax/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
