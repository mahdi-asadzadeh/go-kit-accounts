package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	docs "github.com/mahdi-asadzadeh/go-kit-accounts/clients/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-kit-accounts/clients/pb"
	"github.com/mahdi-asadzadeh/go-kit-accounts/clients/types"
)

func main() {

	// HTTP client
	HTTPCLI := HTTPClient{"http://0.0.0.0:8080"}

	// gRPC client
	GRPCCLI := gRPCClient{GrpcCli: InitUserServiceClient("50051")}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	// gRPC routes
	router.POST("/grpc/register", GRPCCLI.GRPCCreateUser)
	router.DELETE("/grpc/delete/:slug", GRPCCLI.GRPCDeleteUser)
	router.PUT("/grpc/update", GRPCCLI.GRPCUpdateUser)
	router.POST("/grpc/login", GRPCCLI.GRPCLoginUser)
	router.GET("/grpc/get/:slug", GRPCCLI.GRPCGetUser)

	// HTTP routes
	router.POST("/http/register", HTTPCLI.HTTPCreateUser)
	router.DELETE("/http/delete/:slug", HTTPCLI.HTTPDeleteUser)
	router.PUT("/http/update", HTTPCLI.HTTPUpdateUser)
	router.POST("/http/login", HTTPCLI.HTTPLoginUser)
	router.GET("/http/get/:slug", HTTPCLI.HTTPGetUser)

	// Swagger document route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8081")
}

func InitUserServiceClient(port string) pb.USerServiceClient {
	fmt.Println("Init user grpc service ...")
	client, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewUSerServiceClient(client)
}

func ReturnError(ctx *gin.Context, statusCode int, methd string, err string) {
	types.APIErrorResponse(ctx, statusCode, methd, err)
}

type gRPCClient struct {
	GrpcCli pb.USerServiceClient
}

type HTTPClient struct {
	ClientURL string
}

// ********************* gRPC *********************

// @Summary Create an user
// @Description Create an user
// @Tags gRPC client
// @Accept  json
// @Produce  json
// @Param request body types.CreateUserInput true "User"
// @Success 200 {object} types.CreateUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /grpc/register [POST]
func (gcli *gRPCClient) GRPCCreateUser(ctx *gin.Context) {
	var input types.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ReturnError(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	req := pb.CreateUserRequest{Email: input.Email, FullName: input.FullName, Password: input.Password}
	response, err := gcli.GrpcCli.CreateUser(context.Background(), &req)
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	result := types.CreateUserResponse{uint(response.GetId()), response.GetEmail(), response.GetFullName()}
	ctx.JSON(http.StatusCreated, result)
}

// @Summary Delete user by slug
// @Description Delete user by slug
// @Tags gRPC client
// @Accept  json
// @Produce  json
// @Param slug path string true "slug"
// @Success 200 {object} types.DeleteUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /grpc/delete/{slug} [DELETE]
func (gcli *gRPCClient) GRPCDeleteUser(ctx *gin.Context) {
	email := ctx.Param("slug")
	req := pb.DeleteUserRequest{Email: email}
	response, err := gcli.GrpcCli.DeleteUser(context.Background(), &req)
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "DELETE", err.Error())
		return
	}
	result := types.DeleteUserResponse{Ok: response.GetOk()}
	ctx.JSON(http.StatusOK, result)
}

// @Summary Update an user
// @Description Update an user
// @Tags gRPC client
// @Accept  json
// @Produce  json
// @Param request body types.UpdateUserInput true "User"
// @Success 200 {object} types.UpdateUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /grpc/update [PUT]
func (gcli *gRPCClient) GRPCUpdateUser(ctx *gin.Context) {
	var input types.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ReturnError(ctx, http.StatusBadRequest, "PUT", err.Error())
		return
	}
	req := pb.UpdateUserRequest{Email: input.Email, FullName: input.FullName}
	response, err := gcli.GrpcCli.UpdateUser(context.Background(), &req)
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "PUT", err.Error())
		return
	}
	result := types.UpdateUserResponse{uint(response.GetId()), response.GetEmail(), response.GetFullName()}
	ctx.JSON(http.StatusOK, result)
}

// @Summary Login an user
// @Description Login an user
// @Tags gRPC client
// @Accept  json
// @Produce  json
// @Param request body types.LoginUserInput true "User"
// @Success 200 {object} types.LoginUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /grpc/login [POST]
func (gcli *gRPCClient) GRPCLoginUser(ctx *gin.Context) {
	var input types.LoginUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ReturnError(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	req := pb.LoginUserRequest{Email: input.Email, Password: input.Password}
	response, err := gcli.GrpcCli.Login(context.Background(), &req)
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	result := types.LoginUserResponse{response.GetToken()}
	ctx.JSON(http.StatusOK, result)
}

// @Summary Detail user by slug
// @Description Detail article by slug
// @Tags gRPC client
// @Accept json
// @Product json
// @Param slug path string true "slug"
// @Success 200 {object} types.GetUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /grpc/get/{slug} [GET]
func (gcli *gRPCClient) GRPCGetUser(ctx *gin.Context) {
	email := ctx.Param("slug")

	req := pb.GetUserRequest{Email: email}
	response, err := gcli.GrpcCli.GetUser(context.Background(), &req)
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "GET", err.Error())
		return
	}
	result := types.GetUserResponse{uint(response.GetId()), response.GetEmail(), response.GetFullName()}
	ctx.JSON(http.StatusOK, result)
}

