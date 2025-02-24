package execute

import (
	"context"

	"github.com/LinYuXinQing/gooo-ansible/pkg/stdoutcallback"
)

// Executor interface is satisfied by those types which has a Execute(context.Context,[]string,stdoutcallback.StdoutCallbackResultsFunc,...ExecuteOptions)error method
type Executor interface {
	Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, runinsName string, options ...ExecuteOptions) error
}
