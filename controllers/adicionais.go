package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateAdicional(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Adicionais cadastrado com sucesso!"})
}

func ListAdicionais(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Adicionais cadastrado com sucesso!"})
}

func UpdateAdicional(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Adicionais cadastrado com sucesso!"})
}

func DeleteAdicional(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Adicionais cadastrado com sucesso!"})
}
