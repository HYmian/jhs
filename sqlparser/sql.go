//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const OFFSET = 57359
const CALL = 57360
const ALL = 57361
const DISTINCT = 57362
const AS = 57363
const EXISTS = 57364
const IN = 57365
const IS = 57366
const LIKE = 57367
const BETWEEN = 57368
const NULL = 57369
const ASC = 57370
const DESC = 57371
const VALUES = 57372
const INTO = 57373
const DUPLICATE = 57374
const KEY = 57375
const DEFAULT = 57376
const SET = 57377
const LOCK = 57378
const COLLATE = 57379
const ID = 57380
const STRING = 57381
const NUMBER = 57382
const VALUE_ARG = 57383
const COMMENT = 57384
const LE = 57385
const GE = 57386
const NE = 57387
const NULL_SAFE_EQUAL = 57388
const UNION = 57389
const MINUS = 57390
const EXCEPT = 57391
const INTERSECT = 57392
const JOIN = 57393
const STRAIGHT_JOIN = 57394
const LEFT = 57395
const RIGHT = 57396
const INNER = 57397
const OUTER = 57398
const CROSS = 57399
const NATURAL = 57400
const USE = 57401
const FORCE = 57402
const ON = 57403
const AND = 57404
const OR = 57405
const NOT = 57406
const UNARY = 57407
const CASE = 57408
const WHEN = 57409
const THEN = 57410
const ELSE = 57411
const END = 57412
const BEGIN = 57413
const COMMIT = 57414
const ROLLBACK = 57415
const START = 57416
const TRANSACTION = 57417
const NAMES = 57418
const REPLACE = 57419
const ADMIN = 57420
const SHOW = 57421
const DATABASES = 57422
const TABLES = 57423
const PROXY = 57424
const VARIABLES = 57425
const CREATE = 57426
const ALTER = 57427
const DROP = 57428
const RENAME = 57429
const TABLE = 57430
const INDEX = 57431
const VIEW = 57432
const TO = 57433
const IGNORE = 57434
const IF = 57435
const UNIQUE = 57436
const USING = 57437

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"OFFSET",
	"CALL",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"COLLATE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"START",
	"TRANSACTION",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
	"PROXY",
	"VARIABLES",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 224
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 651

