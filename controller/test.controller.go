package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Street string `json:"Street"`
	City   string `json:"City"`
	State  string `json:"State"`
	Zip    string `json:"Zip"`
}

type User struct {
	ID      int     `json:"ID"`
	Name    string  `json:"Name"`
	Email   string  `json:"Email"`
	Phone   string  `json:"Phone"`
	Address Address `json:"Address"`
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
// @Success 200 {object} GetUsersResponse "status message"
// @Router /api/v1/user [get]
func GetUser(c *gin.Context) {
	users := []User{
		{
			ID:    1,
			Name:  "Alice Johnson",
			Email: "alice.johnson@example.com",
			Phone: "+1-202-555-0123",
			Address: Address{
				Street: "123 Maple St",
				City:   "Springfield",
				State:  "IL",
				Zip:    "62701",
			},
		},
		{
			ID:    2,
			Name:  "Bob Smith",
			Email: "bob.smith@example.com",
			Phone: "+1-202-555-0198",
			Address: Address{
				Street: "456 Oak St",
				City:   "Lincoln",
				State:  "NE",
				Zip:    "68502",
			},
		},
		{
			ID:    3,
			Name:  "Charlie Brown",
			Email: "charlie.brown@example.com",
			Phone: "+1-202-555-0145",
			Address: Address{
				Street: "789 Pine St",
				City:   "Madison",
				State:  "WI",
				Zip:    "53703",
			},
		},
		{
			ID:    4,
			Name:  "Diana Prince",
			Email: "diana.prince@example.com",
			Phone: "+1-202-555-0167",
			Address: Address{
				Street: "321 Birch St",
				City:   "Seattle",
				State:  "WA",
				Zip:    "98101",
			},
		},
		{
			ID:    5,
			Name:  "Edward Elric",
			Email: "edward.elric@example.com",
			Phone: "+1-202-555-0189",
			Address: Address{
				Street: "654 Cedar St",
				City:   "Phoenix",
				State:  "AZ",
				Zip:    "85001",
			},
		},
		{
			ID:    6,
			Name:  "Fiona Gallagher",
			Email: "fiona.gallagher@example.com",
			Phone: "+1-202-555-0178",
			Address: Address{
				Street: "987 Spruce St",
				City:   "Chicago",
				State:  "IL",
				Zip:    "60601",
			},
		},
		{
			ID:    7,
			Name:  "George Washington",
			Email: "george.washington@example.com",
			Phone: "+1-202-555-0112",
			Address: Address{
				Street: "135 Elm St",
				City:   "Boston",
				State:  "MA",
				Zip:    "02108",
			},
		},
		{
			ID:    8,
			Name:  "Hannah Baker",
			Email: "hannah.baker@example.com",
			Phone: "+1-202-555-0134",
			Address: Address{
				Street: "246 Willow St",
				City:   "Austin",
				State:  "TX",
				Zip:    "73301",
			},
		},
		{
			ID:    9,
			Name:  "Isaac Newton",
			Email: "isaac.newton@example.com",
			Phone: "+1-202-555-0156",
			Address: Address{
				Street: "357 Ash St",
				City:   "Denver",
				State:  "CO",
				Zip:    "80201",
			},
		},
		{
			ID:    10,
			Name:  "Jack Daniels",
			Email: "jack.daniels@example.com",
			Phone: "+1-202-555-0100",
			Address: Address{
				Street: "468 Chestnut St",
				City:   "San Francisco",
				State:  "CA",
				Zip:    "94101",
			},
		},
	}

	response := GetUsersResponse{
		Status:  "success",
		Message: "List of users retrieved successfully",
		Data:    users,
	}

	c.JSON(http.StatusOK, response)
}
