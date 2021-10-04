package services

import (
	"github.com/gin-gonic/gin"
	"go-rest-crud/models"
	"go-rest-crud/utils"
	"log"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var r response
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "failed to create user", err.Error())
		c.JSON(http.StatusOK, r)
		return
	}

	validate := user.Validation()

	if validate != "" {
		r.setResponse("failed", "request body is not valid", validate)
		c.JSON(http.StatusBadRequest, r)
		return
	}

	err = models.CreateUser(&user)
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "failed to create user", err.Error())
		c.JSON(http.StatusOK, r)
		return
	} else {
		r.setResponse("success", "create user success", "")
		r.User = user
		c.JSON(http.StatusOK, r)
	}
}

func Update(c *gin.Context) {
	var r response
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "failed to update user. please, check your request", err.Error())
		c.JSON(http.StatusBadRequest, r)
		return
	}

	validate := user.Validation()

	if validate != "" {
		r.setResponse("failed", "request body is not valid", validate)
		c.JSON(http.StatusBadRequest, r)
		return
	}

	u := user
	err = models.GetUserByID(&u, u.ID)
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "user is not exist", err.Error())
		c.JSON(http.StatusNotFound, r)
		return
	} else {
		err = models.UpdateUser(&user)
		if err != nil {
			log.Println(err.Error())
			r.setResponse("failed", "failed to update user", err.Error())
			c.JSON(http.StatusInsufficientStorage, r)
			return
		} else {
			r.setResponse("success", "update user success", "")
			r.User = user
			c.JSON(http.StatusOK, r)
			return
		}
	}
}

func FindByID(c *gin.Context) {
	var r response
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "user not found", err.Error())
	}

	var user models.User
	err = models.GetUserByID(&user, uint32(id))
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "user is not exist", err.Error())
	} else {
		r.setResponse("success", "user data retrieved", "")
		r.User = user
		c.JSON(http.StatusOK, r)
		return
	}

	c.JSON(http.StatusNotFound, r)
}

func FindAll(c *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, response{
			Status:    "failed",
			Message:   "something went wrong",
			Error: err.Error(),
			Timestamp: utils.Timestamp(),
		})
		return
	} else {
		if len(user) > 0 {
			r := new(response)
			r.setResponse("success", "user data retrieved", "")
			r.User = user
			c.JSON(http.StatusOK, r)
			return
		} else {
			r := new(response)
			r.Status = "failed"
			r.Message = "no data found"
			r.Timestamp = utils.Timestamp()
			c.JSON(http.StatusNotFound, r)
		}
	}
}

func Delete(c *gin.Context) {
	var r response
	var scode int
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "user is not exist", err.Error())
		scode = 400
	}

	var user models.User
	err = models.GetUserByID(&user, uint32(id))
	if err != nil {
		log.Println(err.Error())
		r.setResponse("failed", "user is not exist", err.Error())
		scode = 404
	} else {
		err = models.DeleteUser(&user, uint32(id))
		if err != nil {
			r.setResponse("failed", "failed to delete user", err.Error())
			scode = 500
		} else {
			r.setResponse("success", "user deleted", "")
			scode = 200
		}
	}

	c.JSON(scode, r)
}