var yyAct = []int{

	117, 329, 159, 76, 114, 282, 364, 398, 197, 145,
	238, 321, 273, 103, 211, 115, 275, 137, 198, 3,
	207, 97, 220, 78, 104, 172, 173, 125, 37, 38,
	39, 40, 90, 108, 407, 64, 83, 297, 298, 299,
	300, 301, 290, 302, 303, 407, 407, 80, 148, 375,
	85, 79, 47, 87, 49, 68, 167, 91, 50, 52,
	374, 53, 327, 96, 346, 348, 124, 167, 265, 130,
	167, 235, 235, 109, 235, 373, 84, 81, 121, 122,
	123, 94, 55, 56, 57, 144, 156, 267, 86, 409,
	128, 350, 58, 152, 54, 355, 81, 147, 155, 66,
	408, 406, 160, 162, 162, 274, 347, 169, 308, 161,
	163, 354, 171, 126, 127, 199, 154, 326, 255, 200,
	131, 141, 316, 102, 143, 314, 266, 236, 370, 234,
	226, 274, 158, 319, 206, 80, 322, 80, 210, 79,
	286, 79, 216, 215, 74, 195, 196, 129, 203, 77,
	164, 224, 217, 384, 385, 227, 185, 186, 187, 162,
	213, 256, 230, 60, 61, 62, 63, 172, 173, 151,
	109, 244, 216, 133, 231, 135, 372, 248, 157, 233,
	253, 254, 357, 257, 258, 259, 260, 261, 262, 263,
	264, 249, 243, 153, 242, 89, 180, 181, 182, 183,
	184, 185, 186, 187, 109, 109, 246, 247, 371, 80,
	80, 65, 278, 79, 280, 269, 271, 223, 225, 222,
	344, 287, 381, 172, 173, 281, 277, 37, 38, 39,
	40, 288, 285, 80, 136, 136, 292, 79, 293, 183,
	184, 185, 186, 187, 343, 322, 284, 342, 291, 245,
	277, 307, 294, 92, 310, 311, 19, 20, 21, 22,
	340, 380, 235, 382, 359, 341, 309, 338, 242, 35,
	140, 166, 339, 109, 212, 212, 180, 181, 182, 183,
	184, 185, 186, 187, 315, 318, 23, 328, 392, 391,
	325, 352, 324, 19, 180, 181, 182, 183, 184, 185,
	186, 187, 390, 156, 204, 202, 336, 337, 320, 132,
	297, 298, 299, 300, 301, 353, 302, 303, 167, 232,
	295, 136, 356, 242, 242, 201, 241, 241, 80, 208,
	361, 209, 360, 362, 365, 240, 240, 28, 30, 31,
	29, 209, 366, 32, 34, 33, 101, 100, 138, 139,
	24, 25, 27, 26, 306, 376, 65, 170, 312, 81,
	377, 180, 181, 182, 183, 184, 185, 186, 187, 351,
	349, 305, 162, 386, 65, 379, 333, 332, 388, 229,
	228, 165, 67, 394, 395, 365, 149, 146, 397, 396,
	142, 399, 399, 399, 80, 88, 404, 402, 79, 400,
	401, 214, 387, 270, 389, 120, 134, 378, 412, 358,
	124, 411, 413, 130, 414, 93, 405, 73, 313, 41,
	120, 107, 121, 122, 123, 124, 19, 218, 130, 250,
	112, 251, 252, 98, 128, 150, 107, 121, 122, 123,
	43, 44, 45, 46, 71, 112, 69, 330, 99, 128,
	369, 276, 59, 111, 331, 283, 368, 126, 127, 105,
	335, 212, 95, 19, 131, 75, 410, 393, 111, 19,
	42, 17, 126, 127, 105, 16, 15, 14, 13, 131,
	120, 12, 219, 48, 289, 124, 221, 51, 130, 82,
	279, 129, 403, 383, 268, 120, 81, 121, 122, 123,
	124, 363, 367, 130, 334, 112, 129, 317, 205, 128,
	272, 81, 121, 122, 123, 119, 19, 116, 118, 323,
	112, 113, 174, 110, 128, 345, 239, 296, 111, 237,
	106, 304, 126, 127, 168, 70, 36, 72, 124, 131,
	11, 130, 10, 111, 9, 8, 18, 126, 127, 81,
	121, 122, 123, 7, 131, 6, 124, 5, 156, 130,
	4, 2, 128, 1, 0, 0, 129, 81, 121, 122,
	123, 0, 0, 0, 0, 0, 156, 0, 0, 0,
	128, 129, 0, 0, 0, 126, 127, 0, 0, 0,
	0, 0, 131, 0, 0, 175, 179, 177, 178, 0,
	0, 0, 0, 126, 127, 0, 0, 0, 0, 0,
	131, 0, 0, 0, 0, 191, 192, 193, 194, 129,
	188, 189, 190, 180, 181, 182, 183, 184, 185, 186,
	187, 0, 0, 0, 0, 0, 0, 129, 0, 0,
	0, 0, 176, 180, 181, 182, 183, 184, 185, 186,
	187,
}
var yyPact = []int{

	251, -1000, -1000, 175, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -51, -46, -9, -21, -1000, 2,
	-1000, -1000, -1000, 68, 318, 344, 464, 427, -1000, -1000,
	-1000, 424, -1000, 386, 344, 456, 58, -72, -28, 318,
	-1000, -15, 318, -1000, 357, -76, 318, -76, -1000, 384,
	-1000, 453, 318, 423, 300, -1000, 299, 44, -1000, -1000,
	398, -1000, 267, 344, 371, 344, 179, 310, -1000, 222,
	-1000, 42, 352, 54, 318, -1000, 349, -1000, -58, 348,
	413, 102, 318, 344, 423, 529, 453, -1000, 473, 529,
	529, 39, 343, 262, -1000, -1000, 336, 33, 155, 572,
	-1000, 473, 458, -1000, -1000, -1000, 529, 278, 258, -1000,
	257, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 529, -1000, 294, 321, 451, 321, 364, -1000, -1000,
	529, 318, -1000, 405, -88, -1000, 117, -1000, 342, -1000,
	-1000, 341, -1000, 284, -1000, 552, 511, 423, 155, 572,
	552, 18, 552, 16, -1000, -1000, 289, 398, -1000, -1000,
	318, 173, 473, 473, 529, 256, 406, 529, 529, 91,
	529, 529, 529, 529, 529, 529, 529, 529, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -43, 15, -24, 572,
	-1000, 383, 398, -1000, 464, 23, 552, 421, 321, 321,
	265, 442, 473, -1000, 310, 552, -1000, -1000, -1000, 73,
	318, -1000, -64, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 421, 321, -1000, -1000, 529, -1000, 264, 253, 333,
	288, 29, -1000, -1000, -1000, -1000, -1000, -1000, 552, -1000,
	256, 529, 529, 552, 290, -1000, 391, 165, 165, 165,
	80, 80, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 14,
	398, 11, 49, -1000, 473, 69, 256, 175, 178, 6,
	-1000, 442, 432, 440, 155, -1000, 339, -1000, -1000, 338,
	-1000, -1000, 179, 552, 449, 289, 289, -1000, -1000, 210,
	203, 190, 187, 163, -1, -1000, 332, -20, 331, -1000,
	552, 223, 529, -1000, -1000, 0, -1000, 10, -1000, 529,
	99, -1000, 377, 208, -1000, -1000, -1000, 321, 432, -1000,
	529, 529, -1000, -1000, 444, 436, 253, 61, -1000, 151,
	-1000, 119, -1000, -1000, -1000, -1000, -29, -44, -55, -1000,
	-1000, -1000, 529, 552, -1000, -1000, 552, 529, 374, 256,
	-1000, -1000, 205, 207, -1000, 125, -1000, 442, 473, 529,
	473, -1000, -1000, 255, 242, 241, 552, 552, 460, -1000,
	529, 529, 529, -1000, -1000, -1000, 432, 155, 206, 155,
	318, 318, 318, 321, 552, 552, -1000, 380, -10, -1000,
	-11, -22, 179, -1000, 459, 388, -1000, 318, -1000, -1000,
	-1000, 318, -1000, 318, -1000,
}
var yyPgo = []int{

	0, 563, 561, 18, 560, 557, 555, 553, 546, 545,
	544, 542, 540, 419, 537, 536, 535, 13, 24, 534,
	531, 530, 529, 10, 527, 526, 99, 525, 7, 14,
	33, 523, 522, 16, 521, 2, 17, 15, 8, 519,
	518, 27, 517, 4, 515, 510, 12, 508, 507, 504,
	502, 5, 501, 6, 493, 1, 492, 20, 490, 11,
	3, 23, 195, 489, 487, 486, 484, 483, 482, 0,
	9, 481, 478, 477, 476, 475, 471, 81, 21, 470,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 4, 4, 74, 74, 5, 6, 7, 7, 7,
	71, 71, 72, 73, 76, 75, 75, 75, 75, 9,
	9, 9, 10, 10, 10, 11, 12, 12, 12, 8,
	8, 79, 13, 14, 14, 15, 15, 15, 15, 15,
	16, 16, 17, 17, 18, 18, 18, 21, 21, 19,
	19, 19, 22, 22, 23, 23, 23, 23, 20, 20,
	20, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	25, 25, 25, 26, 26, 27, 27, 27, 27, 28,
	28, 29, 29, 78, 78, 78, 77, 77, 30, 30,
	30, 30, 30, 31, 31, 31, 31, 31, 31, 31,
	31, 31, 31, 32, 32, 32, 32, 32, 32, 32,
	33, 33, 39, 39, 37, 37, 41, 38, 38, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 40, 40, 42, 42,
	42, 44, 47, 47, 45, 45, 46, 48, 48, 43,
	43, 34, 34, 34, 34, 49, 49, 50, 50, 51,
	51, 52, 52, 53, 54, 54, 54, 55, 55, 55,
	55, 56, 56, 56, 57, 57, 58, 58, 59, 59,
	60, 60, 61, 36, 36, 62, 62, 63, 63, 64,
	64, 65, 65, 65, 65, 65, 66, 66, 67, 67,
	68, 68, 69, 70,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 4, 12,
	3, 7, 7, 6, 6, 8, 7, 3, 4, 6,
	1, 2, 1, 1, 5, 2, 4, 5, 3, 5,
	8, 4, 6, 7, 4, 5, 4, 5, 5, 5,
	4, 0, 2, 0, 2, 1, 2, 1, 1, 1,
	0, 1, 1, 3, 1, 2, 3, 1, 1, 0,
	1, 2, 1, 3, 3, 3, 3, 5, 0, 1,
	2, 1, 1, 2, 3, 2, 3, 2, 2, 2,
	1, 3, 1, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 0, 2, 2, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 1, 3, 3, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	4, 0, 2, 4, 0, 3, 1, 3, 0, 5,
	1, 3, 3, 1, 1, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -9, -10,
	-11, -12, -71, -72, -73, -74, -75, -76, -8, 5,
	6, 7, 8, 35, 99, 100, 102, 101, 86, 89,
	87, 88, 92, 94, 93, 18, -15, 52, 53, 54,
	55, -13, -79, -13, -13, -13, -13, 103, -67, 105,
	109, -64, 105, 107, 103, 103, 104, 105, 90, -13,
	95, 96, 97, 98, -69, 38, -26, 38, -3, 19,
	-16, 20, -14, 31, -26, 9, -60, 91, -61, -43,
	-69, 38, -63, 108, 104, -69, 103, -69, 38, -62,
	108, -69, -62, 31, -77, 9, -69, -78, 10, 25,
	47, 47, 79, -17, -18, 76, -21, 38, -30, -35,
	-31, 70, 47, -34, -43, -37, -42, -69, -40, -44,
	22, 39, 40, 41, 27, -41, 74, 75, 51, 108,
	30, 81, 42, -26, 35, -26, 56, -36, 38, 39,
	48, 79, 38, 70, -69, -70, 38, -70, 106, 38,
	22, 67, -69, -26, -78, -35, 47, -77, -30, -35,
	-35, -38, -35, -38, 111, 38, 9, 56, -19, -69,
	21, 79, 68, 69, -32, 23, 70, 25, 26, 24,
	71, 72, 73, 74, 75, 76, 77, 78, 48, 49,
	50, 43, 44, 45, 46, -30, -30, -38, -3, -35,
	-35, 47, 47, -41, 47, -47, -35, -57, 35, 47,
	-60, -29, 10, -61, 37, -35, -69, -70, 22, -68,
	110, -65, 102, 100, 34, 101, 13, 38, 38, 38,
	-70, -57, 35, -78, 111, 56, 111, -22, -23, -25,
	47, 38, -41, -18, -69, 76, -30, -30, -35, -37,
	23, 25, 26, -35, -35, 27, 70, -35, -35, -35,
	-35, -35, -35, -35, -35, 111, 111, 111, 111, -17,
	20, -17, -45, -46, 82, -33, 30, -3, -60, -58,
	-43, -29, -51, 13, -30, -36, 67, -69, -70, -66,
	106, -33, -60, -35, -29, 56, -24, 57, 58, 59,
	60, 61, 63, 64, -20, 38, 21, -23, 79, -37,
	-35, -35, 68, 27, 111, -17, 111, -48, -46, 84,
	-30, -59, 67, -39, -37, -59, 111, 56, -51, -55,
	15, 14, 38, 38, -49, 11, -23, -23, 57, 62,
	57, 62, 57, 57, 57, -27, 65, 107, 66, 38,
	111, 38, 68, -35, 111, 85, -35, 83, 32, 56,
	-43, -55, -35, -52, -53, -35, -70, -50, 12, 14,
	67, 57, 57, 104, 104, 104, -35, -35, 33, -37,
	56, 17, 56, -54, 28, 29, -51, -30, -38, -30,
	47, 47, 47, 7, -35, -35, -53, -55, -28, -69,
	-28, -28, -60, -56, 16, 36, 111, 56, 111, 111,
	7, 23, -69, -69, -69,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 51,
	51, 51, 51, 51, 218, 209, 0, 0, 30, 0,
	32, 33, 51, 0, 0, 0, 0, 55, 57, 58,
	59, 60, 53, 0, 0, 0, 0, 207, 0, 0,
	219, 0, 0, 210, 0, 205, 0, 205, 31, 0,
	35, 106, 0, 103, 0, 222, 0, 93, 20, 56,
	0, 61, 52, 0, 0, 0, 27, 0, 200, 0,
	169, 222, 0, 0, 0, 223, 0, 223, 0, 0,
	0, 0, 0, 0, 103, 0, 106, 38, 0, 0,
	0, 0, 0, 18, 62, 64, 69, 222, 67, 68,
	108, 0, 0, 139, 140, 141, 0, 169, 0, 155,
	0, 171, 172, 173, 174, 135, 158, 159, 160, 156,
	157, 162, 54, 194, 0, 101, 0, 28, 203, 204,
	0, 0, 223, 0, 220, 41, 0, 44, 0, 46,
	206, 0, 223, 194, 36, 107, 0, 103, 104, 0,
	105, 0, 137, 0, 50, 94, 0, 0, 65, 70,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 123, 124,
	125, 126, 127, 128, 129, 111, 0, 0, 0, 137,
	150, 0, 0, 122, 0, 0, 163, 0, 0, 0,
	101, 179, 0, 201, 0, 202, 170, 39, 208, 0,
	0, 223, 216, 211, 212, 213, 214, 215, 45, 47,
	48, 0, 0, 37, 34, 0, 49, 101, 72, 78,
	0, 90, 92, 63, 71, 66, 109, 110, 113, 114,
	0, 0, 0, 116, 0, 120, 0, 142, 143, 144,
	145, 146, 147, 148, 149, 112, 134, 136, 151, 0,
	0, 0, 167, 164, 0, 198, 0, 131, 198, 0,
	196, 179, 187, 0, 102, 29, 0, 221, 42, 0,
	217, 23, 24, 138, 175, 0, 0, 81, 82, 0,
	0, 0, 0, 0, 95, 79, 0, 0, 0, 115,
	117, 0, 0, 121, 152, 0, 154, 0, 165, 0,
	0, 21, 0, 130, 132, 22, 195, 0, 187, 26,
	0, 0, 223, 43, 177, 0, 73, 76, 83, 0,
	85, 0, 87, 88, 89, 74, 0, 0, 0, 80,
	75, 91, 0, 118, 153, 161, 168, 0, 0, 0,
	197, 25, 188, 180, 181, 184, 40, 179, 0, 0,
	0, 84, 86, 0, 0, 0, 119, 166, 0, 133,
	0, 0, 0, 183, 185, 186, 187, 178, 176, 77,
	0, 0, 0, 0, 189, 190, 182, 191, 0, 99,
	0, 0, 199, 19, 0, 0, 96, 0, 97, 98,
	192, 0, 100, 0, 193,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 78, 71, 3,
	47, 111, 76, 74, 56, 75, 79, 77, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 48, 50, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 51,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 52, 53, 54, 55, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:171
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:177
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		yyVAL.statement = yyS[yypt-0].statement
	case 18:
		//line sql.y:198
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 19:
		//line sql.y:202
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 20:
		//line sql.y:206
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 21:
		//line sql.y:213
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:217
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 23:
		//line sql.y:229
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 24:
		//line sql.y:233
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 25:
		//line sql.y:246
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:252
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 27:
		//line sql.y:258
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 28:
		//line sql.y:262
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyS[yypt-0].valExpr}}}
		}
	case 29:
		//line sql.y:266
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-4].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyS[yypt-2].valExpr}, &UpdateExpr{Name: &ColName{Name: []byte("collate")}, Expr: yyS[yypt-0].valExpr}}}
		}
	case 30:
		//line sql.y:272
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		//line sql.y:276
		{
			yyVAL.statement = &Begin{}
		}
	case 32:
		//line sql.y:282
		{
			yyVAL.statement = &Commit{}
		}
	case 33:
		//line sql.y:288
		{
			yyVAL.statement = &Rollback{}
		}
	case 34:
		//line sql.y:294
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 35:
		//line sql.y:300
		{
			yyVAL.statement = &Show{Section: "databases"}
		}
	case 36:
		//line sql.y:304
		{
			yyVAL.statement = &Show{Section: "tables", From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 37:
		//line sql.y:308
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyS[yypt-2].bytes), From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 38:
		//line sql.y:312
		{
			yyVAL.statement = &Show{Section: "variables", LikeOrWhere: yyS[yypt-0].expr}
		}
	case 39:
		//line sql.y:318
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 40:
		//line sql.y:322
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 41:
		//line sql.y:327
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 42:
		//line sql.y:333
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 43:
		//line sql.y:337
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 44:
		//line sql.y:342
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 45:
		//line sql.y:348
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 46:
		//line sql.y:354
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 47:
		//line sql.y:358
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 48:
		//line sql.y:363
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 49:
		//line sql.y:369
		{
			yyVAL.statement = &Call{Procedure: yyS[yypt-3].tableName, Values: yyS[yypt-1].valExprs}
		}
	case 50:
		//line sql.y:373
		{
			yyVAL.statement = &Call{Procedure: yyS[yypt-2].tableName}
		}
	case 51:
		//line sql.y:378
		{
			SetAllowComments(yylex, true)
		}
	case 52:
		//line sql.y:382
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 53:
		//line sql.y:388
		{
			yyVAL.bytes2 = nil
		}
	case 54:
		//line sql.y:392
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 55:
		//line sql.y:398
		{
			yyVAL.str = AST_UNION
		}
	case 56:
		//line sql.y:402
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 57:
		//line sql.y:406
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 58:
		//line sql.y:410
		{
			yyVAL.str = AST_EXCEPT
		}
	case 59:
		//line sql.y:414
		{
			yyVAL.str = AST_INTERSECT
		}
	case 60:
		//line sql.y:419
		{
			yyVAL.str = ""
		}
	case 61:
		//line sql.y:423
		{
			yyVAL.str = AST_DISTINCT
		}
	case 62:
		//line sql.y:429
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 63:
		//line sql.y:433
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 64:
		//line sql.y:439
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 65:
		//line sql.y:443
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 66:
		//line sql.y:447
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 67:
		//line sql.y:453
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 68:
		//line sql.y:457
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 69:
		//line sql.y:462
		{
			yyVAL.bytes = nil
		}
	case 70:
		//line sql.y:466
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 71:
		//line sql.y:470
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 72:
		//line sql.y:476
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 73:
		//line sql.y:480
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 74:
		//line sql.y:486
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 75:
		//line sql.y:490
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 76:
		//line sql.y:494
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 77:
		//line sql.y:498
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 78:
		//line sql.y:503
		{
			yyVAL.bytes = nil
		}
	case 79:
		//line sql.y:507
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 80:
		//line sql.y:511
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 81:
		//line sql.y:517
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		//line sql.y:521
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 83:
		//line sql.y:525
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 84:
		//line sql.y:529
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 85:
		//line sql.y:533
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 86:
		//line sql.y:537
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 87:
		//line sql.y:541
		{
			yyVAL.str = AST_JOIN
		}
	case 88:
		//line sql.y:545
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 89:
		//line sql.y:549
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 90:
		//line sql.y:555
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 91:
		//line sql.y:559
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 92:
		//line sql.y:563
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 93:
		//line sql.y:569
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 94:
		//line sql.y:573
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 95:
		//line sql.y:578
		{
			yyVAL.indexHints = nil
		}
	case 96:
		//line sql.y:582
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 97:
		//line sql.y:586
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 98:
		//line sql.y:590
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 99:
		//line sql.y:596
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 100:
		//line sql.y:600
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 101:
		//line sql.y:605
		{
			yyVAL.boolExpr = nil
		}
	case 102:
		//line sql.y:609
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 103:
		//line sql.y:614
		{
			yyVAL.expr = nil
		}
	case 104:
		//line sql.y:618
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 105:
		//line sql.y:622
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 106:
		//line sql.y:627
		{
			yyVAL.valExpr = nil
		}
	case 107:
		//line sql.y:631
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 108:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 109:
		//line sql.y:638
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 110:
		//line sql.y:642
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 111:
		//line sql.y:646
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 112:
		//line sql.y:650
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 113:
		//line sql.y:656
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 114:
		//line sql.y:660
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 115:
		//line sql.y:664
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 116:
		//line sql.y:668
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 117:
		//line sql.y:672
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 118:
		//line sql.y:676
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 119:
		//line sql.y:680
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 120:
		//line sql.y:684
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 121:
		//line sql.y:688
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 122:
		//line sql.y:692
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 123:
		//line sql.y:698
		{
			yyVAL.str = AST_EQ
		}
	case 124:
		//line sql.y:702
		{
			yyVAL.str = AST_LT
		}
	case 125:
		//line sql.y:706
		{
			yyVAL.str = AST_GT
		}
	case 126:
		//line sql.y:710
		{
			yyVAL.str = AST_LE
		}
	case 127:
		//line sql.y:714
		{
			yyVAL.str = AST_GE
		}
	case 128:
		//line sql.y:718
		{
			yyVAL.str = AST_NE
		}
	case 129:
		//line sql.y:722
		{
			yyVAL.str = AST_NSE
		}
	case 130:
		//line sql.y:728
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 131:
		//line sql.y:732
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 132:
		//line sql.y:738
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 133:
		//line sql.y:742
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 134:
		//line sql.y:748
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 135:
		//line sql.y:752
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 136:
		//line sql.y:758
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 137:
		//line sql.y:764
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 138:
		//line sql.y:768
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 139:
		//line sql.y:774
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 140:
		//line sql.y:778
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 141:
		//line sql.y:782
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 142:
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 143:
		//line sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 144:
		//line sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 145:
		//line sql.y:798
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 146:
		//line sql.y:802
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 147:
		//line sql.y:806
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 148:
		//line sql.y:810
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 149:
		//line sql.y:814
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 150:
		//line sql.y:818
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 151:
		//line sql.y:833
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 152:
		//line sql.y:837
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 153:
		//line sql.y:841
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 154:
		//line sql.y:845
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 155:
		//line sql.y:849
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 156:
		//line sql.y:855
		{
			yyVAL.bytes = IF_BYTES
		}
	case 157:
		//line sql.y:859
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 158:
		//line sql.y:865
		{
			yyVAL.byt = AST_UPLUS
		}
	case 159:
		//line sql.y:869
		{
			yyVAL.byt = AST_UMINUS
		}
	case 160:
		//line sql.y:873
		{
			yyVAL.byt = AST_TILDA
		}
	case 161:
		//line sql.y:879
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 162:
		//line sql.y:884
		{
			yyVAL.valExpr = nil
		}
	case 163:
		//line sql.y:888
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 164:
		//line sql.y:894
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 165:
		//line sql.y:898
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 166:
		//line sql.y:904
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 167:
		//line sql.y:909
		{
			yyVAL.valExpr = nil
		}
	case 168:
		//line sql.y:913
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 169:
		//line sql.y:919
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 170:
		//line sql.y:923
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 171:
		//line sql.y:929
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 172:
		//line sql.y:933
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 173:
		//line sql.y:937
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 174:
		//line sql.y:941
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 175:
		//line sql.y:946
		{
			yyVAL.valExprs = nil
		}
	case 176:
		//line sql.y:950
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 177:
		//line sql.y:955
		{
			yyVAL.boolExpr = nil
		}
	case 178:
		//line sql.y:959
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 179:
		//line sql.y:964
		{
			yyVAL.orderBy = nil
		}
	case 180:
		//line sql.y:968
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 181:
		//line sql.y:974
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 182:
		//line sql.y:978
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 183:
		//line sql.y:984
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 184:
		//line sql.y:989
		{
			yyVAL.str = AST_ASC
		}
	case 185:
		//line sql.y:993
		{
			yyVAL.str = AST_ASC
		}
	case 186:
		//line sql.y:997
		{
			yyVAL.str = AST_DESC
		}
	case 187:
		//line sql.y:1002
		{
			yyVAL.limit = nil
		}
	case 188:
		//line sql.y:1006
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 189:
		//line sql.y:1010
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 190:
		//line sql.y:1014
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 191:
		//line sql.y:1019
		{
			yyVAL.str = ""
		}
	case 192:
		//line sql.y:1023
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 193:
		//line sql.y:1027
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 194:
		//line sql.y:1040
		{
			yyVAL.columns = nil
		}
	case 195:
		//line sql.y:1044
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 196:
		//line sql.y:1050
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 197:
		//line sql.y:1054
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 198:
		//line sql.y:1059
		{
			yyVAL.updateExprs = nil
		}
	case 199:
		//line sql.y:1063
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 200:
		//line sql.y:1069
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 201:
		//line sql.y:1073
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 202:
		//line sql.y:1079
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 203:
		//line sql.y:1085
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 204:
		//line sql.y:1089
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 205:
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1101
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1104
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1106
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1112
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		//line sql.y:1114
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		//line sql.y:1116
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		//line sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		//line sql.y:1121
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		//line sql.y:1123
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		//line sql.y:1126
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		//line sql.y:1128
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		//line sql.y:1131
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		//line sql.y:1133
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		//line sql.y:1137
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 223:
		//line sql.y:1142
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