// ********************* HTTP *********************

// @Summary Detail user by slug
// @Description Detail article by slug
// @Tags HTTP client
// @Accept json
// @Product json
// @Param slug path string true "slug"
// @Success 200 {object} types.GetUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /http/get/{slug} [GET]
func (hcli *HTTPClient) HTTPGetUser(ctx *gin.Context) {
	email := ctx.Param("slug")

	response, err := http.Get(hcli.ClientURL + "/v1/get?email=" + email)
	if err != nil {
		ReturnError(ctx, 200, "GET", err.Error())
		return
	}
	if response.StatusCode == 404 {
		ReturnError(ctx, 404, "GET", "Record not found.")
		return
	}

	var result types.GetUserResponse
	json.NewDecoder(response.Body).Decode(&result)
	ctx.JSON(http.StatusOK, result)
}

// @Summary Create an user
// @Description Create an user
// @Tags HTTP client
// @Accept  json
// @Produce  json
// @Param request body types.CreateUserInput true "User"
// @Success 200 {object} types.CreateUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /http/register [POST]
func (hcli *HTTPClient) HTTPCreateUser(ctx *gin.Context) {
	var input types.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ReturnError(ctx, 400, "POST", err.Error())
		return
	}

	json_data, err := json.Marshal(&input)
	if err != nil {
		ReturnError(ctx, 400, "POST", err.Error())
		return
	}

	response, err := http.Post(hcli.ClientURL+"/v1/register", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		ReturnError(ctx, 400, "POST", err.Error())
		return
	}
	if response.StatusCode == 400 {
		ReturnError(ctx, 400, "POST", "Invalid data!")
		return
	}

	var result types.CreateUserResponse
	json.NewDecoder(response.Body).Decode(&result)
	ctx.JSON(http.StatusCreated, result)
}

// @Summary Delete user by slug
// @Description Delete user by slug
// @Tags HTTP client
// @Accept  json
// @Produce  json
// @Param slug path string true "slug"
// @Success 200 {object} types.DeleteUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /http/delete/{slug} [DELETE]
func (hcli *HTTPClient) HTTPDeleteUser(ctx *gin.Context) {
	email := ctx.Param("slug")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", hcli.ClientURL+"/v1/delete?email="+email, nil)
	if err != nil {
		ReturnError(ctx, 400, "DELETE", err.Error())
		return
	}

	// Fetch Request
	response, err := client.Do(req)
	if err != nil {
		ReturnError(ctx, 400, "DELETE", err.Error())
		return
	}

	if response.StatusCode == 404 {
		ReturnError(ctx, 404, "DELETE", "Record not found!")
		return
	}

	var result types.DeleteUserResponse
	json.NewDecoder(response.Body).Decode(&result)
	ctx.JSON(http.StatusOK, result)
}

// @Summary Update an user
// @Description Update an user
// @Tags HTTP client
// @Accept  json
// @Produce  json
// @Param request body types.UpdateUserInput true "User"
// @Success 200 {object} types.UpdateUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /http/update [PUT]
func (hcli *HTTPClient) HTTPUpdateUser(ctx *gin.Context) {
	var input types.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ReturnError(ctx, http.StatusBadRequest, "PUT", err.Error())
		return
	}

	client := http.Client{}
	json_data, err := json.Marshal(&input)

	req, err := http.NewRequest(http.MethodPut, hcli.ClientURL+"/v1/update", bytes.NewBuffer(json_data))
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "PUT", err.Error())
		return
	}

	response, err := client.Do(req)
	if err != nil {
		ReturnError(ctx, response.StatusCode, "PUT", err.Error())
		return
	}

	if response.StatusCode == 404 {
		ReturnError(ctx, 404, "PUT", "Record not found!")
		return
	}

	if response.StatusCode == 400 {
		ReturnError(ctx, 400, "PUT", "Invalid data!")
		return
	}
	var result types.UpdateUserInput
	json.NewDecoder(response.Body).Decode(&result)
	ctx.JSON(http.StatusOK, result)
}

// @Summary Login an user
// @Description Login an user
// @Tags HTTP client
// @Accept  json
// @Produce  json
// @Param request body types.LoginUserInput true "User"
// @Success 200 {object} types.LoginUserResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /http/login [POST]
func (hcli *HTTPClient) HTTPLoginUser(ctx *gin.Context) {
	var input types.LoginUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ReturnError(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	json_data, err := json.Marshal(&input)
	response, err := http.Post(hcli.ClientURL+"/v1/login", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		ReturnError(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	if response.StatusCode == 404 {
		ReturnError(ctx, 404, "POST", "Record not found!")
		return
	}

	var result types.LoginUserResponse
	json.NewDecoder(response.Body).Decode(&result)
	ctx.JSON(http.StatusCreated, result)
}
