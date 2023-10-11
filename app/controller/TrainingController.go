package controller

import (
	"idstar-idp/rest-api/app/dto/request"
	"idstar-idp/rest-api/app/service"
	"idstar-idp/rest-api/app/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrainingController struct {
	svc *service.TrainingService
}

func NewTrainingController(svc *service.TrainingService) *TrainingController {
	return &TrainingController{svc}
}

func (ctrl *TrainingController) CreateTraining(ctx *gin.Context) {
	req := request.TrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(false)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.CreateTraining(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) UpdateTraining(ctx *gin.Context) {
	req := request.TrainingRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate(true)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.UpdateTraining(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) GetTrainingById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetTrainingById(id)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) GetTrainingList(ctx *gin.Context) {
	req := request.PagingRequest{}

	err := ctx.Bind(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	validFields := []string{"id", "tema", "pengajar", "created_date", "updated_date"}
	err = req.Validate(validFields)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	result, err := ctrl.svc.GetTrainingList(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponse(ctx, result)
}

func (ctrl *TrainingController) DeleteTraining(ctx *gin.Context) {
	req := request.IdRequest{}

	err := ctx.ShouldBindJSON(&req)
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = req.Validate()
	util.SetErrorResponse(ctx, err, http.StatusBadRequest)

	err = ctrl.svc.DeleteTraining(req)
	util.SetErrorResponse(ctx, err, http.StatusInternalServerError)

	util.SetSuccessResponseNoData(ctx)
}
