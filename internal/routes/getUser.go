package routes

import (
	"net/http"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
)

func (r *Router) getUser(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	userInfo, err := r.services.Authorization.GetUser(userVkInfo.Id)

	if err != nil {
		userType, err1 := r.services.UserType.GetUserType("simp")
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err1.Error(),
			})
			return
		}

		slaveLevel, err1 := r.services.SlaveLevel.GetSlaveLevel(0)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err1.Error(),
			})
			return
		}

		slaveStats := domain.SlaveStats{
			Level:         &slaveLevel,
			MoneyQuantity: 0,
		}

		slaveStatsId, err1 := r.services.SlaveStats.CreateSlaveStats(slaveStats)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err1.Error(),
			})
			return
		}
		slaveStats.Id = slaveStatsId

		defenderLevel, err1 := r.services.DefenderLevel.GetDefenderLevel(0)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err1.Error(),
			})
			return
		}

		defenderStats := domain.DefenderStats{
			Level:          &defenderLevel,
			DamageQuantity: 0,
		}
		defStatsId, err1 := r.services.DefenderStats.CreateDefenderStats(defenderStats)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err1.Error(),
			})
			return
		}
		defenderStats.Id = defStatsId

		fetter, err1 := r.services.Fetter.GetFetterByName("common")
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err1.Error(),
			})
			return
		}

		us := domain.User{
			Id:              userVkInfo.Id,
			SlavesCount:     0,
			Balance:         100,
			Gold:            0,
			Income:          0,
			LastUpdate:      time.Now(),
			JobName:         "",
			UserType:        &userType,
			SlaveStats:      &slaveStats,
			DefenderStats:   &defenderStats,
			PurchasePriceSm: 20,
			SalePriceSm:     0,
			PurchasePriceGm: 0,
			SalePriceGm:     0,
			HasFetter:       false,
			FetterTime:      time.Now(),
			FetterType:      &fetter,
		}

		err := r.services.Authorization.CreateUser(us)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		us.VkInfo = &userVkInfo

		c.JSON(http.StatusOK, gin.H{
			"user": us,
		})

	} else {
		userInfo.VkInfo = &userVkInfo

		c.JSON(http.StatusOK, gin.H{
			"user": userInfo,
		})
	}

}
