package header

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const XRequestID = "x-request-id"
const TodoID = "todo-id"

func GetRequestID(ctx context.Context) string {
	var reqId string
	mdOut, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		values := mdOut.Get(XRequestID)
		if len(values) >= 1 {
			reqId = values[0]
		}
	}
	return reqId
}
