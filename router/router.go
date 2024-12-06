package router

import (
	"github.com/ayoadeoye1/restapi-gin-gorm/controller"
	usercontroller "github.com/ayoadeoye1/restapi-gin-gorm/controller/user_controller"
	"github.com/gin-gonic/gin"

	//swagger files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userController *usercontroller.UserController) *gin.Engine {
	routes := gin.Default()

	// routes.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"*"}, github.com/gin-contrib/cors
	//     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//     AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	// }))

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.GET("/", controller.Home)

	router := routes.Group("/api/v1")
	userRouter := router.Group("/user")
	userRouter.POST("/", userController.Create)

	userRouter.GET("/", controller.GetUser)
	return routes
}

//Last Matches

// {
// 	"success": true,
// 	"data": {
// 	  "tournamentTeamEvents": {
// 		"1": {
// 		  "3": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 2,
// 				"period2": 2,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436947,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "everton-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dskb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 3,
// 				"period2": 1,
// 				"normaltime": 4
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436922,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "bournemouth-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsT",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 1,
// 				"period2": 3,
// 				"normaltime": 4
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436490,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "fulham-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436485,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "southampton-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsh",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437060,
// 			  "startTimestamp": 1730568600,
// 			  "slug": "crystal-palace-wolverhampton",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "7": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hsH",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436961,
// 			  "startTimestamp": 1733254200,
// 			  "slug": "ipswich-town-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hO",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436502,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "newcastle-united-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hP",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 2,
// 				"period2": 0,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436493,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "aston-villa-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hsT",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436483,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "fulham-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsh",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437060,
// 			  "startTimestamp": 1730568600,
// 			  "slug": "crystal-palace-wolverhampton",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "14": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osr",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 2,
// 				"period2": 1,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436972,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "manchester-city-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osH",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436503,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "ipswich-town-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436489,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "arsenal-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osO",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 0,
// 				"period2": 3,
// 				"normaltime": 3
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436487,
// 			  "startTimestamp": 1731247200,
// 			  "slug": "newcastle-united-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osM",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437057,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "west-ham-united-nottingham-forest",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "17": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osr",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 2,
// 				"period2": 1,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436972,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "manchester-city-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436504,
// 			  "startTimestamp": 1733068800,
// 			  "slug": "liverpool-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rI",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 2,
// 				"period2": 2,
// 				"normaltime": 4
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436499,
// 			  "startTimestamp": 1732383000,
// 			  "slug": "tottenham-hotspur-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rsF",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436882,
// 			  "startTimestamp": 1731173400,
// 			  "slug": "brighton-and-hove-albion-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rkb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437051,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "bournemouth-manchester-city",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "30": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "FV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436501,
// 			  "startTimestamp": 1732910400,
// 			  "slug": "southampton-brighton-and-hove-albion",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Fskb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436491,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "bournemouth-brighton-and-hove-albion",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rsF",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436882,
// 			  "startTimestamp": 1731173400,
// 			  "slug": "brighton-and-hove-albion-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "FsU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437054,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "liverpool-brighton-and-hove-albion",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsF",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437044,
// 			  "startTimestamp": 1729951200,
// 			  "slug": "brighton-and-hove-albion-wolverhampton",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "31": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GsM",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436966,
// 			  "startTimestamp": 1733256900,
// 			  "slug": "west-ham-united-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Gsab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 3,
// 				"period2": 1,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436500,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "brentford-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GsN",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436495,
// 			  "startTimestamp": 1732365000,
// 			  "slug": "chelsea-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GK",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 2,
// 				"period2": 1,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436494,
// 			  "startTimestamp": 1731247200,
// 			  "slug": "manchester-united-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GsH",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437053,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "ipswich-town-leicester-city",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "32": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hsH",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436961,
// 			  "startTimestamp": 1733254200,
// 			  "slug": "ipswich-town-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osH",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436503,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "ipswich-town-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "HsK",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436497,
// 			  "startTimestamp": 1732465800,
// 			  "slug": "manchester-united-ipswich-town",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "HsI",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 2,
// 				"period2": 0,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436492,
// 			  "startTimestamp": 1731247200,
// 			  "slug": "tottenham-hotspur-ipswich-town",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GsH",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437053,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "ipswich-town-leicester-city",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "33": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "IsT",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436507,
// 			  "startTimestamp": 1733059800,
// 			  "slug": "fulham-tottenham-hotspur",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rI",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 2,
// 				"period2": 2,
// 				"normaltime": 4
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436499,
// 			  "startTimestamp": 1732383000,
// 			  "slug": "tottenham-hotspur-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "HsI",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 2,
// 				"period2": 0,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436492,
// 			  "startTimestamp": 1731247200,
// 			  "slug": "tottenham-hotspur-ipswich-town",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "IP",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 0,
// 				"period2": 4,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437059,
// 			  "startTimestamp": 1730642400,
// 			  "slug": "aston-villa-tottenham-hotspur",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hI",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437046,
// 			  "startTimestamp": 1730037600,
// 			  "slug": "tottenham-hotspur-crystal-palace",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "35": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "KR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436935,
// 			  "startTimestamp": 1733343300,
// 			  "slug": "arsenal-manchester-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "KsY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 2,
// 				"period2": 2,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436505,
// 			  "startTimestamp": 1733059800,
// 			  "slug": "everton-manchester-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "HsK",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Ipswich Town",
// 				"slug": "ipswich-town",
// 				"shortName": "Ipswich",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 132520,
// 				"nameCode": "IPS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 32,
// 				"teamColors": {
// 				  "primary": "#0000ff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ايبسويتش تاون",
// 					"ru": "Ипсвич Таун"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436497,
// 			  "startTimestamp": 1732465800,
// 			  "slug": "manchester-united-ipswich-town",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GK",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 2,
// 				"period2": 1,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436494,
// 			  "startTimestamp": 1731247200,
// 			  "slug": "manchester-united-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "mcYb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437055,
// 			  "startTimestamp": 1730651400,
// 			  "slug": "chelsea-manchester-united",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "37": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GsM",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436966,
// 			  "startTimestamp": 1733256900,
// 			  "slug": "west-ham-united-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "MR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 2,
// 				"period2": 0,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 5,
// 				"display": 5,
// 				"period1": 5,
// 				"period2": 0,
// 				"normaltime": 5
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436914,
// 			  "startTimestamp": 1732987800,
// 			  "slug": "arsenal-west-ham-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "MO",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436496,
// 			  "startTimestamp": 1732564800,
// 			  "slug": "newcastle-united-west-ham-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "MY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436486,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "everton-west-ham-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osM",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437057,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "west-ham-united-nottingham-forest",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "38": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "NV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 5,
// 				"display": 5,
// 				"period1": 3,
// 				"period2": 2,
// 				"normaltime": 5
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436984,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "southampton-chelsea",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "NP",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 2,
// 				"period2": 1,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436498,
// 			  "startTimestamp": 1733059800,
// 			  "slug": "aston-villa-chelsea",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "GsN",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436495,
// 			  "startTimestamp": 1732365000,
// 			  "slug": "chelsea-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "NR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436890,
// 			  "startTimestamp": 1731256200,
// 			  "slug": "arsenal-chelsea",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "mcYb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437055,
// 			  "startTimestamp": 1730651400,
// 			  "slug": "chelsea-manchester-united",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "39": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "OU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 0,
// 				"period2": 3,
// 				"normaltime": 3
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436978,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "liverpool-newcastle-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hO",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436502,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "newcastle-united-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "MO",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436496,
// 			  "startTimestamp": 1732564800,
// 			  "slug": "newcastle-united-west-ham-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osO",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 0,
// 				"period2": 3,
// 				"normaltime": 3
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436487,
// 			  "startTimestamp": 1731247200,
// 			  "slug": "newcastle-united-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Ardc",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437056,
// 			  "startTimestamp": 1730550600,
// 			  "slug": "arsenal-newcastle-united",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "40": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Psab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 3,
// 				"period2": 0,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436942,
// 			  "startTimestamp": 1733343300,
// 			  "slug": "brentford-aston-villa",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "NP",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 2,
// 				"period2": 1,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436498,
// 			  "startTimestamp": 1733059800,
// 			  "slug": "aston-villa-chelsea",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hP",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 2,
// 				"period2": 0,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436493,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "aston-villa-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "PU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436484,
// 			  "startTimestamp": 1731182400,
// 			  "slug": "liverpool-aston-villa",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "IP",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 0,
// 				"period2": 4,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437059,
// 			  "startTimestamp": 1730642400,
// 			  "slug": "aston-villa-tottenham-hotspur",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "42": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "KR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436935,
// 			  "startTimestamp": 1733343300,
// 			  "slug": "arsenal-manchester-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "MR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 2,
// 				"period2": 0,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 5,
// 				"display": 5,
// 				"period1": 5,
// 				"period2": 0,
// 				"normaltime": 5
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436914,
// 			  "startTimestamp": 1732987800,
// 			  "slug": "arsenal-west-ham-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "osR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Nottingham Forest",
// 				"slug": "nottingham-forest",
// 				"shortName": "Forest",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 289871,
// 				"nameCode": "NFO",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 14,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نوتنجهام فوريست",
// 					"ru": "Ноттингем Форест"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436489,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "arsenal-nottingham-forest",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "NR",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436890,
// 			  "startTimestamp": 1731256200,
// 			  "slug": "arsenal-chelsea",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Ardc",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Arsenal",
// 				"slug": "arsenal",
// 				"shortName": "Arsenal",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2210565,
// 				"nameCode": "ARS",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 42,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أرسنال",
// 					"ru": "Арсенал"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437056,
// 			  "startTimestamp": 1730550600,
// 			  "slug": "arsenal-newcastle-united",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "43": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "IsT",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Tottenham Hotspur",
// 				"slug": "tottenham-hotspur",
// 				"shortName": "Tottenham ",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1248951,
// 				"nameCode": "TOT",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 33,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000066",
// 				  "text": "#000066"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "توتنهام هوتسبير",
// 					"ru": "Тоттенхэм Хотспур"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436507,
// 			  "startTimestamp": 1733059800,
// 			  "slug": "fulham-tottenham-hotspur",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsT",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 1,
// 				"period2": 3,
// 				"normaltime": 4
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436490,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "fulham-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "hsT",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Crystal Palace",
// 				"slug": "crystal-palace",
// 				"shortName": "Crystal Palace",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 377993,
// 				"nameCode": "CRY",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 7,
// 				"teamColors": {
// 				  "primary": "#0033ff",
// 				  "secondary": "#b90d2b",
// 				  "text": "#b90d2b"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "كريستال بالاس",
// 					"ru": "Кристал Пэлас"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436483,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "fulham-crystal-palace",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Tsab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437052,
// 			  "startTimestamp": 1730750400,
// 			  "slug": "brentford-fulham",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "TsY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437047,
// 			  "startTimestamp": 1729960200,
// 			  "slug": "everton-fulham",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "44": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "OU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Newcastle United",
// 				"slug": "newcastle-united",
// 				"shortName": "Newcastle",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 848307,
// 				"nameCode": "NEW",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 39,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "نيوكاسل يونايتد",
// 					"ru": "Ньюкасл Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 0,
// 				"period2": 3,
// 				"normaltime": 3
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436978,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "liverpool-newcastle-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436504,
// 			  "startTimestamp": 1733068800,
// 			  "slug": "liverpool-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "UsV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436506,
// 			  "startTimestamp": 1732456800,
// 			  "slug": "southampton-liverpool",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "PU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436484,
// 			  "startTimestamp": 1731182400,
// 			  "slug": "liverpool-aston-villa",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "FsU",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437054,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "liverpool-brighton-and-hove-albion",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "45": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "NV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Chelsea",
// 				"slug": "chelsea",
// 				"shortName": "Chelsea",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 1996984,
// 				"nameCode": "CHE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 38,
// 				"teamColors": {
// 				  "primary": "#0310a7",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "تشيلسي",
// 					"ru": "Челси"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 5,
// 				"display": 5,
// 				"period1": 3,
// 				"period2": 2,
// 				"normaltime": 5
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436984,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "southampton-chelsea",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "FV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436501,
// 			  "startTimestamp": 1732910400,
// 			  "slug": "southampton-brighton-and-hove-albion",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "UsV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Liverpool",
// 				"slug": "liverpool",
// 				"shortName": "Liverpool",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2416877,
// 				"nameCode": "LIV",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 44,
// 				"teamColors": {
// 				  "primary": "#cc0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليفربول",
// 					"ru": "Ливерпуль"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436506,
// 			  "startTimestamp": 1732456800,
// 			  "slug": "southampton-liverpool",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsV",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436485,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "southampton-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "VY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437058,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "everton-southampton",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "48": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dsY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 2,
// 				"period2": 2,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436947,
// 			  "startTimestamp": 1733340600,
// 			  "slug": "everton-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "KsY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Manchester United",
// 				"slug": "manchester-united",
// 				"shortName": "Man Utd",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2454572,
// 				"nameCode": "MUN",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 35,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#373737",
// 				  "text": "#373737"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانتشستر يونايتد",
// 					"ru": "Манчестер Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 2,
// 				"period2": 2,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436505,
// 			  "startTimestamp": 1733059800,
// 			  "slug": "everton-manchester-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Ysab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436488,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "brentford-everton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "MY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "West Ham United",
// 				"slug": "west-ham-united",
// 				"shortName": "West Ham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 649571,
// 				"nameCode": "WHU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 37,
// 				"teamColors": {
// 				  "primary": "#66192c",
// 				  "secondary": "#59b3e4",
// 				  "text": "#59b3e4"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وست هام يونايتد",
// 					"ru": "Вест Хэм Юнайтед"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436486,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "everton-west-ham-united",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "VY",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Southampton",
// 				"slug": "southampton",
// 				"shortName": "Southampton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 197127,
// 				"nameCode": "SOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 45,
// 				"teamColors": {
// 				  "primary": "#ff0000",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ساوثهامتون",
// 					"ru": "Саутгемптон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437058,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "everton-southampton",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "50": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Psab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 3,
// 				"period2": 0,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436942,
// 			  "startTimestamp": 1733343300,
// 			  "slug": "brentford-aston-villa",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Gsab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Leicester City",
// 				"slug": "leicester-city",
// 				"shortName": "Leicester",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 506527,
// 				"nameCode": "LEI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 31,
// 				"teamColors": {
// 				  "primary": "#003090",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "ليستر سيتي",
// 					"ru": "Лестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 3,
// 				"period2": 1,
// 				"normaltime": 4
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436500,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "brentford-leicester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Ysab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Everton",
// 				"slug": "everton",
// 				"shortName": "Everton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 432074,
// 				"nameCode": "EVE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 48,
// 				"teamColors": {
// 				  "primary": "#274488",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "إيفرتون",
// 					"ru": "Эвертон"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "awayScore": {
// 				"current": 0,
// 				"display": 0,
// 				"period1": 0,
// 				"period2": 0,
// 				"normaltime": 0
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436488,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "brentford-everton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "abskb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436876,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "bournemouth-brentford",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Tsab",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Fulham",
// 				"slug": "fulham",
// 				"shortName": "Fulham",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 375454,
// 				"nameCode": "FUL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 43,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "فولهام",
// 					"ru": "Фулхэм"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 0,
// 				"period2": 2,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 1,
// 				"period2": 0,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437052,
// 			  "startTimestamp": 1730750400,
// 			  "slug": "brentford-fulham",
// 			  "finalResultOnly": false
// 			}
// 		  ],
// 		  "60": [
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "dskb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Wolverhampton",
// 				"slug": "wolverhampton",
// 				"shortName": "Wolves",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 390904,
// 				"nameCode": "WOL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 3,
// 				"teamColors": {
// 				  "primary": "#ff9900",
// 				  "secondary": "#000000",
// 				  "text": "#000000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "وولفرهامبتون واندررز",
// 					"ru": "Вулверхэмптон Уондерерс"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 4,
// 				"display": 4,
// 				"period1": 3,
// 				"period2": 1,
// 				"normaltime": 4
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436922,
// 			  "startTimestamp": 1732978800,
// 			  "slug": "bournemouth-wolverhampton",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Fskb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 2,
// 			  "homeTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Brighton & Hove Albion",
// 				"slug": "brighton-and-hove-albion",
// 				"shortName": "Brighton",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 596880,
// 				"nameCode": "BHA",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 30,
// 				"teamColors": {
// 				  "primary": "#0054a6",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برايتون أند هوف ألبيون",
// 					"ru": "Брайтон энд Хоув Альбион"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436491,
// 			  "startTimestamp": 1732374000,
// 			  "slug": "bournemouth-brighton-and-hove-albion",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "abskb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Brentford",
// 				"slug": "brentford",
// 				"shortName": "Brentford",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 340288,
// 				"nameCode": "BRE",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 50,
// 				"teamColors": {
// 				  "primary": "#ffffff",
// 				  "secondary": "#ff0000",
// 				  "text": "#ff0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "برينتفورد",
// 					"ru": "Брентфорд"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 3,
// 				"display": 3,
// 				"period1": 1,
// 				"period2": 2,
// 				"normaltime": 3
// 			  },
// 			  "awayScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "hasXg": true,
// 			  "id": 12436876,
// 			  "startTimestamp": 1731164400,
// 			  "slug": "bournemouth-brentford",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "rkb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 1,
// 			  "homeTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Manchester City",
// 				"slug": "manchester-city",
// 				"shortName": "Man City",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 2766537,
// 				"nameCode": "MCI",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 17,
// 				"teamColors": {
// 				  "primary": "#66ccff",
// 				  "secondary": "#ffffff",
// 				  "text": "#ffffff"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "مانشستر سيتي",
// 					"ru": "Манчестер Сити"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 2,
// 				"display": 2,
// 				"period1": 1,
// 				"period2": 1,
// 				"normaltime": 2
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437051,
// 			  "startTimestamp": 1730559600,
// 			  "slug": "bournemouth-manchester-city",
// 			  "finalResultOnly": false
// 			},
// 			{
// 			  "tournament": {
// 				"name": "Premier League",
// 				"slug": "premier-league",
// 				"category": {
// 				  "name": "England",
// 				  "slug": "england",
// 				  "sport": {
// 					"name": "Football",
// 					"slug": "football",
// 					"id": 1
// 				  },
// 				  "id": 1,
// 				  "flag": "england",
// 				  "alpha2": "EN"
// 				},
// 				"uniqueTournament": {
// 				  "name": "Premier League",
// 				  "slug": "premier-league",
// 				  "primaryColorHex": "#3c1c5a",
// 				  "secondaryColorHex": "#f80158",
// 				  "category": {
// 					"name": "England",
// 					"slug": "england",
// 					"sport": {
// 					  "name": "Football",
// 					  "slug": "football",
// 					  "id": 1
// 					},
// 					"id": 1,
// 					"flag": "england",
// 					"alpha2": "EN"
// 				  },
// 				  "userCount": 1392631,
// 				  "id": 17,
// 				  "displayInverseHomeAwayTeams": false
// 				},
// 				"priority": 617,
// 				"isLive": false,
// 				"id": 1
// 			  },
// 			  "customId": "Pkb",
// 			  "status": {
// 				"code": 100,
// 				"description": "Ended",
// 				"type": "finished"
// 			  },
// 			  "winnerCode": 3,
// 			  "homeTeam": {
// 				"name": "Aston Villa",
// 				"slug": "aston-villa",
// 				"shortName": "Aston Villa",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 817661,
// 				"nameCode": "AVL",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 40,
// 				"teamColors": {
// 				  "primary": "#670e36",
// 				  "secondary": "#94bee5",
// 				  "text": "#94bee5"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "أستون فيلا",
// 					"ru": "Астон Вилла"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "awayTeam": {
// 				"name": "Bournemouth",
// 				"slug": "bournemouth",
// 				"shortName": "Bournemouth",
// 				"gender": "M",
// 				"sport": {
// 				  "name": "Football",
// 				  "slug": "football",
// 				  "id": 1
// 				},
// 				"userCount": 266534,
// 				"nameCode": "BOU",
// 				"disabled": false,
// 				"national": false,
// 				"type": 0,
// 				"id": 60,
// 				"teamColors": {
// 				  "primary": "#000000",
// 				  "secondary": "#cc0000",
// 				  "text": "#cc0000"
// 				},
// 				"fieldTranslations": {
// 				  "nameTranslation": {
// 					"ar": "بورنموث",
// 					"ru": "Борнмут"
// 				  },
// 				  "shortNameTranslation": {}
// 				}
// 			  },
// 			  "homeScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "awayScore": {
// 				"current": 1,
// 				"display": 1,
// 				"period1": 0,
// 				"period2": 1,
// 				"normaltime": 1
// 			  },
// 			  "hasXg": true,
// 			  "id": 12437041,
// 			  "startTimestamp": 1729951200,
// 			  "slug": "bournemouth-aston-villa",
// 			  "finalResultOnly": false
// 			}
// 		  ]
// 		}
// 	  }
//   }
// 	}
