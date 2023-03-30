package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proffesional/model"
	"strconv"
)

func (h *Handler) AddGroup(c *gin.Context) {
	var group model.AddGroup
	err := c.BindJSON(&group)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.AddGroup(group)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) GetAllGroups(c *gin.Context) {
	group, err := h.services.GetAllGroups()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) GetGroupForId(c *gin.Context) {

}

func (h *Handler) ChangeSomeGroup(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var group model.AddGroup
	err = c.BindJSON(&group)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.services.ChangeSomeGroup(groupId, group)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteGroupForId(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.DeleteGroup(groupId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
