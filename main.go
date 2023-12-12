package main

import (
	"AFL_3_Rest_API/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/finances", getFinances)
	router.GET("/finances/:id", getFinancesByID)
	router.POST("/finances", postFinances)
	router.GET("/finances/accounts/", getAccounts)
	router.GET("/finances/accounts/:id", getAccountsByID)
	router.POST("/finances/accounts", postAccounts)
	router.Run("localhost:8080")
}

func postAccounts(c *gin.Context) {
	var newAccount entities.Account

	if err := c.BindJSON(&newAccount); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if newAccount.Name == "" || newAccount.Category == "" || newAccount.Number == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	entities.Accounts = append(entities.Accounts, newAccount)
	c.IndentedJSON(http.StatusCreated, newAccount)
}

func getAccountsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, a := range entities.Accounts {
		if a.Number == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "finance not found"})
}

func getFinances(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entities.Finances)
}

func postFinances(c *gin.Context) {
	var newFinance entities.Finance

	if err := c.BindJSON(&newFinance); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if newFinance.ID == "" || newFinance.Account == 0 || newFinance.Date == "" || newFinance.Description == "" || newFinance.Nominal == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	entities.Finances = append(entities.Finances, newFinance)
	c.IndentedJSON(http.StatusCreated, newFinance)
}

func getFinancesByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range entities.Finances {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "finance not found"})
}

func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entities.Accounts)
}
