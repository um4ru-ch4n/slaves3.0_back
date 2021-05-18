package routes

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type JobName struct {
	Name    string `json:"job_name" binding:"required"`
	SlaveId int32  `json:"slave_id" binding:"required"`
}

func (r *Router) setJobName(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	var jobName JobName

	if err := c.ShouldBindJSON(&jobName); err != nil {
		c.JSON(http.StatusBadRequest, errors.Cause(err).Error())
		return
	}

	space := regexp.MustCompile(`\s+`)
	trimJobName := space.ReplaceAllString(strings.Trim(jobName.Name, " "), " ")

	if trimJobName == "" {
		c.JSON(http.StatusBadRequest, "Job name is empty")
		return
	}

	if len(trimJobName) > 50 {
		c.JSON(http.StatusBadRequest, "The length of job name is too long")
		return
	}

	if err := r.services.User.SetJobName(userVkInfo.Id, jobName.SlaveId, trimJobName); err != nil {
		c.JSON(http.StatusInternalServerError, errors.Cause(err).Error())
		return
	}

	c.Status(http.StatusOK)
}
