/*
	Inspired by github.com/grpc-ecosystem/go-grpc-middleware/blob/master/logging/zap/ctxzap
	but return a default logger if no logger is found in the context.
	Also has a few modifications to work with echo.
*/

package ctxlogger
