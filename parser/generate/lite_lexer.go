// Code generated from LiteLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 98, 592,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52,
	3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3,
	57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62,
	3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 67, 3,
	67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70,
	3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3,
	72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75,
	3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77, 3, 77, 3,
	78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 80, 3, 80,
	3, 80, 3, 80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82, 3,
	82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 85,
	3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3,
	86, 3, 87, 6, 87, 495, 10, 87, 13, 87, 14, 87, 496, 3, 88, 3, 88, 3, 89,
	3, 89, 3, 89, 3, 89, 7, 89, 505, 10, 89, 12, 89, 14, 89, 508, 11, 89, 3,
	89, 3, 89, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 7, 90, 518, 10, 90,
	12, 90, 14, 90, 521, 11, 90, 3, 90, 3, 90, 3, 91, 3, 91, 6, 91, 527, 10,
	91, 13, 91, 14, 91, 528, 3, 92, 3, 92, 7, 92, 533, 10, 92, 12, 92, 14,
	92, 536, 11, 92, 3, 93, 3, 93, 3, 94, 3, 94, 3, 94, 3, 94, 3, 94, 7, 94,
	545, 10, 94, 12, 94, 14, 94, 548, 11, 94, 3, 94, 3, 94, 3, 94, 3, 94, 3,
	94, 3, 94, 3, 95, 3, 95, 3, 95, 3, 95, 7, 95, 560, 10, 95, 12, 95, 14,
	95, 563, 11, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 96, 3, 96, 7, 96,
	572, 10, 96, 12, 96, 14, 96, 575, 11, 96, 3, 96, 3, 96, 3, 96, 3, 96, 3,
	97, 5, 97, 582, 10, 97, 3, 97, 3, 97, 3, 98, 6, 98, 587, 10, 98, 13, 98,
	14, 98, 588, 3, 98, 3, 98, 7, 506, 519, 546, 561, 573, 2, 99, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129,
	66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71, 141, 72, 143, 73, 145,
	74, 147, 75, 149, 76, 151, 77, 153, 78, 155, 79, 157, 80, 159, 81, 161,
	82, 163, 83, 165, 84, 167, 85, 169, 86, 171, 87, 173, 88, 175, 2, 177,
	89, 179, 90, 181, 91, 183, 92, 185, 93, 187, 94, 189, 95, 191, 96, 193,
	97, 195, 98, 3, 2, 8, 3, 2, 50, 59, 9, 2, 36, 36, 94, 94, 100, 100, 104,
	104, 112, 112, 116, 116, 118, 118, 8, 2, 94, 94, 100, 100, 104, 104, 112,
	112, 116, 116, 118, 118, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 67,
	92, 99, 124, 4, 2, 11, 11, 34, 34, 2, 603, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67,
	3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2,
	75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2,
	2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2,
	2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2,
	2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105,
	3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2,
	2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3,
	2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2,
	127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2,
	2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141,
	3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2,
	2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3,
	2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2,
	163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2,
	2, 2, 2, 171, 3, 2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179,
	3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2,
	2, 187, 3, 2, 2, 2, 2, 189, 3, 2, 2, 2, 2, 191, 3, 2, 2, 2, 2, 193, 3,
	2, 2, 2, 2, 195, 3, 2, 2, 2, 3, 197, 3, 2, 2, 2, 5, 202, 3, 2, 2, 2, 7,
	205, 3, 2, 2, 2, 9, 212, 3, 2, 2, 2, 11, 218, 3, 2, 2, 2, 13, 224, 3, 2,
	2, 2, 15, 229, 3, 2, 2, 2, 17, 237, 3, 2, 2, 2, 19, 242, 3, 2, 2, 2, 21,
	246, 3, 2, 2, 2, 23, 249, 3, 2, 2, 2, 25, 252, 3, 2, 2, 2, 27, 259, 3,
	2, 2, 2, 29, 269, 3, 2, 2, 2, 31, 280, 3, 2, 2, 2, 33, 283, 3, 2, 2, 2,
	35, 286, 3, 2, 2, 2, 37, 289, 3, 2, 2, 2, 39, 292, 3, 2, 2, 2, 41, 295,
	3, 2, 2, 2, 43, 298, 3, 2, 2, 2, 45, 301, 3, 2, 2, 2, 47, 304, 3, 2, 2,
	2, 49, 307, 3, 2, 2, 2, 51, 310, 3, 2, 2, 2, 53, 313, 3, 2, 2, 2, 55, 316,
	3, 2, 2, 2, 57, 319, 3, 2, 2, 2, 59, 322, 3, 2, 2, 2, 61, 325, 3, 2, 2,
	2, 63, 329, 3, 2, 2, 2, 65, 332, 3, 2, 2, 2, 67, 334, 3, 2, 2, 2, 69, 336,
	3, 2, 2, 2, 71, 339, 3, 2, 2, 2, 73, 342, 3, 2, 2, 2, 75, 345, 3, 2, 2,
	2, 77, 348, 3, 2, 2, 2, 79, 351, 3, 2, 2, 2, 81, 353, 3, 2, 2, 2, 83, 355,
	3, 2, 2, 2, 85, 357, 3, 2, 2, 2, 87, 359, 3, 2, 2, 2, 89, 361, 3, 2, 2,
	2, 91, 363, 3, 2, 2, 2, 93, 365, 3, 2, 2, 2, 95, 367, 3, 2, 2, 2, 97, 369,
	3, 2, 2, 2, 99, 371, 3, 2, 2, 2, 101, 374, 3, 2, 2, 2, 103, 376, 3, 2,
	2, 2, 105, 378, 3, 2, 2, 2, 107, 380, 3, 2, 2, 2, 109, 382, 3, 2, 2, 2,
	111, 384, 3, 2, 2, 2, 113, 386, 3, 2, 2, 2, 115, 388, 3, 2, 2, 2, 117,
	390, 3, 2, 2, 2, 119, 392, 3, 2, 2, 2, 121, 394, 3, 2, 2, 2, 123, 396,
	3, 2, 2, 2, 125, 398, 3, 2, 2, 2, 127, 400, 3, 2, 2, 2, 129, 402, 3, 2,
	2, 2, 131, 404, 3, 2, 2, 2, 133, 407, 3, 2, 2, 2, 135, 410, 3, 2, 2, 2,
	137, 414, 3, 2, 2, 2, 139, 418, 3, 2, 2, 2, 141, 422, 3, 2, 2, 2, 143,
	426, 3, 2, 2, 2, 145, 430, 3, 2, 2, 2, 147, 434, 3, 2, 2, 2, 149, 438,
	3, 2, 2, 2, 151, 442, 3, 2, 2, 2, 153, 446, 3, 2, 2, 2, 155, 450, 3, 2,
	2, 2, 157, 455, 3, 2, 2, 2, 159, 459, 3, 2, 2, 2, 161, 463, 3, 2, 2, 2,
	163, 468, 3, 2, 2, 2, 165, 472, 3, 2, 2, 2, 167, 476, 3, 2, 2, 2, 169,
	481, 3, 2, 2, 2, 171, 487, 3, 2, 2, 2, 173, 494, 3, 2, 2, 2, 175, 498,
	3, 2, 2, 2, 177, 500, 3, 2, 2, 2, 179, 511, 3, 2, 2, 2, 181, 524, 3, 2,
	2, 2, 183, 530, 3, 2, 2, 2, 185, 537, 3, 2, 2, 2, 187, 539, 3, 2, 2, 2,
	189, 555, 3, 2, 2, 2, 191, 569, 3, 2, 2, 2, 193, 581, 3, 2, 2, 2, 195,
	586, 3, 2, 2, 2, 197, 198, 7, 104, 2, 2, 198, 199, 7, 116, 2, 2, 199, 200,
	7, 113, 2, 2, 200, 201, 7, 111, 2, 2, 201, 4, 3, 2, 2, 2, 202, 203, 7,
	100, 2, 2, 203, 204, 7, 123, 2, 2, 204, 6, 3, 2, 2, 2, 205, 206, 7, 117,
	2, 2, 206, 207, 7, 103, 2, 2, 207, 208, 7, 110, 2, 2, 208, 209, 7, 103,
	2, 2, 209, 210, 7, 101, 2, 2, 210, 211, 7, 118, 2, 2, 211, 8, 3, 2, 2,
	2, 212, 213, 7, 121, 2, 2, 213, 214, 7, 106, 2, 2, 214, 215, 7, 103, 2,
	2, 215, 216, 7, 116, 2, 2, 216, 217, 7, 103, 2, 2, 217, 10, 3, 2, 2, 2,
	218, 219, 7, 105, 2, 2, 219, 220, 7, 116, 2, 2, 220, 221, 7, 113, 2, 2,
	221, 222, 7, 119, 2, 2, 222, 223, 7, 114, 2, 2, 223, 12, 3, 2, 2, 2, 224,
	225, 7, 107, 2, 2, 225, 226, 7, 112, 2, 2, 226, 227, 7, 118, 2, 2, 227,
	228, 7, 113, 2, 2, 228, 14, 3, 2, 2, 2, 229, 230, 7, 113, 2, 2, 230, 231,
	7, 116, 2, 2, 231, 232, 7, 102, 2, 2, 232, 233, 7, 103, 2, 2, 233, 234,
	7, 116, 2, 2, 234, 235, 7, 100, 2, 2, 235, 236, 7, 123, 2, 2, 236, 16,
	3, 2, 2, 2, 237, 238, 7, 108, 2, 2, 238, 239, 7, 113, 2, 2, 239, 240, 7,
	107, 2, 2, 240, 241, 7, 112, 2, 2, 241, 18, 3, 2, 2, 2, 242, 243, 7, 110,
	2, 2, 243, 244, 7, 103, 2, 2, 244, 245, 7, 118, 2, 2, 245, 20, 3, 2, 2,
	2, 246, 247, 7, 107, 2, 2, 247, 248, 7, 112, 2, 2, 248, 22, 3, 2, 2, 2,
	249, 250, 7, 113, 2, 2, 250, 251, 7, 112, 2, 2, 251, 24, 3, 2, 2, 2, 252,
	253, 7, 103, 2, 2, 253, 254, 7, 115, 2, 2, 254, 255, 7, 119, 2, 2, 255,
	256, 7, 99, 2, 2, 256, 257, 7, 110, 2, 2, 257, 258, 7, 117, 2, 2, 258,
	26, 3, 2, 2, 2, 259, 260, 7, 99, 2, 2, 260, 261, 7, 117, 2, 2, 261, 262,
	7, 101, 2, 2, 262, 263, 7, 103, 2, 2, 263, 264, 7, 112, 2, 2, 264, 265,
	7, 102, 2, 2, 265, 266, 7, 107, 2, 2, 266, 267, 7, 112, 2, 2, 267, 268,
	7, 105, 2, 2, 268, 28, 3, 2, 2, 2, 269, 270, 7, 102, 2, 2, 270, 271, 7,
	103, 2, 2, 271, 272, 7, 117, 2, 2, 272, 273, 7, 101, 2, 2, 273, 274, 7,
	103, 2, 2, 274, 275, 7, 112, 2, 2, 275, 276, 7, 102, 2, 2, 276, 277, 7,
	107, 2, 2, 277, 278, 7, 112, 2, 2, 278, 279, 7, 105, 2, 2, 279, 30, 3,
	2, 2, 2, 280, 281, 7, 44, 2, 2, 281, 282, 7, 96, 2, 2, 282, 32, 3, 2, 2,
	2, 283, 284, 7, 49, 2, 2, 284, 285, 7, 96, 2, 2, 285, 34, 3, 2, 2, 2, 286,
	287, 7, 94, 2, 2, 287, 288, 7, 96, 2, 2, 288, 36, 3, 2, 2, 2, 289, 290,
	7, 45, 2, 2, 290, 291, 7, 45, 2, 2, 291, 38, 3, 2, 2, 2, 292, 293, 7, 47,
	2, 2, 293, 294, 7, 47, 2, 2, 294, 40, 3, 2, 2, 2, 295, 296, 7, 45, 2, 2,
	296, 297, 7, 63, 2, 2, 297, 42, 3, 2, 2, 2, 298, 299, 7, 47, 2, 2, 299,
	300, 7, 63, 2, 2, 300, 44, 3, 2, 2, 2, 301, 302, 7, 44, 2, 2, 302, 303,
	7, 63, 2, 2, 303, 46, 3, 2, 2, 2, 304, 305, 7, 49, 2, 2, 305, 306, 7, 63,
	2, 2, 306, 48, 3, 2, 2, 2, 307, 308, 7, 94, 2, 2, 308, 309, 7, 63, 2, 2,
	309, 50, 3, 2, 2, 2, 310, 311, 7, 60, 2, 2, 311, 312, 7, 63, 2, 2, 312,
	52, 3, 2, 2, 2, 313, 314, 7, 63, 2, 2, 314, 315, 7, 63, 2, 2, 315, 54,
	3, 2, 2, 2, 316, 317, 7, 62, 2, 2, 317, 318, 7, 63, 2, 2, 318, 56, 3, 2,
	2, 2, 319, 320, 7, 64, 2, 2, 320, 321, 7, 63, 2, 2, 321, 58, 3, 2, 2, 2,
	322, 323, 7, 64, 2, 2, 323, 324, 7, 62, 2, 2, 324, 60, 3, 2, 2, 2, 325,
	326, 7, 48, 2, 2, 326, 327, 7, 48, 2, 2, 327, 328, 7, 48, 2, 2, 328, 62,
	3, 2, 2, 2, 329, 330, 7, 48, 2, 2, 330, 331, 7, 48, 2, 2, 331, 64, 3, 2,
	2, 2, 332, 333, 7, 48, 2, 2, 333, 66, 3, 2, 2, 2, 334, 335, 7, 46, 2, 2,
	335, 68, 3, 2, 2, 2, 336, 337, 7, 63, 2, 2, 337, 338, 7, 64, 2, 2, 338,
	70, 3, 2, 2, 2, 339, 340, 7, 47, 2, 2, 340, 341, 7, 64, 2, 2, 341, 72,
	3, 2, 2, 2, 342, 343, 7, 62, 2, 2, 343, 344, 7, 47, 2, 2, 344, 74, 3, 2,
	2, 2, 345, 346, 7, 128, 2, 2, 346, 347, 7, 64, 2, 2, 347, 76, 3, 2, 2,
	2, 348, 349, 7, 62, 2, 2, 349, 350, 7, 128, 2, 2, 350, 78, 3, 2, 2, 2,
	351, 352, 7, 63, 2, 2, 352, 80, 3, 2, 2, 2, 353, 354, 7, 62, 2, 2, 354,
	82, 3, 2, 2, 2, 355, 356, 7, 64, 2, 2, 356, 84, 3, 2, 2, 2, 357, 358, 7,
	61, 2, 2, 358, 86, 3, 2, 2, 2, 359, 360, 7, 42, 2, 2, 360, 88, 3, 2, 2,
	2, 361, 362, 7, 43, 2, 2, 362, 90, 3, 2, 2, 2, 363, 364, 7, 125, 2, 2,
	364, 92, 3, 2, 2, 2, 365, 366, 7, 127, 2, 2, 366, 94, 3, 2, 2, 2, 367,
	368, 7, 93, 2, 2, 368, 96, 3, 2, 2, 2, 369, 370, 7, 95, 2, 2, 370, 98,
	3, 2, 2, 2, 371, 372, 7, 60, 2, 2, 372, 373, 7, 60, 2, 2, 373, 100, 3,
	2, 2, 2, 374, 375, 7, 60, 2, 2, 375, 102, 3, 2, 2, 2, 376, 377, 7, 65,
	2, 2, 377, 104, 3, 2, 2, 2, 378, 379, 7, 66, 2, 2, 379, 106, 3, 2, 2, 2,
	380, 381, 7, 35, 2, 2, 381, 108, 3, 2, 2, 2, 382, 383, 7, 38, 2, 2, 383,
	110, 3, 2, 2, 2, 384, 385, 7, 39, 2, 2, 385, 112, 3, 2, 2, 2, 386, 387,
	7, 128, 2, 2, 387, 114, 3, 2, 2, 2, 388, 389, 7, 45, 2, 2, 389, 116, 3,
	2, 2, 2, 390, 391, 7, 47, 2, 2, 391, 118, 3, 2, 2, 2, 392, 393, 7, 44,
	2, 2, 393, 120, 3, 2, 2, 2, 394, 395, 7, 49, 2, 2, 395, 122, 3, 2, 2, 2,
	396, 397, 7, 94, 2, 2, 397, 124, 3, 2, 2, 2, 398, 399, 7, 40, 2, 2, 399,
	126, 3, 2, 2, 2, 400, 401, 7, 126, 2, 2, 401, 128, 3, 2, 2, 2, 402, 403,
	7, 96, 2, 2, 403, 130, 3, 2, 2, 2, 404, 405, 7, 107, 2, 2, 405, 406, 7,
	58, 2, 2, 406, 132, 3, 2, 2, 2, 407, 408, 7, 119, 2, 2, 408, 409, 7, 58,
	2, 2, 409, 134, 3, 2, 2, 2, 410, 411, 7, 107, 2, 2, 411, 412, 7, 51, 2,
	2, 412, 413, 7, 56, 2, 2, 413, 136, 3, 2, 2, 2, 414, 415, 7, 119, 2, 2,
	415, 416, 7, 51, 2, 2, 416, 417, 7, 56, 2, 2, 417, 138, 3, 2, 2, 2, 418,
	419, 7, 107, 2, 2, 419, 420, 7, 53, 2, 2, 420, 421, 7, 52, 2, 2, 421, 140,
	3, 2, 2, 2, 422, 423, 7, 119, 2, 2, 423, 424, 7, 53, 2, 2, 424, 425, 7,
	52, 2, 2, 425, 142, 3, 2, 2, 2, 426, 427, 7, 107, 2, 2, 427, 428, 7, 56,
	2, 2, 428, 429, 7, 54, 2, 2, 429, 144, 3, 2, 2, 2, 430, 431, 7, 119, 2,
	2, 431, 432, 7, 56, 2, 2, 432, 433, 7, 54, 2, 2, 433, 146, 3, 2, 2, 2,
	434, 435, 7, 104, 2, 2, 435, 436, 7, 53, 2, 2, 436, 437, 7, 52, 2, 2, 437,
	148, 3, 2, 2, 2, 438, 439, 7, 104, 2, 2, 439, 440, 7, 56, 2, 2, 440, 441,
	7, 54, 2, 2, 441, 150, 3, 2, 2, 2, 442, 443, 7, 101, 2, 2, 443, 444, 7,
	106, 2, 2, 444, 445, 7, 116, 2, 2, 445, 152, 3, 2, 2, 2, 446, 447, 7, 117,
	2, 2, 447, 448, 7, 118, 2, 2, 448, 449, 7, 116, 2, 2, 449, 154, 3, 2, 2,
	2, 450, 451, 7, 100, 2, 2, 451, 452, 7, 113, 2, 2, 452, 453, 7, 113, 2,
	2, 453, 454, 7, 110, 2, 2, 454, 156, 3, 2, 2, 2, 455, 456, 7, 107, 2, 2,
	456, 457, 7, 112, 2, 2, 457, 458, 7, 118, 2, 2, 458, 158, 3, 2, 2, 2, 459,
	460, 7, 112, 2, 2, 460, 461, 7, 119, 2, 2, 461, 462, 7, 111, 2, 2, 462,
	160, 3, 2, 2, 2, 463, 464, 7, 100, 2, 2, 464, 465, 7, 123, 2, 2, 465, 466,
	7, 118, 2, 2, 466, 467, 7, 103, 2, 2, 467, 162, 3, 2, 2, 2, 468, 469, 7,
	99, 2, 2, 469, 470, 7, 112, 2, 2, 470, 471, 7, 123, 2, 2, 471, 164, 3,
	2, 2, 2, 472, 473, 7, 112, 2, 2, 473, 474, 7, 107, 2, 2, 474, 475, 7, 110,
	2, 2, 475, 166, 3, 2, 2, 2, 476, 477, 7, 118, 2, 2, 477, 478, 7, 116, 2,
	2, 478, 479, 7, 119, 2, 2, 479, 480, 7, 103, 2, 2, 480, 168, 3, 2, 2, 2,
	481, 482, 7, 104, 2, 2, 482, 483, 7, 99, 2, 2, 483, 484, 7, 110, 2, 2,
	484, 485, 7, 117, 2, 2, 485, 486, 7, 103, 2, 2, 486, 170, 3, 2, 2, 2, 487,
	488, 7, 119, 2, 2, 488, 489, 7, 112, 2, 2, 489, 490, 7, 102, 2, 2, 490,
	491, 7, 103, 2, 2, 491, 492, 7, 104, 2, 2, 492, 172, 3, 2, 2, 2, 493, 495,
	5, 175, 88, 2, 494, 493, 3, 2, 2, 2, 495, 496, 3, 2, 2, 2, 496, 494, 3,
	2, 2, 2, 496, 497, 3, 2, 2, 2, 497, 174, 3, 2, 2, 2, 498, 499, 9, 2, 2,
	2, 499, 176, 3, 2, 2, 2, 500, 506, 7, 36, 2, 2, 501, 502, 7, 94, 2, 2,
	502, 505, 9, 3, 2, 2, 503, 505, 11, 2, 2, 2, 504, 501, 3, 2, 2, 2, 504,
	503, 3, 2, 2, 2, 505, 508, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 506, 504,
	3, 2, 2, 2, 507, 509, 3, 2, 2, 2, 508, 506, 3, 2, 2, 2, 509, 510, 7, 36,
	2, 2, 510, 178, 3, 2, 2, 2, 511, 519, 7, 41, 2, 2, 512, 513, 7, 94, 2,
	2, 513, 518, 7, 41, 2, 2, 514, 515, 7, 94, 2, 2, 515, 518, 9, 4, 2, 2,
	516, 518, 11, 2, 2, 2, 517, 512, 3, 2, 2, 2, 517, 514, 3, 2, 2, 2, 517,
	516, 3, 2, 2, 2, 518, 521, 3, 2, 2, 2, 519, 520, 3, 2, 2, 2, 519, 517,
	3, 2, 2, 2, 520, 522, 3, 2, 2, 2, 521, 519, 3, 2, 2, 2, 522, 523, 7, 41,
	2, 2, 523, 180, 3, 2, 2, 2, 524, 526, 7, 97, 2, 2, 525, 527, 9, 5, 2, 2,
	526, 525, 3, 2, 2, 2, 527, 528, 3, 2, 2, 2, 528, 526, 3, 2, 2, 2, 528,
	529, 3, 2, 2, 2, 529, 182, 3, 2, 2, 2, 530, 534, 9, 6, 2, 2, 531, 533,
	9, 5, 2, 2, 532, 531, 3, 2, 2, 2, 533, 536, 3, 2, 2, 2, 534, 532, 3, 2,
	2, 2, 534, 535, 3, 2, 2, 2, 535, 184, 3, 2, 2, 2, 536, 534, 3, 2, 2, 2,
	537, 538, 7, 97, 2, 2, 538, 186, 3, 2, 2, 2, 539, 540, 7, 37, 2, 2, 540,
	541, 7, 37, 2, 2, 541, 542, 7, 37, 2, 2, 542, 546, 3, 2, 2, 2, 543, 545,
	11, 2, 2, 2, 544, 543, 3, 2, 2, 2, 545, 548, 3, 2, 2, 2, 546, 547, 3, 2,
	2, 2, 546, 544, 3, 2, 2, 2, 547, 549, 3, 2, 2, 2, 548, 546, 3, 2, 2, 2,
	549, 550, 7, 37, 2, 2, 550, 551, 7, 37, 2, 2, 551, 552, 7, 37, 2, 2, 552,
	553, 3, 2, 2, 2, 553, 554, 8, 94, 2, 2, 554, 188, 3, 2, 2, 2, 555, 556,
	7, 37, 2, 2, 556, 557, 7, 37, 2, 2, 557, 561, 3, 2, 2, 2, 558, 560, 11,
	2, 2, 2, 559, 558, 3, 2, 2, 2, 560, 563, 3, 2, 2, 2, 561, 562, 3, 2, 2,
	2, 561, 559, 3, 2, 2, 2, 562, 564, 3, 2, 2, 2, 563, 561, 3, 2, 2, 2, 564,
	565, 7, 37, 2, 2, 565, 566, 7, 37, 2, 2, 566, 567, 3, 2, 2, 2, 567, 568,
	8, 95, 2, 2, 568, 190, 3, 2, 2, 2, 569, 573, 7, 37, 2, 2, 570, 572, 11,
	2, 2, 2, 571, 570, 3, 2, 2, 2, 572, 575, 3, 2, 2, 2, 573, 574, 3, 2, 2,
	2, 573, 571, 3, 2, 2, 2, 574, 576, 3, 2, 2, 2, 575, 573, 3, 2, 2, 2, 576,
	577, 7, 37, 2, 2, 577, 578, 3, 2, 2, 2, 578, 579, 8, 96, 2, 2, 579, 192,
	3, 2, 2, 2, 580, 582, 7, 15, 2, 2, 581, 580, 3, 2, 2, 2, 581, 582, 3, 2,
	2, 2, 582, 583, 3, 2, 2, 2, 583, 584, 7, 12, 2, 2, 584, 194, 3, 2, 2, 2,
	585, 587, 9, 7, 2, 2, 586, 585, 3, 2, 2, 2, 587, 588, 3, 2, 2, 2, 588,
	586, 3, 2, 2, 2, 588, 589, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 591,
	8, 98, 2, 2, 591, 196, 3, 2, 2, 2, 15, 2, 496, 504, 506, 517, 519, 528,
	534, 546, 561, 573, 581, 588, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'from'", "'by'", "'select'", "'where'", "'group'", "'into'", "'orderby'",
	"'join'", "'let'", "'in'", "'on'", "'equals'", "'ascending'", "'descending'",
	"'*^'", "'/^'", "'\\^'", "'++'", "'--'", "'+='", "'-='", "'*='", "'/='",
	"'\\='", "':='", "'=='", "'<='", "'>='", "'><'", "'...'", "'..'", "'.'",
	"','", "'=>'", "'->'", "'<-'", "'~>'", "'<~'", "'='", "'<'", "'>'", "';'",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "'::'", "':'", "'?'", "'@'",
	"'!'", "'$'", "'%'", "'~'", "'+'", "'-'", "'*'", "'/'", "'\\'", "'&'",
	"'|'", "'^'", "'i8'", "'u8'", "'i16'", "'u16'", "'i32'", "'u32'", "'i64'",
	"'u64'", "'f32'", "'f64'", "'chr'", "'str'", "'bool'", "'int'", "'num'",
	"'byte'", "'any'", "'nil'", "'true'", "'false'", "'undef'", "", "", "",
	"", "", "'_'",
}

