package apis

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/vtomar01/user-service/src/main/context"
	"github.com/vtomar01/user-service/src/main/dtos"
	"github.com/vtomar01/user-service/src/main/logging"
	"github.com/vtomar01/user-service/src/main/services"
	"net/http"
)

type UserController struct {
	orch services.UserOrch
}

func NewUserController(apiMux *mux.Router, orch services.UserOrch) *UserController {

	controller := &UserController{orch: orch}
	v0mux := apiMux.PathPrefix("/v1").Subrouter()

	v0mux.Methods("POST").Path("/user/").HandlerFunc(
		Make(
			CorrelationIdMiddleWare(controller.createUser),
		),
	)
	v0mux.Methods("GET").Path("/user/{id}/").HandlerFunc(
		Make(
			CorrelationIdMiddleWare(controller.getUser),
		),
	)
	v0mux.Methods("PUT").Path("/user/{id}/").HandlerFunc(
		Make(
			CorrelationIdMiddleWare(controller.updateUser),
		),
	)
	return controller
}

func (controller *UserController) createUser(req *http.Request) *Response {
	ctx := context.CreateLoggableContext(req, logging.Log)
	var request dtos.CreateUserRequest
	if err := read(req, &request); err != nil {
		return HandleError(err)
	}
	return Process(
		ctx,
		http.StatusOK,
		func() (interface{}, error) {
			return controller.orch.CreateUser(ctx, &request)

		},
	)

}

func (controller *UserController) updateUser(req *http.Request) *Response {
	ctx := context.CreateLoggableContext(req, logging.Log)
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		return HandleError(errors.New("invalid route"))
	}

	var request dtos.UpdateUserRequest
	if err := read(req, &request); err != nil {
		return HandleError(err)
	}
	return Process(
		ctx,
		http.StatusOK,
		func() (interface{}, error) {
			return controller.orch.UpdateUser(ctx, id, &request)
		},
	)
}

func (controller *UserController) getUser(req *http.Request) *Response {
	ctx := context.CreateLoggableContext(req, logging.Log)
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		return HandleError(errors.New("invalid route"))
	}
	return Process(
		ctx,
		http.StatusOK,
		func() (interface{}, error) {
			return controller.orch.GetUser(ctx, id)
		},
	)
}
