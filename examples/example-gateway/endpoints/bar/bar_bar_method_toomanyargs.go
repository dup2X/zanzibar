// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/clients/bar"
	"github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"

	clientTypeBar "github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/clients/bar/bar"

	clientTypeFoo "github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/clients/foo/foo"
)

// HandleTooManyArgsRequest handles "/bar/too-many-args-path".
func HandleTooManyArgsRequest(
	ctx context.Context,
	inc *zanzibar.IncomingMessage,
	g *zanzibar.Gateway,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}
	for _, header := range []string{"x-uuid", "x-token"} {
		h.Set(header, inc.Header.Get(header))
	}

	// Handle request body.
	rawBody, ok := inc.ReadAll()
	if !ok {
		return
	}
	var body TooManyArgsHTTPRequest
	if ok := inc.UnmarshalBody(&body, rawBody); !ok {
		return
	}
	clientRequest := convertToTooManyArgsClientRequest(&body)
	clientResp, err := clients.Bar.TooManyArgs(ctx, clientRequest, h)
	if err != nil {
		g.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		inc.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	// Handle client respnse.
	if !inc.IsOKResponse(clientResp.StatusCode, []int{200}) {
		g.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	b, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		inc.SendError(500, errors.Wrap(err, "could not read client response body:"))
		return
	}
	var clientRespBody bar.BarResponse
	if err := clientRespBody.UnmarshalJSON(b); err != nil {
		inc.SendError(500, errors.Wrap(err, "could not unmarshal client response body:"))
		return
	}
	response := convertTooManyArgsClientResponse(&clientRespBody)
	inc.WriteJSON(clientResp.StatusCode, response)
}

func convertToTooManyArgsClientRequest(body *TooManyArgsHTTPRequest) *barClient.TooManyArgsHTTPRequest {
	clientRequest := barClient.TooManyArgsHTTPRequest{}

	clientRequest.Foo = clientTypeFoo.FooStruct(body.Foo)

	clientRequest.Request = clientTypeBar.BarRequest(body.Request)

	return &clientRequest
}
func convertTooManyArgsClientResponse(body *bar.BarResponse) *bar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := bar.BarResponse{}
	return &downstreamResponse
}
