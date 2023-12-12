package client

import (
	"context"
	"github.com/sedraxnet/sedrax/cmd/sedraxwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/sedraxnet/sedrax/cmd/sedraxwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the sedraxwalletd server, and returns the client instance
func Connect(address string) (pb.sedraxwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("sedraxwallet daemon is not running, start it with `sedraxwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewsedraxwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
