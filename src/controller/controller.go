package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/denizcamalan/key-value-store/model"
	"github.com/denizcamalan/key-value-store/repository"
	"github.com/gin-gonic/gin"
)

// createData godoc
// @Summary create data workflow
// @Description get {object}
// @Tags Workflow
// @Accept  json
// @Produce  json
// @Param workflow body model.Workflow true "workflow"
// @Success 200 {object} model.Workflow
// @Failure 400 {object} model.Message
// @Failure 404 {object} model.Workflow
// @Router /keys [post]
func CreateData(c *gin.Context) {

	var workflow model.Workflow
	if body, err := io.ReadAll(c.Request.Body); err != nil{
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	} else {
		json.Unmarshal(body, &workflow)
	}

	if err := repository.SettoRedis(workflow); err != nil {
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
    }

	if (model.Workflow{}) == workflow {
		c.JSON(http.StatusNotFound, workflow)
	} else {
		fmt.Println(repository.ListfromRedis())
		c.JSON(http.StatusOK, workflow)
	}
}

// getDataById godoc
// @Summary show workflow by ID
// @Description get string by ID
// @Tags Workflow
// @Accept  json
// @Produce  json
// @Param id path string true "Workflow ID"
// @Success 200 {object} model.Workflow
// @Failure 204 {object} model.Message
// @Router /keys/{id} [get]
func GetDataById(c *gin.Context) {

	var workflow model.Workflow

	workflow.ID = c.Param("id")

	strData, err := repository.GetRedis(workflow.ID)
	if  err != nil {
		c.JSON(http.StatusNoContent, model.Message{Message: err.Error()})
		return
	}else {
		json.Unmarshal([]byte(strData),&workflow)
	}
	if (model.Workflow{}) == workflow {
		c.JSON(http.StatusNotFound, workflow)
		return
	} else {
		c.JSON(http.StatusOK, workflow)
		return
	}
}

// @Summary check a model.Workflow item by ID
// @Accept  json
// @Produce json
// @Tags Workflow
// @Param id path string true "Workflow ID"
// @Success 204 string header
// @Failure 202 string header
// @Router /keys/{id} [head]
func CheckIfExist(c *gin.Context) {
	var workflow model.Workflow
	workflow.ID = c.Param("id")

	if !repository.CheckRedis(workflow.ID){
		c.Header("Message","there is no id "+`"`+workflow.ID+`"`)
		c.Status(http.StatusNoContent)
	}else{
		c.Header(workflow.ID, "Accepted")
		c.Status(http.StatusAccepted)
	}	
}

// UpdateData godoc
// @Summary update Workflow by ID
// @Tags Workflow
// @Description update by json Workflow
// @Accept  json
// @Produce json
// @Param id path string true "model.Workflow ID"
// @Param company body 	 model.Company true "Company"
// @Success 200 {object} model.Workflow
// @Failure 400 {object} model.Message
// @Failure 400 {object} model.Message
// @Failure 404 {object} model.Message
// @Router /keys/{id} [put]
func UpdateData(c *gin.Context) {
		
	var workflow model.Workflow

	workflow.ID = c.Param("id")

	if body, err := io.ReadAll(c.Request.Body); err != nil{
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	} else {
		json.Unmarshal(body, &workflow.Company)
	}

	if err := repository.UpdateRedis(workflow.ID,workflow.Company); err != nil {
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
		return
	}
	if (model.Workflow{}) == workflow {
		c.JSON(http.StatusNotFound, workflow)
	} else {
		c.JSON(http.StatusOK, workflow)
	}
}

// deleteDataByID godoc
// @Summary delete a model.Workflow item by ID
// @Description delete workflow by ID
// @Tags Workflow
// @Accept  json
// @Produce json
// @Param id path string true "model.Workflow ID"
// @Success 200 {object} model.Message
// @Failure 500 {object} model.Message
// @Router /keys/{id} [delete]
func DeleteDataByID(c *gin.Context) {
		
	var workflow model.Workflow

	workflow.ID = c.Param("id")
	
	err := repository.DeleteKey(workflow.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Message{Message: err.Error()})

	}else {
		r := model.Message{Message: "Deleted"}
		c.JSON(http.StatusOK, r)
	}
}

// listAll godoc
// @Summary get all items in the model.Workflow list
// @Tags Workflow
// @Accept  json
// @Produce json
// @Success 200 {array} []model.Workflow
// @Failure 500 {object} model.Message
// @Failure 400 {array} []model.Workflow
// @Router /keys [get]
func GetAll(c *gin.Context) {
	var workflows []model.Workflow

	workflows, err:=repository.ListfromRedis()
	if err != nil{
		c.JSON(http.StatusInternalServerError,model.Message{Message: err.Error()})
		return
	}
	if workflows == nil {
		c.JSON(http.StatusBadRequest, workflows)
		return
	}else {
		c.JSON(http.StatusOK, workflows)	
		return
	}
}

// deleteAll godoc
// @Summary delete all model.Workflow item
// @Tags Workflow
// @Accept  json
// @Produce json
// @Success 200 {object} model.Message
// @Failure 400 {object} model.Message
// @Router /keys [delete]
func DeleteAll(c *gin.Context) {
	
	if err:=repository.DeleteAll(); err != nil{
		c.JSON(http.StatusBadRequest, model.Message{Message: err.Error()})
	}else {
		c.JSON(http.StatusOK, model.Message{Message: "Deleted all keys"})
	}
}
