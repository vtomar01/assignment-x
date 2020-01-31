package clients

import (
	"encoding/json"
	"errors"
	"github.com/vtomar01/user-service/src/main/context"
	"io/ioutil"
	"net/http"
)

func Execute(context *context.Context, client *http.Client,
	request *http.Request, response interface{}, responseStatusCode int) error {

	return ProcessRequest(context, client, request, response, responseStatusCode)
}

func ProcessRequest(ctx *context.Context, client *http.Client, req *http.Request,
	response interface{}, responseStatusCode int) error {

	req.Header.Add("Content-Type", "application/json")
	context.AttachCorrelationIdFromContext(req, ctx)

	var resp *http.Response
	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()

		if resp.StatusCode == responseStatusCode {

			var body []byte
			body, err = ioutil.ReadAll(resp.Body)
			if err == nil {
				err = json.Unmarshal(body, &response)
				if err == nil {
					return err
				}
			}
		}
	}

	ctx.Logger.Error("Error while processing request with url: ", req.URL, " error: ", err.Error())
	return errors.New("error in client processing")
}
