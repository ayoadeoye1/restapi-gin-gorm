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