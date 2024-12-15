package middleware

// func authMiddleware() gin.HandlerFunc {
//  return func(c *gin.Context) {
//   tokenString := c.GetHeader("Authorization")

//   // Parse the token
//   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//     return nil, http.ErrAbortHandler
//    }
//    return []byte(secretKey), nil
//   })

//   if err != nil || !token.Valid {
//    c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//    c.Abort() // Stop further processing if unauthorized
//    return
//   }

//   // Set the token claims to the context
//   if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//    c.Set("claims", claims)
//   } else {
//    c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//    c.Abort()
//    return
//   }

//   c.Next() // Proceed to the next handler if authorized
//  }
// }

// func adminMiddleware() gin.HandlerFunc {
//  return func(c *gin.Context) {
//   claims := c.MustGet("claims").(jwt.MapClaims)
//   role := claims["role"].(string)

//   if role != "admin" {
//    c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
//    c.Abort()
//    return
//   }

//   c.Next()
//  }
// }

//ehhhrhrhehhe

{
    "success": true,
    "data": {
      "event": {
        "tournament": {
          "name": "UEFA Conference League",
          "slug": "uefa-conference-league",
          "category": {
            "name": "Europe",
            "slug": "europe",
            "sport": {
              "name": "Football",
              "slug": "football",
              "id": 1
            },
            "id": 1465,
            "country": {},
            "flag": "europe"
          },
          "uniqueTournament": {
            "name": "UEFA Conference League",
            "slug": "uefa-europa-conference-league",
            "secondaryColorHex": "#04be15",
            "category": {
              "name": "Europe",
              "slug": "europe",
              "sport": {
                "name": "Football",
                "slug": "football",
                "id": 1
              },
              "id": 1465,
              "country": {},
              "flag": "europe"
            },
            "userCount": 304496,
            "id": 17015,
            "country": {},
            "hasPerformanceGraphFeature": false,
            "hasEventPlayerStatistics": true,
            "displayInverseHomeAwayTeams": false
          },
          "priority": 579,
          "isGroup": false,
          "competitionType": 1,
          "id": 138316
        },
        "season": {
          "name": "UEFA Conference League 24/25",
          "year": "24/25",
          "editor": false,
          "id": 61648
        },
        "roundInfo": {
          "round": 5
        },
        "customId": "jKsiN",
        "status": {
          "code": 100,
          "description": "Ended",
          "type": "finished"
        },
        "winnerCode": 2,
        "venue": {
          "city": {
            "name": "Kópavogur"
          },
          "venueCoordinates": {
            "latitude": 64.104198,
            "longitude": -21.896615
          },
          "hidden": true,
          "slug": "k-pavogsv-llur",
          "name": "Kópavogsvöllur",
          "capacity": 5501,
          "id": 3976,
          "country": {
            "alpha2": "IS",
            "alpha3": "ISL",
            "name": "Iceland",
            "slug": "iceland"
          },
          "fieldTranslations": {
            "nameTranslation": {
              "ar": "كوپاڢوغس ڢولور"
            },
            "shortNameTranslation": {}
          },
          "stadium": {
            "name": "Kópavogsvöllur",
            "capacity": 5501
          }
        },
        "referee": {
          "name": "Luka Bilbija",
          "slug": "bilbija-luka",
          "yellowCards": 32,
          "redCards": 2,
          "yellowRedCards": 1,
          "games": 11,
          "sport": {
            "id": 1,
            "slug": "football",
            "name": "Football"
          },
          "id": 787817,
          "country": {
            "alpha2": "BA",
            "alpha3": "BIH",
            "name": "Bosnia & Herzegovina",
            "slug": "bosnia-and-herzegovina"
          }
        },
        "homeTeam": {
          "name": "Víkingur Reykjavík",
          "slug": "vikingur-reykjavik",
          "shortName": "Víkingur",
          "gender": "M",
          "sport": {
            "name": "Football",
            "slug": "football",
            "id": 1
          },
          "userCount": 27409,
          "manager": {
            "name": "Arnar Bergmann Gunnlaugsson",
            "slug": "arnar-bergmann-gunnlaugsson",
            "shortName": "A. B. Gunnlaugsson",
            "id": 790785,
            "country": {
              "alpha2": "IS",
              "alpha3": "ISL",
              "name": "Iceland",
              "slug": "iceland"
            }
          },
          "venue": {
            "city": {
              "name": "Reykjavik"
            },
            "venueCoordinates": {
              "latitude": 64.13548,
              "longitude": -21.89541
            },
            "hidden": true,
            "slug": "v-kingsvollur-stadium",
            "name": "Víkingsvollur Stadium",
            "capacity": 1500,
            "id": 16714,
            "country": {
              "alpha2": "IS",
              "alpha3": "ISL",
              "name": "Iceland",
              "slug": "iceland"
            },
            "fieldTranslations": {
              "nameTranslation": {
                "ar": "Víkingsvollur"
              },
              "shortNameTranslation": {}
            },
            "stadium": {
              "name": "Víkingsvollur Stadium",
              "capacity": 1500
            }
          },
          "nameCode": "VIK",
          "class": 4,
          "disabled": false,
          "national": false,
          "type": 0,
          "id": 1908,
          "country": {
            "alpha2": "IS",
            "alpha3": "ISL",
            "name": "Iceland",
            "slug": "iceland"
          },
          "subTeams": [],
          "fullName": "Víkingur Reykjavík",
          "teamColors": {
            "primary": "#ff0100",
            "secondary": "#000000",
            "text": "#000000"
          },
          "foundationDateTimestamp": -1947024000,
          "fieldTranslations": {
            "nameTranslation": {
              "ar": "فيكينجور ريكيافيك",
              "ru": "Викингур  Рейкьявик"
            },
            "shortNameTranslation": {}
          },
          "logo": "data:image/jpeg;base64,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"
        },
        "awayTeam": {
          "name": "Djurgårdens IF",
          "slug": "djurgardens-if",
          "shortName": "Djurgården",
          "gender": "M",
          "sport": {
            "name": "Football",
            "slug": "football",
            "id": 1
          },
          "userCount": 40559,
          "manager": {
            "name": "Roberth Bjorkensjo",
            "slug": "roberth-bjorkensjo",
            "shortName": "R. Bjorkensjo",
            "id": 139143,
            "country": {
              "alpha2": "SE",
              "alpha3": "SWE",
              "name": "Sweden",
              "slug": "sweden"
            }
          },
          "venue": {
            "city": {
              "name": "Stockholm"
            },
            "venueCoordinates": {
              "latitude": 59.291033,
              "longitude": 18.084543
            },
            "hidden": true,
            "slug": "tele2-arena",
            "name": "Tele2 Arena",
            "capacity": 32000,
            "id": 12373,
            "country": {
              "alpha2": "SE",
              "alpha3": "SWE",
              "name": "Sweden",
              "slug": "sweden"
            },
            "fieldTranslations": {
              "nameTranslation": {
                "ar": "Tele2 Arena"
              },
              "shortNameTranslation": {}
            },
            "stadium": {
              "name": "Tele2 Arena",
              "capacity": 32000
            }
          },
          "nameCode": "DIF",
          "class": 1,
          "disabled": false,
          "national": false,
          "type": 0,
          "id": 1759,
          "country": {
            "alpha2": "SE",
            "alpha3": "SWE",
            "name": "Sweden",
            "slug": "sweden"
          },
          "subTeams": [],
          "fullName": "Djurgårdens IF",
          "teamColors": {
            "primary": "#33ccff",
            "secondary": "#000099",
            "text": "#000099"
          },
          "foundationDateTimestamp": -2486937600,
          "fieldTranslations": {
            "nameTranslation": {
              "ar": "ديورجادنز",
              "ru": "ФК Юргорден"
            },
            "shortNameTranslation": {}
          },
          "logo": "data:image/jpeg;base64,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"
        },
        "homeScore": {
          "current": 1,
          "display": 1,
          "period1": 0,
          "period2": 1,
          "normaltime": 1
        },
        "awayScore": {
          "current": 2,
          "display": 2,
          "period1": 0,
          "period2": 2,
          "normaltime": 2
        },
        "coverage": -1,
        "time": {
          "injuryTime1": 2,
          "injuryTime2": 5,
          "currentPeriodStartTimestamp": 1734012294
        },
        "changes": {
          "changes": [
            "status.code",
            "status.description",
            "status.type"
          ],
          "changeTimestamp": 1734015357
        },
        "hasGlobalHighlights": false,
        "hasXg": true,
        "hasEventPlayerStatistics": true,
        "hasEventPlayerHeatMap": true,
        "detailId": 1,
        "crowdsourcingDataDisplayEnabled": false,
        "id": 12765132,
        "awayRedCards": 1,
        "defaultPeriodCount": 2,
        "defaultPeriodLength": 45,
        "defaultOvertimeLength": 15,
        "currentPeriodStartTimestamp": 1734012294,
        "startTimestamp": 1734008400,
        "slug": "vikingur-reykjavik-djurgardens-if",
        "finalResultOnly": false,
        "feedLocked": true,
        "fanRatingEvent": false,
        "seasonStatisticsType": "overall",
        "showTotoPromo": true,
        "isEditor": false
      }
    }
  }
