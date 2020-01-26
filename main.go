package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Company contains information about a single company
type Company struct {
	ID            int    `json:"id" binding:"required"`
	Likes         int    `json:"likes"`
	Company       string `json:"company" binding:"required"`
	FoundedYear   int    `json:"founded_year" binding:"required"`
	Description   string `json:"description" binding:"required"`
	OfficialURL   string `json:"official_url" binding:"required"`
	CrunchbaseURL string `json:"crunchbase_url" binding:"required"`
}

var companies = []Company{
	Company{1, 0, "120WaterAudit", 2016, "120WaterAudit offers cloud-based water management software.", "https://120wateraudit.com/", "https://www.crunchbase.com/organization/120wateraudit#section-overview"},
	Company{2, 0, "3AM Innovations", 2015, "3AM has designed an IPS solution that utilizes state-of-the-art sensors, communication technologies, and software called FLARE.", "https://www.3aminnovations.com/", "https://www.crunchbase.com/organization/3am-innovations"},
	Company{3, 0, "Accela", 1999, "Accela provides civic engagement solutions and improves core processes for city, county, state and federal governments.", "https://www.accela.com/", "https://www.crunchbase.com/organization/accela#section-overview"},
	Company{4, 0, "Acivilate", 2014, "They are a Software-as-a-Service (SaaS) company committed to reducing recidivism.", "https://acivilate.com/", "https://www.crunchbase.com/organization/acivilate"},
	Company{5, 0, "APPCityLife®, Inc.", 2009, "APPCityLife is an end-to-end platform transforming the way cities and agencies produce mobile apps.", "https://www.appcitylife.com/", "https://www.crunchbase.com/organization/appcitylife"},
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve fronted static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// Our API will consit of just two routes
		// /companies - which will retrieve a list of companies a user can see
		// /companies/like/:companyID - which will capture likes sent to a particular company
		api.GET("/companies", CompanyHandler)
		// api.POST("/companies/like/:companyID", LikeCompany)
	}

	// Start and run the server
	router.Run(":3000")
}

// CompanyHandler retrieves a list of available companies
func CompanyHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, companies)
}

// // LikeCompany increments the likes of a particular Company Item
// func LikeCompany(c *gin.Context) {
// 	// confirm Company ID sent is valid
// 	// remember to import the `strconv` package
// 	if companyid, err := strconv.Atoi(c.Param("companyID")); err == nil {
// 		// find company, and increment likes
// 		for i := 0; i < len(companies); i++ {
// 			if companies[i].ID == companyid {
// 				companies[i].Likes++
// 			}
// 		}
// 		// return a pointer to the updated companies list
// 		c.JSON(http.StatusOK, &companies)
// 	} else {
// 		// Company ID is invalid
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
}
