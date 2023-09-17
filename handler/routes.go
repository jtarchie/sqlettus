package handler

import (
	"context"

	"github.com/jtarchie/sqlettus/db"
	"github.com/jtarchie/sqlettus/router"
)

//nolint:funlen
func NewRoutes(
	ctx context.Context,
	client *db.Client,
) router.Command {
	commands := router.Command{
		"APPEND": appendRouter(ctx, client),
		"CONFIG": router.Command{
			"GET": router.Command{
				"save":       router.StaticResponseRouter(router.EmptyStringResponse),
				"appendonly": router.StaticResponseRouter("+no\r\n"),
			},
		},
		"COMMAND": router.Command{
			"DOCS": router.StaticResponseRouter(router.EmptyStringResponse),
		},
		"DECR":        decrRouter(ctx, client),
		"DECRBY":      decrByRouter(ctx, client),
		"DEL":         delRouter(ctx, client),
		"ECHO":        echoRouter(),
		"FLUSHALL":    flushAllRouter(ctx, client),
		"GET":         getRouter(ctx, client),
		"GETDEL":      getDelRouter(ctx, client),
		"GETRANGE":    getRangeRouter(ctx, client),
		"INCR":        incrRouter(ctx, client),
		"INCRBY":      incrByRouter(ctx, client),
		"INCRBYFLOAT": incrByFloatRouter(ctx, client),
		"MGET":        mgetRouter(ctx, client),
		"MSET":        msetRouter(ctx, client),
		"PING":        router.StaticResponseRouter("+PONG\r\n"),
		"SET":         setRouter(ctx, client),
		"STRLEN":      strlenRouter(ctx, client),
		"RPUSH":       rpushRouter(ctx, client),
	}

	commands["FLUSHDB"] = commands["FLUSHALL"]
	commands["SUBSTR"] = commands["GETRANGE"]
	commands["UNLINK"] = commands["DEL"]

	return commands
}
