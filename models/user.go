package models

import (
	"time"
)

type Users struct {
	ID         int     `gorm:"type:int;primaryKey;autoIncrement"`
	FirstName  string  `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName   string  `gorm:"type:varchar(255);not null" json:"last_name"`
	Email      string  `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password   string  `gorm:"type:varchar(255);not null" json:"password"`
	Occupation *string `gorm:"type:varchar(255)" json:"occupation"`
	Address    *string `gorm:"type:varchar(255)" json:"address"`

	//time stamps
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	//associations
	Otp *Otp `gorm:"foreignKey:UserID"`
}

// {
//   "success": true,
//   "data": {
//     "standings": [
//       {
//         "tournament": {
//           "name": "U21 Premier League International Cup, Group A",
//           "slug": "u21-premier-league-international-cup-group-a",
//           "category": {
//             "name": "World",
//             "slug": "world",
//             "sport": {
//               "name": "Football",
//               "slug": "football",
//               "id": 1
//             },
//             "id": 1468,
//             "flag": "international"
//           },
//           "uniqueTournament": {
//             "name": "Premier League International Cup",
//             "slug": "premier-league-international-cup",
//             "category": {
//               "name": "World",
//               "slug": "world",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "id": 1468,
//               "flag": "international"
//             },
//             "userCount": 1576,
//             "hasPerformanceGraphFeature": false,
//             "id": 20048,
//             "displayInverseHomeAwayTeams": false
//           },
//           "priority": 0,
//           "isGroup": true,
//           "groupName": "Group A",
//           "isLive": false,
//           "id": 119905
//         },
//         "type": "total",
//         "name": "U21 Premier League International Cup 24/25, Group A",
//         "descriptions": [],
//         "rows": [
//           {
//             "team": {
//               "name": "Nottingham Forest U21",
//               "slug": "nottingham-forest-u21",
//               "shortName": "Forest U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 3695,
//               "nameCode": "NOF",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 324738,
//               "teamColors": {
//                 "primary": "#ff0000",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 1,
//             "matches": 3,
//             "wins": 2,
//             "scoresFor": 4,
//             "scoresAgainst": 1,
//             "id": 1229152,
//             "losses": 1,
//             "draws": 0,
//             "points": 6,
//             "scoreDiffFormatted": "+3"
//           },
//           {
//             "team": {
//               "name": "Valencia Mestalla",
//               "slug": "valencia-b",
//               "shortName": "Valencia B ",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 1770,
//               "nameCode": "VAM",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 24337,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#000000",
//                 "text": "#000000"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "КФ Валенсия Б"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 2,
//             "matches": 2,
//             "wins": 2,
//             "scoresFor": 4,
//             "scoresAgainst": 1,
//             "id": 1229155,
//             "losses": 0,
//             "draws": 0,
//             "points": 6,
//             "scoreDiffFormatted": "+3"
//           },
//           {
//             "team": {
//               "name": "AS Monaco U21",
//               "slug": "as-monaco-u21",
//               "shortName": "AS Monaco U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 649,
//               "nameCode": "MON",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 293450,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#ff0000",
//                 "text": "#ff0000"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ar": "AS Monaco تحت 21"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 3,
//             "matches": 2,
//             "wins": 1,
//             "scoresFor": 5,
//             "scoresAgainst": 2,
//             "id": 1229150,
//             "losses": 0,
//             "draws": 1,
//             "points": 4,
//             "scoreDiffFormatted": "+3"
//           },
//           {
//             "team": {
//               "name": "Reading U21",
//               "slug": "reading-u21",
//               "shortName": "Reading U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2175,
//               "nameCode": "REA",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 77867,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#003399",
//                 "text": "#003399"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "ФК Рединг (дубль)"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 4,
//             "matches": 2,
//             "wins": 1,
//             "scoresFor": 3,
//             "scoresAgainst": 5,
//             "id": 1229153,
//             "losses": 1,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "-2"
//           },
//           {
//             "team": {
//               "name": "Wolverhampton U21",
//               "slug": "wolverhampton-u21",
//               "shortName": "Wolves U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 4362,
//               "nameCode": "WOL",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36551,
//               "teamColors": {
//                 "primary": "#ff9900",
//                 "secondary": "#000000",
//                 "text": "#000000"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Вулверхэмптон  Уондерерс"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 5,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 4,
//             "scoresAgainst": 4,
//             "id": 1229156,
//             "losses": 0,
//             "draws": 2,
//             "points": 2,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "GNK Dinamo Zagreb U21",
//               "slug": "gnk-dinamo-zagreb-u21",
//               "shortName": "GNK Dinamo Zagreb U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 521,
//               "nameCode": "GNK",
//               "national": false,
//               "type": 0,
//               "id": 318029,
//               "teamColors": {
//                 "primary": "#0b59a1",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               }
//             },
//             "descriptions": [],
//             "position": 6,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 3,
//             "scoresAgainst": 6,
//             "id": 1229151,
//             "losses": 1,
//             "draws": 1,
//             "points": 1,
//             "scoreDiffFormatted": "-3"
//           },
//           {
//             "team": {
//               "name": "Jong Ajax",
//               "slug": "jong-ajax",
//               "shortName": "Jong Ajax",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 20517,
//               "nameCode": "AJA",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 46160,
//               "teamColors": {
//                 "primary": "#bb002b",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               }
//             },
//             "descriptions": [],
//             "position": 7,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 1,
//             "scoresAgainst": 3,
//             "id": 1229149,
//             "losses": 2,
//             "draws": 0,
//             "points": 0,
//             "scoreDiffFormatted": "-2"
//           },
//           {
//             "team": {
//               "name": "Tottenham U21",
//               "slug": "tottenham-u21",
//               "shortName": "Tottenham U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 5712,
//               "nameCode": "TOT",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 77865,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#000066",
//                 "text": "#000066"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Тоттенхэм"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 8,
//             "matches": 1,
//             "wins": 0,
//             "scoresFor": 1,
//             "scoresAgainst": 3,
//             "id": 1229154,
//             "losses": 1,
//             "draws": 0,
//             "points": 0,
//             "scoreDiffFormatted": "-2"
//           }
//         ],
//         "id": 134093,
//         "updatedAtTimestamp": 1724277156
//       },
//       {
//         "tournament": {
//           "name": "U21 Premier League International Cup, Group B",
//           "slug": "u21-premier-league-international-cup-group-b",
//           "category": {
//             "name": "World",
//             "slug": "world",
//             "sport": {
//               "name": "Football",
//               "slug": "football",
//               "id": 1
//             },
//             "id": 1468,
//             "flag": "international"
//           },
//           "uniqueTournament": {
//             "name": "Premier League International Cup",
//             "slug": "premier-league-international-cup",
//             "category": {
//               "name": "World",
//               "slug": "world",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "id": 1468,
//               "flag": "international"
//             },
//             "userCount": 1576,
//             "hasPerformanceGraphFeature": false,
//             "id": 20048,
//             "displayInverseHomeAwayTeams": false
//           },
//           "priority": 0,
//           "isGroup": true,
//           "groupName": "Group B",
//           "isLive": false,
//           "id": 119904
//         },
//         "type": "total",
//         "name": "U21 Premier League International Cup 24/25, Group B",
//         "descriptions": [],
//         "rows": [
//           {
//             "team": {
//               "name": "Olympique Lyonnais 2",
//               "slug": "olympique-lyonnais-2",
//               "shortName": "Olympique Lyonnais 2 U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 753,
//               "nameCode": "LYO",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 305409,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#004080",
//                 "text": "#004080"
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 1,
//             "matches": 4,
//             "wins": 3,
//             "scoresFor": 5,
//             "scoresAgainst": 2,
//             "id": 1229162,
//             "losses": 1,
//             "draws": 0,
//             "points": 9,
//             "scoreDiffFormatted": "+3"
//           },
//           {
//             "team": {
//               "name": "Blackburn U21",
//               "slug": "blackburn-rovers-u21",
//               "shortName": "Blackburn U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2462,
//               "nameCode": "BRO",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36543,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#0000ff",
//                 "text": "#0000ff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ar": "بلاكبرن روفرز الإحتياط تحت 21",
//                   "ru": "Блэкберн Роверс U21"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 2,
//             "matches": 3,
//             "wins": 1,
//             "scoresFor": 4,
//             "scoresAgainst": 4,
//             "id": 1229159,
//             "losses": 1,
//             "draws": 1,
//             "points": 4,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "Sunderland U21",
//               "slug": "sunderland-u21",
//               "shortName": "Sunderland U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2839,
//               "nameCode": "SUN",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36546,
//               "teamColors": {
//                 "primary": "#ca0000",
//                 "secondary": "#000000",
//                 "text": "#000000"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ar": "سندرلاند تحت 21",
//                   "ru": "Сандерленд U21"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 3,
//             "matches": 3,
//             "wins": 1,
//             "scoresFor": 4,
//             "scoresAgainst": 5,
//             "id": 1229163,
//             "losses": 1,
//             "draws": 1,
//             "points": 4,
//             "scoreDiffFormatted": "-1"
//           },
//           {
//             "team": {
//               "name": "Borussia Mönchengladbach II U23",
//               "slug": "borussia-monchengladbach-ii-u23",
//               "shortName": "Borussia M'gladbach II U23",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2107,
//               "nameCode": "BMG",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 5811,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#000000",
//                 "text": "#000000"
//               }
//             },
//             "descriptions": [],
//             "position": 4,
//             "matches": 3,
//             "wins": 1,
//             "scoresFor": 6,
//             "scoresAgainst": 6,
//             "id": 1229160,
//             "losses": 2,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "Middlesbrough U21",
//               "slug": "middlesbrough-u21",
//               "shortName": "Middlesbrough U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 1790,
//               "nameCode": "MIB",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 37483,
//               "teamColors": {
//                 "primary": "#d90000",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Мидлсбро"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 5,
//             "matches": 2,
//             "wins": 1,
//             "scoresFor": 3,
//             "scoresAgainst": 3,
//             "id": 1229161,
//             "losses": 1,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "Athletic Club B",
//               "slug": "athletic-bilbao-b",
//               "shortName": "Athletic Club B",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 3954,
//               "nameCode": "ACB",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 24324,
//               "teamColors": {
//                 "primary": "#aa0000",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ar": "اتلتيكو بلباو",
//                   "ru": "Атлетик Бильбао Б"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 6,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 4,
//             "scoresAgainst": 4,
//             "id": 1229157,
//             "losses": 0,
//             "draws": 2,
//             "points": 2,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "Benfica B",
//               "slug": "benfica-b",
//               "shortName": "Benfica B",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 16477,
//               "nameCode": "SLB",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 74465,
//               "teamColors": {
//                 "primary": "#cc0000",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ar": "بنفيكا B",
//                   "ru": "Бенфика Лиссабон Б"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 7,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 3,
//             "scoresAgainst": 3,
//             "id": 1229158,
//             "losses": 0,
//             "draws": 2,
//             "points": 2,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "West Ham U21",
//               "slug": "west-ham-u21",
//               "shortName": "West Ham U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 4535,
//               "nameCode": "WES",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36552,
//               "teamColors": {
//                 "primary": "#66192c",
//                 "secondary": "#59b3e4",
//                 "text": "#59b3e4"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Вест Хэм Юнайтед"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 8,
//             "matches": 3,
//             "wins": 0,
//             "scoresFor": 4,
//             "scoresAgainst": 6,
//             "id": 1229164,
//             "losses": 1,
//             "draws": 2,
//             "points": 2,
//             "scoreDiffFormatted": "-2"
//           }
//         ],
//         "id": 134094,
//         "updatedAtTimestamp": 1724277156
//       },
//       {
//         "tournament": {
//           "name": "U21 Premier League International Cup, Group C",
//           "slug": "u21-premier-league-international-cup-group-c",
//           "category": {
//             "name": "World",
//             "slug": "world",
//             "sport": {
//               "name": "Football",
//               "slug": "football",
//               "id": 1
//             },
//             "id": 1468,
//             "flag": "international"
//           },
//           "uniqueTournament": {
//             "name": "Premier League International Cup",
//             "slug": "premier-league-international-cup",
//             "category": {
//               "name": "World",
//               "slug": "world",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "id": 1468,
//               "flag": "international"
//             },
//             "userCount": 1576,
//             "hasPerformanceGraphFeature": false,
//             "id": 20048,
//             "displayInverseHomeAwayTeams": false
//           },
//           "priority": 0,
//           "isGroup": true,
//           "groupName": "Group C",
//           "isLive": false,
//           "id": 119906
//         },
//         "type": "total",
//         "name": "U21 Premier League International Cup 24/25, Group C",
//         "descriptions": [],
//         "rows": [
//           {
//             "team": {
//               "name": "Jong PSV Eindhoven",
//               "slug": "jong-psv-eindhoven",
//               "shortName": "Jong PSV",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 14303,
//               "nameCode": "PSV",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 46168,
//               "teamColors": {
//                 "primary": "#ff0000",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "ПСВ Эйндховен"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 1,
//             "matches": 3,
//             "wins": 1,
//             "scoresFor": 9,
//             "scoresAgainst": 5,
//             "id": 1292893,
//             "losses": 0,
//             "draws": 2,
//             "points": 5,
//             "scoreDiffFormatted": "+4"
//           },
//           {
//             "team": {
//               "name": "Fulham U21",
//               "slug": "fulham-u21",
//               "shortName": "Fulham U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 5140,
//               "nameCode": "FUL",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36548,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#000000",
//                 "text": "#000000"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Фулхэм"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 2,
//             "matches": 2,
//             "wins": 1,
//             "scoresFor": 7,
//             "scoresAgainst": 5,
//             "id": 1229167,
//             "losses": 0,
//             "draws": 1,
//             "points": 4,
//             "scoreDiffFormatted": "+2"
//           },
//           {
//             "team": {
//               "name": "Hertha BSC II",
//               "slug": "hertha-bsc-ii",
//               "shortName": "Hertha II ",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2091,
//               "nameCode": "BSC",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 2626,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#0131ad",
//                 "text": "#0131ad"
//               }
//             },
//             "descriptions": [],
//             "position": 3,
//             "matches": 2,
//             "wins": 1,
//             "scoresFor": 5,
//             "scoresAgainst": 4,
//             "id": 1229168,
//             "losses": 0,
//             "draws": 1,
//             "points": 4,
//             "scoreDiffFormatted": "+1"
//           },
//           {
//             "team": {
//               "name": "FC Nordsjælland Reserve",
//               "slug": "fc-nordsjaelland-reserve",
//               "shortName": "FC Nordsjælland",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 257,
//               "nameCode": "FCN",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 246759,
//               "teamColors": {
//                 "primary": "#374df5",
//                 "secondary": "#374df5",
//                 "text": "#ffffff"
//               }
//             },
//             "descriptions": [],
//             "position": 4,
//             "matches": 1,
//             "wins": 1,
//             "scoresFor": 2,
//             "scoresAgainst": 0,
//             "id": 1229166,
//             "losses": 0,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "+2"
//           },
//           {
//             "team": {
//               "name": "Norwich City U21",
//               "slug": "norwich-city-u21",
//               "shortName": "Norwich U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 1987,
//               "nameCode": "NOR",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 54367,
//               "teamColors": {
//                 "primary": "#336600",
//                 "secondary": "#ffcc33",
//                 "text": "#ffcc33"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Норвич Сити"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 5,
//             "matches": 2,
//             "wins": 1,
//             "scoresFor": 2,
//             "scoresAgainst": 3,
//             "id": 1229170,
//             "losses": 1,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "-1"
//           },
//           {
//             "team": {
//               "name": "Manchester United U21",
//               "slug": "manchester-united-u21",
//               "shortName": "Man Utd U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 14497,
//               "nameCode": "MU",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36380,
//               "teamColors": {
//                 "primary": "#ff0000",
//                 "secondary": "#373737",
//                 "text": "#373737"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Манчестер Юнайтед"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 6,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 2,
//             "scoresAgainst": 2,
//             "id": 1229169,
//             "losses": 0,
//             "draws": 2,
//             "points": 2,
//             "scoreDiffFormatted": "0"
//           },
//           {
//             "team": {
//               "name": "Sparta Praha B U21",
//               "slug": "sparta-praha-b-u21",
//               "shortName": "Sparta Praha B  U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2767,
//               "nameCode": "SPA",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 4866,
//               "teamColors": {
//                 "primary": "#942536",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "АК Спарта Прага Б"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 7,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 2,
//             "scoresAgainst": 5,
//             "id": 1229171,
//             "losses": 2,
//             "draws": 0,
//             "points": 0,
//             "scoreDiffFormatted": "-3"
//           },
//           {
//             "team": {
//               "name": "Liverpool U21",
//               "slug": "liverpool-u21",
//               "shortName": "Liverpool U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 7771,
//               "nameCode": "LFC",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36556,
//               "teamColors": {
//                 "primary": "#cc0000",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Ливерпуль"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 8,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 3,
//             "scoresAgainst": 8,
//             "id": 1229172,
//             "losses": 2,
//             "draws": 0,
//             "points": 0,
//             "scoreDiffFormatted": "-5"
//           }
//         ],
//         "id": 134095,
//         "updatedAtTimestamp": 1724277156
//       },
//       {
//         "tournament": {
//           "name": "U21 Premier League International Cup, Group D",
//           "slug": "u21-premier-league-international-cup-group-d",
//           "category": {
//             "name": "World",
//             "slug": "world",
//             "sport": {
//               "name": "Football",
//               "slug": "football",
//               "id": 1
//             },
//             "id": 1468,
//             "flag": "international"
//           },
//           "uniqueTournament": {
//             "name": "Premier League International Cup",
//             "slug": "premier-league-international-cup",
//             "category": {
//               "name": "World",
//               "slug": "world",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "id": 1468,
//               "flag": "international"
//             },
//             "userCount": 1576,
//             "hasPerformanceGraphFeature": false,
//             "id": 20048,
//             "displayInverseHomeAwayTeams": false
//           },
//           "priority": 0,
//           "isGroup": true,
//           "groupName": "Group D",
//           "id": 137564
//         },
//         "type": "total",
//         "name": "U21 Premier League International Cup 24/25, Group D",
//         "descriptions": [],
//         "rows": [
//           {
//             "team": {
//               "name": "Brighton & Hove Albion U21",
//               "slug": "brighton-and-hove-albion-u21",
//               "shortName": "Brighton U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 3238,
//               "nameCode": "BHA",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 169150,
//               "teamColors": {
//                 "primary": "#0054a6",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Брайтон энд Хоув Альбион"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 1,
//             "matches": 3,
//             "wins": 3,
//             "scoresFor": 6,
//             "scoresAgainst": 0,
//             "id": 1229173,
//             "losses": 0,
//             "draws": 0,
//             "points": 9,
//             "scoreDiffFormatted": "+6"
//           },
//           {
//             "team": {
//               "name": "Chelsea U21",
//               "slug": "chelsea-u21",
//               "shortName": "Chelsea U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 13535,
//               "nameCode": "CFC",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 36539,
//               "teamColors": {
//                 "primary": "#0310a7",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Челси"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "promotion": {
//               "text": "Playoffs",
//               "id": 6
//             },
//             "position": 2,
//             "matches": 3,
//             "wins": 2,
//             "scoresFor": 8,
//             "scoresAgainst": 5,
//             "id": 1229174,
//             "losses": 1,
//             "draws": 0,
//             "points": 6,
//             "scoreDiffFormatted": "+3"
//           },
//           {
//             "team": {
//               "name": "Southampton U21",
//               "slug": "southampton-u21",
//               "shortName": "Southampton U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 1657,
//               "nameCode": "SOU",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 77871,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#ff0000",
//                 "text": "#ff0000"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Саутгемптон"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 3,
//             "matches": 1,
//             "wins": 1,
//             "scoresFor": 4,
//             "scoresAgainst": 1,
//             "id": 1229179,
//             "losses": 0,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "+3"
//           },
//           {
//             "team": {
//               "name": "Sporting CP B",
//               "slug": "sporting-cp-b",
//               "shortName": "Sporting B",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 5677,
//               "nameCode": "SCP",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 74467,
//               "teamColors": {
//                 "primary": "#009966",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Спортинг Лиссабон Б"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 4,
//             "matches": 1,
//             "wins": 1,
//             "scoresFor": 2,
//             "scoresAgainst": 1,
//             "id": 1229180,
//             "losses": 0,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "+1"
//           },
//           {
//             "team": {
//               "name": "RB Leipzig U21",
//               "slug": "rb-leipzig-u21",
//               "shortName": "RB Leipzig U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 67,
//               "nameCode": "LEI",
//               "national": false,
//               "type": 0,
//               "id": 828227,
//               "teamColors": {
//                 "primary": "#374df5",
//                 "secondary": "#374df5",
//                 "text": "#ffffff"
//               }
//             },
//             "descriptions": [],
//             "position": 5,
//             "matches": 3,
//             "wins": 1,
//             "scoresFor": 5,
//             "scoresAgainst": 9,
//             "id": 1229176,
//             "losses": 2,
//             "draws": 0,
//             "points": 3,
//             "scoreDiffFormatted": "-4"
//           },
//           {
//             "team": {
//               "name": "Crystal Palace U21",
//               "slug": "crystal-palace-u21",
//               "shortName": "Crystal Palace U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 2926,
//               "nameCode": "CRP",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 77869,
//               "teamColors": {
//                 "primary": "#0033ff",
//                 "secondary": "#b90d2b",
//                 "text": "#b90d2b"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Кристал Пэлас ФК"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 6,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 3,
//             "scoresAgainst": 4,
//             "id": 1229175,
//             "losses": 1,
//             "draws": 1,
//             "points": 1,
//             "scoreDiffFormatted": "-1"
//           },
//           {
//             "team": {
//               "name": "Real Sociedad B",
//               "slug": "real-sociedad-b",
//               "shortName": "Real Sociedad B",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 4226,
//               "nameCode": "RSO",
//               "national": false,
//               "type": 0,
//               "id": 24360,
//               "teamColors": {
//                 "primary": "#ffffff",
//                 "secondary": "#0000ff",
//                 "text": "#0000ff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "Реал Сосьедад Б"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 7,
//             "matches": 3,
//             "wins": 0,
//             "scoresFor": 3,
//             "scoresAgainst": 5,
//             "id": 1229177,
//             "losses": 2,
//             "draws": 1,
//             "points": 1,
//             "scoreDiffFormatted": "-2"
//           },
//           {
//             "team": {
//               "name": "Anderlecht Reserve U21",
//               "slug": "anderlecht-reserve",
//               "shortName": "Anderlecht U21",
//               "gender": "M",
//               "sport": {
//                 "name": "Football",
//                 "slug": "football",
//                 "id": 1
//               },
//               "userCount": 1364,
//               "nameCode": "AND",
//               "disabled": false,
//               "national": false,
//               "type": 0,
//               "id": 76465,
//               "teamColors": {
//                 "primary": "#330066",
//                 "secondary": "#ffffff",
//                 "text": "#ffffff"
//               },
//               "fieldTranslations": {
//                 "nameTranslation": {
//                   "ru": "РСК Андерлехт"
//                 },
//                 "shortNameTranslation": {}
//               }
//             },
//             "descriptions": [],
//             "position": 8,
//             "matches": 2,
//             "wins": 0,
//             "scoresFor": 0,
//             "scoresAgainst": 6,
//             "id": 1229178,
//             "losses": 2,
//             "draws": 0,
//             "points": 0,
//             "scoreDiffFormatted": "-6"
//           }
//         ],
//         "id": 134096,
//         "updatedAtTimestamp": 1724277156
//       }
//     ]
//   }
// }