var lexerSymbolicNames = []string{
	"", "LinqFrom", "LinqBy", "LinqSelect", "LinqWhere", "LinqGroup", "LinqInto",
	"LinqOrderby", "LinqJoin", "LinqLet", "LinqIn", "LinqOn", "LinqEquals",
	"LinqAscending", "LinqDescending", "Pow", "Root", "Log", "Add_Add", "Sub_Sub",
	"Add_Equal", "Sub_Equal", "Mul_Equal", "Div_Equal", "Mod_Equal", "Colon_Equal",
	"Equal_Equal", "Less_Equal", "Greater_Equal", "Not_Equal", "Dot_Dot_Dot",
	"Dot_Dot", "Dot", "Comma", "Equal_Arrow", "Right_Arrow", "Left_Arrow",
	"Right_Flow", "Left_Flow", "Equal", "Less", "Greater", "Semi", "Left_Paren",
	"Right_Paren", "Left_Brace", "Right_Brace", "Left_Brack", "Right_Brack",
	"Colon_Colon", "Colon", "Question", "At", "Bang", "Coin", "Cent", "Wave",
	"Add", "Sub", "Mul", "Div", "Mod", "And", "Or", "Xor", "TypeI8", "TypeU8",
	"TypeI16", "TypeU16", "TypeI32", "TypeU32", "TypeI64", "TypeU64", "TypeF32",
	"TypeF64", "TypeChr", "TypeStr", "TypeBool", "TypeInt", "TypeNum", "TypeByte",
	"TypeAny", "NilLiteral", "TrueLiteral", "FalseLiteral", "UndefinedLiteral",
	"NumberLiteral", "TextLiteral", "CharLiteral", "IDPrivate", "IDPublic",
	"Discard", "Big_Big_Comment", "Big_Comment", "Comment", "New_Line", "WS",
}

