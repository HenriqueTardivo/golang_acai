package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateAcai(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Açai cadastrado com sucesso!"})
}

func ListAcais(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Açai cadastrado com sucesso!"})
}

func ListAcaisById(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Açai cadastrado com sucesso!"})
}

func UpdateAcai(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Açai cadastrado com sucesso!"})
}

func DeleteAcai(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Açai cadastrado com sucesso!"})
}
