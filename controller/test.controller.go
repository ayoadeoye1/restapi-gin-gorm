package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

type GetUsersResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

// @Summary Show server running status
// @Description Get a message indicating that the Gin-Gonic server is up and running
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "status message"
// @Router / [get]
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Gin-Gonic Server is Up and Running ~~~~",
	})
}

// @Summary Get list of dummy users
// @Description Get list of dummy users from the endpoint
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "status message"
// @Router /api/v1/user [get]
func GetUser(c *gin.Context) {
	users := []gin.H{
		{
			"id":    1,
			"name":  "Alice Johnson",
			"email": "alice.johnson@example.com",
			"phone": "+1-202-555-0123",
			"address": gin.H{
				"street": "123 Maple St",
				"city":   "Springfield",
				"state":  "IL",
				"zip":    "62701",
			},
		},
		{
			"id":    2,
			"name":  "Bob Smith",
			"email": "bob.smith@example.com",
			"phone": "+1-202-555-0198",
			"address": gin.H{
				"street": "456 Oak St",
				"city":   "Lincoln",
				"state":  "NE",
				"zip":    "68502",
			},
		},
		{
			"id":    3,
			"name":  "Charlie Brown",
			"email": "charlie.brown@example.com",
			"phone": "+1-202-555-0145",
			"address": gin.H{
				"street": "789 Pine St",
				"city":   "Madison",
				"state":  "WI",
				"zip":    "53703",
			},
		},
		{
			"id":    4,
			"name":  "Diana Prince",
			"email": "diana.prince@example.com",
			"phone": "+1-202-555-0167",
			"address": gin.H{
				"street": "321 Birch St",
				"city":   "Seattle",
				"state":  "WA",
				"zip":    "98101",
			},
		},
		{
			"id":    5,
			"name":  "Edward Elric",
			"email": "edward.elric@example.com",
			"phone": "+1-202-555-0189",
			"address": gin.H{
				"street": "654 Cedar St",
				"city":   "Phoenix",
				"state":  "AZ",
				"zip":    "85001",
			},
		},
		{
			"id":    6,
			"name":  "Fiona Gallagher",
			"email": "fiona.gallagher@example.com",
			"phone": "+1-202-555-0178",
			"address": gin.H{
				"street": "987 Spruce St",
				"city":   "Chicago",
				"state":  "IL",
				"zip":    "60601",
			},
		},
		{
			"id":    7,
			"name":  "George Washington",
			"email": "george.washington@example.com",
			"phone": "+1-202-555-0112",
			"address": gin.H{
				"street": "135 Elm St",
				"city":   "Boston",
				"state":  "MA",
				"zip":    "02108",
			},
		},
		{
			"id":    8,
			"name":  "Hannah Baker",
			"email": "hannah.baker@example.com",
			"phone": "+1-202-555-0134",
			"address": gin.H{
				"street": "246 Willow St",
				"city":   "Austin",
				"state":  "TX",
				"zip":    "73301",
			},
		},
		{
			"id":    9,
			"name":  "Isaac Newton",
			"email": "isaac.newton@example.com",
			"phone": "+1-202-555-0156",
			"address": gin.H{
				"street": "357 Ash St",
				"city":   "Denver",
				"state":  "CO",
				"zip":    "80201",
			},
		},
		{
			"id":    10,
			"name":  "Jack Daniels",
			"email": "jack.daniels@example.com",
			"phone": "+1-202-555-0100",
			"address": gin.H{
				"street": "468 Chestnut St",
				"city":   "San Francisco",
				"state":  "CA",
				"zip":    "94101",
			},
		},
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "List of users retrieved successfully",
		"data":    users,
	}

	c.JSON(http.StatusOK, response)
}