var lexerRuleNames = []string{
	"LinqFrom", "LinqBy", "LinqSelect", "LinqWhere", "LinqGroup", "LinqInto",
	"LinqOrderby", "LinqJoin", "LinqLet", "LinqIn", "LinqOn", "LinqEquals",
	"LinqAscending", "LinqDescending", "Pow", "Root", "Log", "Add_Add", "Sub_Sub",
	"Add_Equal", "Sub_Equal", "Mul_Equal", "Div_Equal", "Mod_Equal", "Colon_Equal",
	"Equal_Equal", "Less_Equal", "Greater_Equal", "Not_Equal", "Dot_Dot_Dot",
	"Dot_Dot", "Dot", "Comma", "Equal_Arrow", "Right_Arrow", "Left_Arrow",
	"Right_Flow", "Left_Flow", "Equal", "Less", "Greater", "Semi", "Left_Paren",
	"Right_Paren", "Left_Brace", "Right_Brace", "Left_Brack", "Right_Brack",
	"Colon_Colon", "Colon", "Question", "At", "Bang", "Coin", "Cent", "Wave",
	"Add", "Sub", "Mul", "Div", "Mod", "And", "Or", "Xor", "TypeI8", "TypeU8",
	"TypeI16", "TypeU16", "TypeI32", "TypeU32", "TypeI64", "TypeU64", "TypeF32",
	"TypeF64", "TypeChr", "TypeStr", "TypeBool", "TypeInt", "TypeNum", "TypeByte",
	"TypeAny", "NilLiteral", "TrueLiteral", "FalseLiteral", "UndefinedLiteral",
	"NumberLiteral", "DIGIT", "TextLiteral", "CharLiteral", "IDPrivate", "IDPublic",
	"Discard", "Big_Big_Comment", "Big_Comment", "Comment", "New_Line", "WS",
}

type LiteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewLiteLexer(input antlr.CharStream) *LiteLexer {

	l := new(LiteLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "LiteLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LiteLexer tokens.
const (
	LiteLexerLinqFrom         = 1
	LiteLexerLinqBy           = 2
	LiteLexerLinqSelect       = 3
	LiteLexerLinqWhere        = 4
	LiteLexerLinqGroup        = 5
	LiteLexerLinqInto         = 6
	LiteLexerLinqOrderby      = 7
	LiteLexerLinqJoin         = 8
	LiteLexerLinqLet          = 9
	LiteLexerLinqIn           = 10
	LiteLexerLinqOn           = 11
	LiteLexerLinqEquals       = 12
	LiteLexerLinqAscending    = 13
	LiteLexerLinqDescending   = 14
	LiteLexerPow              = 15
	LiteLexerRoot             = 16
	LiteLexerLog              = 17
	LiteLexerAdd_Add          = 18
	LiteLexerSub_Sub          = 19
	LiteLexerAdd_Equal        = 20
	LiteLexerSub_Equal        = 21
	LiteLexerMul_Equal        = 22
	LiteLexerDiv_Equal        = 23
	LiteLexerMod_Equal        = 24
	LiteLexerColon_Equal      = 25
	LiteLexerEqual_Equal      = 26
	LiteLexerLess_Equal       = 27
	LiteLexerGreater_Equal    = 28
	LiteLexerNot_Equal        = 29
	LiteLexerDot_Dot_Dot      = 30
	LiteLexerDot_Dot          = 31
	LiteLexerDot              = 32
	LiteLexerComma            = 33
	LiteLexerEqual_Arrow      = 34
	LiteLexerRight_Arrow      = 35
	LiteLexerLeft_Arrow       = 36
	LiteLexerRight_Flow       = 37
	LiteLexerLeft_Flow        = 38
	LiteLexerEqual            = 39
	LiteLexerLess             = 40
	LiteLexerGreater          = 41
	LiteLexerSemi             = 42
	LiteLexerLeft_Paren       = 43
	LiteLexerRight_Paren      = 44
	LiteLexerLeft_Brace       = 45
	LiteLexerRight_Brace      = 46
	LiteLexerLeft_Brack       = 47
	LiteLexerRight_Brack      = 48
	LiteLexerColon_Colon      = 49
	LiteLexerColon            = 50
	LiteLexerQuestion         = 51
	LiteLexerAt               = 52
	LiteLexerBang             = 53
	LiteLexerCoin             = 54
	LiteLexerCent             = 55
	LiteLexerWave             = 56
	LiteLexerAdd              = 57
	LiteLexerSub              = 58
	LiteLexerMul              = 59
	LiteLexerDiv              = 60
	LiteLexerMod              = 61
	LiteLexerAnd              = 62
	LiteLexerOr               = 63
	LiteLexerXor              = 64
	LiteLexerTypeI8           = 65
	LiteLexerTypeU8           = 66
	LiteLexerTypeI16          = 67
	LiteLexerTypeU16          = 68
	LiteLexerTypeI32          = 69
	LiteLexerTypeU32          = 70
	LiteLexerTypeI64          = 71
	LiteLexerTypeU64          = 72
	LiteLexerTypeF32          = 73
	LiteLexerTypeF64          = 74
	LiteLexerTypeChr          = 75
	LiteLexerTypeStr          = 76
	LiteLexerTypeBool         = 77
	LiteLexerTypeInt          = 78
	LiteLexerTypeNum          = 79
	LiteLexerTypeByte         = 80
	LiteLexerTypeAny          = 81
	LiteLexerNilLiteral       = 82
	LiteLexerTrueLiteral      = 83
	LiteLexerFalseLiteral     = 84
	LiteLexerUndefinedLiteral = 85
	LiteLexerNumberLiteral    = 86
	LiteLexerTextLiteral      = 87
	LiteLexerCharLiteral      = 88
	LiteLexerIDPrivate        = 89
	LiteLexerIDPublic         = 90
	LiteLexerDiscard          = 91
	LiteLexerBig_Big_Comment  = 92
	LiteLexerBig_Comment      = 93
	LiteLexerComment          = 94
	LiteLexerNew_Line         = 95
	LiteLexerWS               = 96
)
