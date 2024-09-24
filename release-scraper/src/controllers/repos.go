// controllers/repos.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"platypi.dev/release-scraper/models"
)

// GET /repos
// Get all the repos

func FindRepos(c *gin.Context) {
	var repos []models.Repo
	models.DB.Find(&repos)

	c.JSON(http.StatusOK, gin.H{"data": repos})
}

// POST /repo
// Add a repo to the database. Requires org and project.
func AddRepo(c *gin.Context) {
	// use models.AddRepoInput as the model for input
	var input models.AddRepoInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repo := models.Repo{Org: input.Org, Repo: input.Repo, Name: input.Name}
	models.DB.Create(&repo)

	c.JSON(http.StatusOK, gin.H{"data": repo})
}

// GET /repo/:id
// get a specific repo from the database
func FindRepo(c *gin.Context) {
	var repo models.Repo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&repo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Org/repo record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"repo": repo})
}

// PATCH /repo/:id
// update a repo record with a new org or repo

func UpdateRepo(c *gin.Context) {
	// use models.Repo
	var repo models.Repo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&repo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Org/repo record not found"})
		return
	}
	var input models.UpdateRepoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&repo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": repo})
}

// DELETE /repo/:id
// Delete a repo from the DB

func DeleteRepo(c *gin.Context) {
	// check if the repo exists in the DB
	var repo models.Repo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&repo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Org/repo record not found"})
		return
	}

	models.DB.Delete(&repo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
