package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func UserAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		token, err := c.Cookie("Authorization")
		if err != nil {
			helper.SendError(c, http.StatusBadRequest, "Authorization Token is missing in the header/cookie", err.Error())
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		tokenString = token
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		helper.SendError(c, http.StatusBadRequest, "Authorization Token coud not be verified", err.Error())
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	c.Next()
}

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

// export const getLeagueOverallTopTeam = async (
//     req: Request,
//     res: Response,
//     next: NextFunction,
// ) => {
//     try {
//         const { tournamentId } = req.params;
//         const { seasonId } = req.query;

//         const getLeagueOverallTopTeam = await allSportsFootball.getLeagueOverallTopTeam(
//             tournamentId,
//             String(seasonId),
//         );

//         if(!getLeagueOverallTopTeam.topTeams){
//             return res.status(StatusCodes.OK).json({
//                 success: true,
//                 data: getLeagueOverallTopTeam,
//             });
//         }

//         const modifiedTopTeams = await Promise.all(
//             Object.entries(getLeagueOverallTopTeam?.topTeams).map(async ([key, value]) => {
//                 if (Array.isArray(value)) {
//                     const modifiedArray = await Promise.all(
//                         value.map?.(async (ele: any) => {
//                             const teamLogo = await allSportsFootball.getTeamLogoImage(ele?.team?.id);
//                             ele.team.teamLogo = teamLogo || "";
//                             return ele;
//                         }) || []
//                     );
//                     return [key, modifiedArray];
//                 }
//                 return [key, value];
//             })
//         );

//         const updatedTopTeams = Object.fromEntries(modifiedTopTeams);
//         getLeagueOverallTopTeam.topTeams = updatedTopTeams

//         return res.status(StatusCodes.OK).json({
//             success: true,
//             data: getLeagueOverallTopTeam,
//         });
//     } catch (error) {
//         next(error);
//     }
// };
