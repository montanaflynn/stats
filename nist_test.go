package stats_test

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

var (
	lew = stats.Float64Data{
		-213, -564, -35, -15, 141, 115, -420, -360, 203, -338, -431, 194,
		-220, -513, 154, -125, -559, 92, -21, -579, -52, 99, -543, -175,
		162, -457, -346, 204, -300, -474, 164, -107, -572, -8, 83, -541,
		-224, 180, -420, -374, 201, -236, -531, 83, 27, -564, -112, 131,
		-507, -254, 199, -311, -495, 143, -46, -579, -90, 136, -472, -338,
		202, -287, -477, 169, -124, -568, 17, 48, -568, -135, 162, -430,
		-422, 172, -74, -577, -13, 92, -534, -243, 194, -355, -465, 156,
		-81, -578, -64, 139, -449, -384, 193, -198, -538, 110, -44, -577,
		-6, 66, -552, -164, 161, -460, -344, 205, -281, -504, 134, -28,
		-576, -118, 156, -437, -381, 200, -220, -540, 83, 11, -568, -160,
		172, -414, -408, 188, -125, -572, -32, 139, -492, -321, 205, -262,
		-504, 142, -83, -574, 0, 48, -571, -106, 137, -501, -266, 190,
		-391, -406, 194, -186, -553, 83, -13, -577, -49, 103, -515, -280,
		201, 300, -506, 131, -45, -578, -80, 138, -462, -361, 201, -211,
		-554, 32, 74, -533, -235, 187, -372, -442, 182, -147, -566, 25,
		68, -535, -244, 194, -351, -463, 174, -125, -570, 15, 72, -550,
		-190, 172, -424, -385, 198, -218, -536, 96}

	lottery = stats.Float64Data{
		162, 671, 933, 414, 788, 730, 817, 33, 536, 875, 670, 236, 473, 167,
		877, 980, 316, 950, 456, 92, 517, 557, 956, 954, 104, 178, 794, 278,
		147, 773, 437, 435, 502, 610, 582, 780, 689, 562, 964, 791, 28, 97,
		848, 281, 858, 538, 660, 972, 671, 613, 867, 448, 738, 966, 139, 636,
		847, 659, 754, 243, 122, 455, 195, 968, 793, 59, 730, 361, 574, 522,
		97, 762, 431, 158, 429, 414, 22, 629, 788, 999, 187, 215, 810, 782,
		47, 34, 108, 986, 25, 644, 829, 630, 315, 567, 919, 331, 207, 412,
		242, 607, 668, 944, 749, 168, 864, 442, 533, 805, 372, 63, 458, 777,
		416, 340, 436, 140, 919, 350, 510, 572, 905, 900, 85, 389, 473, 758,
		444, 169, 625, 692, 140, 897, 672, 288, 312, 860, 724, 226, 884, 508,
		976, 741, 476, 417, 831, 15, 318, 432, 241, 114, 799, 955, 833, 358,
		935, 146, 630, 830, 440, 642, 356, 373, 271, 715, 367, 393, 190, 669,
		8, 861, 108, 795, 269, 590, 326, 866, 64, 523, 862, 840, 219, 382,
		998, 4, 628, 305, 747, 247, 34, 747, 729, 645, 856, 974, 24, 568, 24,
		694, 608, 480, 410, 729, 947, 293, 53, 930, 223, 203, 677, 227, 62,
		455, 387, 318, 562, 242, 428, 968}

	mavro = stats.Float64Data{
		2.00180, 2.00170, 2.00180, 2.00190, 2.00180, 2.00170, 2.00150,
		2.00140, 2.00150, 2.00150, 2.00170, 2.00180, 2.00180, 2.00190,
		2.00190, 2.00210, 2.00200, 2.00160, 2.00140, 2.00130, 2.00130,
		2.00150, 2.00150, 2.00160, 2.00150, 2.00140, 2.00130, 2.00140,
		2.00150, 2.00140, 2.00150, 2.00160, 2.00150, 2.00160, 2.00190,
		2.00200, 2.00200, 2.00210, 2.00220, 2.00230, 2.00240, 2.00250,
		2.00270, 2.00260, 2.00260, 2.00260, 2.00270, 2.00260, 2.00250,
		2.00240}

	michelson = stats.Float64Data{
		299.85, 299.74, 299.90, 300.07, 299.93, 299.85, 299.95, 299.98,
		299.98, 299.88, 300.00, 299.98, 299.93, 299.65, 299.76, 299.81,
		300.00, 300.00, 299.96, 299.96, 299.96, 299.94, 299.96, 299.94,
		299.88, 299.80, 299.85, 299.88, 299.90, 299.84, 299.83, 299.79,
		299.81, 299.88, 299.88, 299.83, 299.80, 299.79, 299.76, 299.80,
		299.88, 299.88, 299.88, 299.86, 299.72, 299.72, 299.62, 299.86,
		299.97, 299.95, 299.88, 299.91, 299.85, 299.87, 299.84, 299.84,
		299.85, 299.84, 299.84, 299.84, 299.89, 299.81, 299.81, 299.82,
		299.80, 299.77, 299.76, 299.74, 299.75, 299.76, 299.91, 299.92,
		299.89, 299.86, 299.88, 299.72, 299.84, 299.85, 299.85, 299.78,
		299.89, 299.84, 299.78, 299.81, 299.76, 299.81, 299.79, 299.81,
		299.82, 299.85, 299.87, 299.87, 299.81, 299.74, 299.81, 299.94,
		299.95, 299.80, 299.81, 299.87}

	pidigits = stats.Float64Data{
		3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8,
		9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 2, 7, 9, 5, 0, 2,
		8, 8, 4, 1, 9, 7, 1, 6, 9, 3, 9, 9, 3, 7, 5, 1, 0, 5, 8, 2, 0, 9,
		7, 4, 9, 4, 4, 5, 9, 2, 3, 0, 7, 8, 1, 6, 4, 0, 6, 2, 8, 6, 2, 0,
		8, 9, 9, 8, 6, 2, 8, 0, 3, 4, 8, 2, 5, 3, 4, 2, 1, 1, 7, 0, 6, 7,
		9, 8, 2, 1, 4, 8, 0, 8, 6, 5, 1, 3, 2, 8, 2, 3, 0, 6, 6, 4, 7, 0,
		9, 3, 8, 4, 4, 6, 0, 9, 5, 5, 0, 5, 8, 2, 2, 3, 1, 7, 2, 5, 3, 5,
		9, 4, 0, 8, 1, 2, 8, 4, 8, 1, 1, 1, 7, 4, 5, 0, 2, 8, 4, 1, 0, 2,
		7, 0, 1, 9, 3, 8, 5, 2, 1, 1, 0, 5, 5, 5, 9, 6, 4, 4, 6, 2, 2, 9,
		4, 8, 9, 5, 4, 9, 3, 0, 3, 8, 1, 9, 6, 4, 4, 2, 8, 8, 1, 0, 9, 7,
		5, 6, 6, 5, 9, 3, 3, 4, 4, 6, 1, 2, 8, 4, 7, 5, 6, 4, 8, 2, 3, 3,
		7, 8, 6, 7, 8, 3, 1, 6, 5, 2, 7, 1, 2, 0, 1, 9, 0, 9, 1, 4, 5, 6,
		4, 8, 5, 6, 6, 9, 2, 3, 4, 6, 0, 3, 4, 8, 6, 1, 0, 4, 5, 4, 3, 2,
		6, 6, 4, 8, 2, 1, 3, 3, 9, 3, 6, 0, 7, 2, 6, 0, 2, 4, 9, 1, 4, 1,
		2, 7, 3, 7, 2, 4, 5, 8, 7, 0, 0, 6, 6, 0, 6, 3, 1, 5, 5, 8, 8, 1,
		7, 4, 8, 8, 1, 5, 2, 0, 9, 2, 0, 9, 6, 2, 8, 2, 9, 2, 5, 4, 0, 9,
		1, 7, 1, 5, 3, 6, 4, 3, 6, 7, 8, 9, 2, 5, 9, 0, 3, 6, 0, 0, 1, 1,
		3, 3, 0, 5, 3, 0, 5, 4, 8, 8, 2, 0, 4, 6, 6, 5, 2, 1, 3, 8, 4, 1,
		4, 6, 9, 5, 1, 9, 4, 1, 5, 1, 1, 6, 0, 9, 4, 3, 3, 0, 5, 7, 2, 7,
		0, 3, 6, 5, 7, 5, 9, 5, 9, 1, 9, 5, 3, 0, 9, 2, 1, 8, 6, 1, 1, 7,
		3, 8, 1, 9, 3, 2, 6, 1, 1, 7, 9, 3, 1, 0, 5, 1, 1, 8, 5, 4, 8, 0,
		7, 4, 4, 6, 2, 3, 7, 9, 9, 6, 2, 7, 4, 9, 5, 6, 7, 3, 5, 1, 8, 8,
		5, 7, 5, 2, 7, 2, 4, 8, 9, 1, 2, 2, 7, 9, 3, 8, 1, 8, 3, 0, 1, 1,
		9, 4, 9, 1, 2, 9, 8, 3, 3, 6, 7, 3, 3, 6, 2, 4, 4, 0, 6, 5, 6, 6,
		4, 3, 0, 8, 6, 0, 2, 1, 3, 9, 4, 9, 4, 6, 3, 9, 5, 2, 2, 4, 7, 3,
		7, 1, 9, 0, 7, 0, 2, 1, 7, 9, 8, 6, 0, 9, 4, 3, 7, 0, 2, 7, 7, 0,
		5, 3, 9, 2, 1, 7, 1, 7, 6, 2, 9, 3, 1, 7, 6, 7, 5, 2, 3, 8, 4, 6,
		7, 4, 8, 1, 8, 4, 6, 7, 6, 6, 9, 4, 0, 5, 1, 3, 2, 0, 0, 0, 5, 6,
		8, 1, 2, 7, 1, 4, 5, 2, 6, 3, 5, 6, 0, 8, 2, 7, 7, 8, 5, 7, 7, 1,
		3, 4, 2, 7, 5, 7, 7, 8, 9, 6, 0, 9, 1, 7, 3, 6, 3, 7, 1, 7, 8, 7,
		2, 1, 4, 6, 8, 4, 4, 0, 9, 0, 1, 2, 2, 4, 9, 5, 3, 4, 3, 0, 1, 4,
		6, 5, 4, 9, 5, 8, 5, 3, 7, 1, 0, 5, 0, 7, 9, 2, 2, 7, 9, 6, 8, 9,
		2, 5, 8, 9, 2, 3, 5, 4, 2, 0, 1, 9, 9, 5, 6, 1, 1, 2, 1, 2, 9, 0,
		2, 1, 9, 6, 0, 8, 6, 4, 0, 3, 4, 4, 1, 8, 1, 5, 9, 8, 1, 3, 6, 2,
		9, 7, 7, 4, 7, 7, 1, 3, 0, 9, 9, 6, 0, 5, 1, 8, 7, 0, 7, 2, 1, 1,
		3, 4, 9, 9, 9, 9, 9, 9, 8, 3, 7, 2, 9, 7, 8, 0, 4, 9, 9, 5, 1, 0,
		5, 9, 7, 3, 1, 7, 3, 2, 8, 1, 6, 0, 9, 6, 3, 1, 8, 5, 9, 5, 0, 2,
		4, 4, 5, 9, 4, 5, 5, 3, 4, 6, 9, 0, 8, 3, 0, 2, 6, 4, 2, 5, 2, 2,
		3, 0, 8, 2, 5, 3, 3, 4, 4, 6, 8, 5, 0, 3, 5, 2, 6, 1, 9, 3, 1, 1,
		8, 8, 1, 7, 1, 0, 1, 0, 0, 0, 3, 1, 3, 7, 8, 3, 8, 7, 5, 2, 8, 8,
		6, 5, 8, 7, 5, 3, 3, 2, 0, 8, 3, 8, 1, 4, 2, 0, 6, 1, 7, 1, 7, 7,
		6, 6, 9, 1, 4, 7, 3, 0, 3, 5, 9, 8, 2, 5, 3, 4, 9, 0, 4, 2, 8, 7,
		5, 5, 4, 6, 8, 7, 3, 1, 1, 5, 9, 5, 6, 2, 8, 6, 3, 8, 8, 2, 3, 5,
		3, 7, 8, 7, 5, 9, 3, 7, 5, 1, 9, 5, 7, 7, 8, 1, 8, 5, 7, 7, 3, 0,
		5, 3, 2, 1, 7, 1, 2, 2, 6, 8, 0, 6, 6, 1, 3, 0, 0, 1, 9, 2, 7, 8,
		7, 6, 6, 1, 1, 1, 9, 5, 9, 0, 9, 2, 1, 6, 4, 2, 0, 1, 9, 8, 9, 3,
		8, 0, 9, 5, 2, 5, 7, 2, 0, 1, 0, 6, 5, 4, 8, 5, 8, 6, 3, 2, 7, 8,
		8, 6, 5, 9, 3, 6, 1, 5, 3, 3, 8, 1, 8, 2, 7, 9, 6, 8, 2, 3, 0, 3,
		0, 1, 9, 5, 2, 0, 3, 5, 3, 0, 1, 8, 5, 2, 9, 6, 8, 9, 9, 5, 7, 7,
		3, 6, 2, 2, 5, 9, 9, 4, 1, 3, 8, 9, 1, 2, 4, 9, 7, 2, 1, 7, 7, 5,
		2, 8, 3, 4, 7, 9, 1, 3, 1, 5, 1, 5, 5, 7, 4, 8, 5, 7, 2, 4, 2, 4,
		5, 4, 1, 5, 0, 6, 9, 5, 9, 5, 0, 8, 2, 9, 5, 3, 3, 1, 1, 6, 8, 6,
		1, 7, 2, 7, 8, 5, 5, 8, 8, 9, 0, 7, 5, 0, 9, 8, 3, 8, 1, 7, 5, 4,
		6, 3, 7, 4, 6, 4, 9, 3, 9, 3, 1, 9, 2, 5, 5, 0, 6, 0, 4, 0, 0, 9,
		2, 7, 7, 0, 1, 6, 7, 1, 1, 3, 9, 0, 0, 9, 8, 4, 8, 8, 2, 4, 0, 1,
		2, 8, 5, 8, 3, 6, 1, 6, 0, 3, 5, 6, 3, 7, 0, 7, 6, 6, 0, 1, 0, 4,
		7, 1, 0, 1, 8, 1, 9, 4, 2, 9, 5, 5, 5, 9, 6, 1, 9, 8, 9, 4, 6, 7,
		6, 7, 8, 3, 7, 4, 4, 9, 4, 4, 8, 2, 5, 5, 3, 7, 9, 7, 7, 4, 7, 2,
		6, 8, 4, 7, 1, 0, 4, 0, 4, 7, 5, 3, 4, 6, 4, 6, 2, 0, 8, 0, 4, 6,
		6, 8, 4, 2, 5, 9, 0, 6, 9, 4, 9, 1, 2, 9, 3, 3, 1, 3, 6, 7, 7, 0,
		2, 8, 9, 8, 9, 1, 5, 2, 1, 0, 4, 7, 5, 2, 1, 6, 2, 0, 5, 6, 9, 6,
		6, 0, 2, 4, 0, 5, 8, 0, 3, 8, 1, 5, 0, 1, 9, 3, 5, 1, 1, 2, 5, 3,
		3, 8, 2, 4, 3, 0, 0, 3, 5, 5, 8, 7, 6, 4, 0, 2, 4, 7, 4, 9, 6, 4,
		7, 3, 2, 6, 3, 9, 1, 4, 1, 9, 9, 2, 7, 2, 6, 0, 4, 2, 6, 9, 9, 2,
		2, 7, 9, 6, 7, 8, 2, 3, 5, 4, 7, 8, 1, 6, 3, 6, 0, 0, 9, 3, 4, 1,
		7, 2, 1, 6, 4, 1, 2, 1, 9, 9, 2, 4, 5, 8, 6, 3, 1, 5, 0, 3, 0, 2,
		8, 6, 1, 8, 2, 9, 7, 4, 5, 5, 5, 7, 0, 6, 7, 4, 9, 8, 3, 8, 5, 0,
		5, 4, 9, 4, 5, 8, 8, 5, 8, 6, 9, 2, 6, 9, 9, 5, 6, 9, 0, 9, 2, 7,
		2, 1, 0, 7, 9, 7, 5, 0, 9, 3, 0, 2, 9, 5, 5, 3, 2, 1, 1, 6, 5, 3,
		4, 4, 9, 8, 7, 2, 0, 2, 7, 5, 5, 9, 6, 0, 2, 3, 6, 4, 8, 0, 6, 6,
		5, 4, 9, 9, 1, 1, 9, 8, 8, 1, 8, 3, 4, 7, 9, 7, 7, 5, 3, 5, 6, 6,
		3, 6, 9, 8, 0, 7, 4, 2, 6, 5, 4, 2, 5, 2, 7, 8, 6, 2, 5, 5, 1, 8,
		1, 8, 4, 1, 7, 5, 7, 4, 6, 7, 2, 8, 9, 0, 9, 7, 7, 7, 7, 2, 7, 9,
		3, 8, 0, 0, 0, 8, 1, 6, 4, 7, 0, 6, 0, 0, 1, 6, 1, 4, 5, 2, 4, 9,
		1, 9, 2, 1, 7, 3, 2, 1, 7, 2, 1, 4, 7, 7, 2, 3, 5, 0, 1, 4, 1, 4,
		4, 1, 9, 7, 3, 5, 6, 8, 5, 4, 8, 1, 6, 1, 3, 6, 1, 1, 5, 7, 3, 5,
		2, 5, 5, 2, 1, 3, 3, 4, 7, 5, 7, 4, 1, 8, 4, 9, 4, 6, 8, 4, 3, 8,
		5, 2, 3, 3, 2, 3, 9, 0, 7, 3, 9, 4, 1, 4, 3, 3, 3, 4, 5, 4, 7, 7,
		6, 2, 4, 1, 6, 8, 6, 2, 5, 1, 8, 9, 8, 3, 5, 6, 9, 4, 8, 5, 5, 6,
		2, 0, 9, 9, 2, 1, 9, 2, 2, 2, 1, 8, 4, 2, 7, 2, 5, 5, 0, 2, 5, 4,
		2, 5, 6, 8, 8, 7, 6, 7, 1, 7, 9, 0, 4, 9, 4, 6, 0, 1, 6, 5, 3, 4,
		6, 6, 8, 0, 4, 9, 8, 8, 6, 2, 7, 2, 3, 2, 7, 9, 1, 7, 8, 6, 0, 8,
		5, 7, 8, 4, 3, 8, 3, 8, 2, 7, 9, 6, 7, 9, 7, 6, 6, 8, 1, 4, 5, 4,
		1, 0, 0, 9, 5, 3, 8, 8, 3, 7, 8, 6, 3, 6, 0, 9, 5, 0, 6, 8, 0, 0,
		6, 4, 2, 2, 5, 1, 2, 5, 2, 0, 5, 1, 1, 7, 3, 9, 2, 9, 8, 4, 8, 9,
		6, 0, 8, 4, 1, 2, 8, 4, 8, 8, 6, 2, 6, 9, 4, 5, 6, 0, 4, 2, 4, 1,
		9, 6, 5, 2, 8, 5, 0, 2, 2, 2, 1, 0, 6, 6, 1, 1, 8, 6, 3, 0, 6, 7,
		4, 4, 2, 7, 8, 6, 2, 2, 0, 3, 9, 1, 9, 4, 9, 4, 5, 0, 4, 7, 1, 2,
		3, 7, 1, 3, 7, 8, 6, 9, 6, 0, 9, 5, 6, 3, 6, 4, 3, 7, 1, 9, 1, 7,
		2, 8, 7, 4, 6, 7, 7, 6, 4, 6, 5, 7, 5, 7, 3, 9, 6, 2, 4, 1, 3, 8,
		9, 0, 8, 6, 5, 8, 3, 2, 6, 4, 5, 9, 9, 5, 8, 1, 3, 3, 9, 0, 4, 7,
		8, 0, 2, 7, 5, 9, 0, 0, 9, 9, 4, 6, 5, 7, 6, 4, 0, 7, 8, 9, 5, 1,
		2, 6, 9, 4, 6, 8, 3, 9, 8, 3, 5, 2, 5, 9, 5, 7, 0, 9, 8, 2, 5, 8,
		2, 2, 6, 2, 0, 5, 2, 2, 4, 8, 9, 4, 0, 7, 7, 2, 6, 7, 1, 9, 4, 7,
		8, 2, 6, 8, 4, 8, 2, 6, 0, 1, 4, 7, 6, 9, 9, 0, 9, 0, 2, 6, 4, 0,
		1, 3, 6, 3, 9, 4, 4, 3, 7, 4, 5, 5, 3, 0, 5, 0, 6, 8, 2, 0, 3, 4,
		9, 6, 2, 5, 2, 4, 5, 1, 7, 4, 9, 3, 9, 9, 6, 5, 1, 4, 3, 1, 4, 2,
		9, 8, 0, 9, 1, 9, 0, 6, 5, 9, 2, 5, 0, 9, 3, 7, 2, 2, 1, 6, 9, 6,
		4, 6, 1, 5, 1, 5, 7, 0, 9, 8, 5, 8, 3, 8, 7, 4, 1, 0, 5, 9, 7, 8,
		8, 5, 9, 5, 9, 7, 7, 2, 9, 7, 5, 4, 9, 8, 9, 3, 0, 1, 6, 1, 7, 5,
		3, 9, 2, 8, 4, 6, 8, 1, 3, 8, 2, 6, 8, 6, 8, 3, 8, 6, 8, 9, 4, 2,
		7, 7, 4, 1, 5, 5, 9, 9, 1, 8, 5, 5, 9, 2, 5, 2, 4, 5, 9, 5, 3, 9,
		5, 9, 4, 3, 1, 0, 4, 9, 9, 7, 2, 5, 2, 4, 6, 8, 0, 8, 4, 5, 9, 8,
		7, 2, 7, 3, 6, 4, 4, 6, 9, 5, 8, 4, 8, 6, 5, 3, 8, 3, 6, 7, 3, 6,
		2, 2, 2, 6, 2, 6, 0, 9, 9, 1, 2, 4, 6, 0, 8, 0, 5, 1, 2, 4, 3, 8,
		8, 4, 3, 9, 0, 4, 5, 1, 2, 4, 4, 1, 3, 6, 5, 4, 9, 7, 6, 2, 7, 8,
		0, 7, 9, 7, 7, 1, 5, 6, 9, 1, 4, 3, 5, 9, 9, 7, 7, 0, 0, 1, 2, 9,
		6, 1, 6, 0, 8, 9, 4, 4, 1, 6, 9, 4, 8, 6, 8, 5, 5, 5, 8, 4, 8, 4,
		0, 6, 3, 5, 3, 4, 2, 2, 0, 7, 2, 2, 2, 5, 8, 2, 8, 4, 8, 8, 6, 4,
		8, 1, 5, 8, 4, 5, 6, 0, 2, 8, 5, 0, 6, 0, 1, 6, 8, 4, 2, 7, 3, 9,
		4, 5, 2, 2, 6, 7, 4, 6, 7, 6, 7, 8, 8, 9, 5, 2, 5, 2, 1, 3, 8, 5,
		2, 2, 5, 4, 9, 9, 5, 4, 6, 6, 6, 7, 2, 7, 8, 2, 3, 9, 8, 6, 4, 5,
		6, 5, 9, 6, 1, 1, 6, 3, 5, 4, 8, 8, 6, 2, 3, 0, 5, 7, 7, 4, 5, 6,
		4, 9, 8, 0, 3, 5, 5, 9, 3, 6, 3, 4, 5, 6, 8, 1, 7, 4, 3, 2, 4, 1,
		1, 2, 5, 1, 5, 0, 7, 6, 0, 6, 9, 4, 7, 9, 4, 5, 1, 0, 9, 6, 5, 9,
		6, 0, 9, 4, 0, 2, 5, 2, 2, 8, 8, 7, 9, 7, 1, 0, 8, 9, 3, 1, 4, 5,
		6, 6, 9, 1, 3, 6, 8, 6, 7, 2, 2, 8, 7, 4, 8, 9, 4, 0, 5, 6, 0, 1,
		0, 1, 5, 0, 3, 3, 0, 8, 6, 1, 7, 9, 2, 8, 6, 8, 0, 9, 2, 0, 8, 7,
		4, 7, 6, 0, 9, 1, 7, 8, 2, 4, 9, 3, 8, 5, 8, 9, 0, 0, 9, 7, 1, 4,
		9, 0, 9, 6, 7, 5, 9, 8, 5, 2, 6, 1, 3, 6, 5, 5, 4, 9, 7, 8, 1, 8,
		9, 3, 1, 2, 9, 7, 8, 4, 8, 2, 1, 6, 8, 2, 9, 9, 8, 9, 4, 8, 7, 2,
		2, 6, 5, 8, 8, 0, 4, 8, 5, 7, 5, 6, 4, 0, 1, 4, 2, 7, 0, 4, 7, 7,
		5, 5, 5, 1, 3, 2, 3, 7, 9, 6, 4, 1, 4, 5, 1, 5, 2, 3, 7, 4, 6, 2,
		3, 4, 3, 6, 4, 5, 4, 2, 8, 5, 8, 4, 4, 4, 7, 9, 5, 2, 6, 5, 8, 6,
		7, 8, 2, 1, 0, 5, 1, 1, 4, 1, 3, 5, 4, 7, 3, 5, 7, 3, 9, 5, 2, 3,
		1, 1, 3, 4, 2, 7, 1, 6, 6, 1, 0, 2, 1, 3, 5, 9, 6, 9, 5, 3, 6, 2,
		3, 1, 4, 4, 2, 9, 5, 2, 4, 8, 4, 9, 3, 7, 1, 8, 7, 1, 1, 0, 1, 4,
		5, 7, 6, 5, 4, 0, 3, 5, 9, 0, 2, 7, 9, 9, 3, 4, 4, 0, 3, 7, 4, 2,
		0, 0, 7, 3, 1, 0, 5, 7, 8, 5, 3, 9, 0, 6, 2, 1, 9, 8, 3, 8, 7, 4,
		4, 7, 8, 0, 8, 4, 7, 8, 4, 8, 9, 6, 8, 3, 3, 2, 1, 4, 4, 5, 7, 1,
		3, 8, 6, 8, 7, 5, 1, 9, 4, 3, 5, 0, 6, 4, 3, 0, 2, 1, 8, 4, 5, 3,
		1, 9, 1, 0, 4, 8, 4, 8, 1, 0, 0, 5, 3, 7, 0, 6, 1, 4, 6, 8, 0, 6,
		7, 4, 9, 1, 9, 2, 7, 8, 1, 9, 1, 1, 9, 7, 9, 3, 9, 9, 5, 2, 0, 6,
		1, 4, 1, 9, 6, 6, 3, 4, 2, 8, 7, 5, 4, 4, 4, 0, 6, 4, 3, 7, 4, 5,
		1, 2, 3, 7, 1, 8, 1, 9, 2, 1, 7, 9, 9, 9, 8, 3, 9, 1, 0, 1, 5, 9,
		1, 9, 5, 6, 1, 8, 1, 4, 6, 7, 5, 1, 4, 2, 6, 9, 1, 2, 3, 9, 7, 4,
		8, 9, 4, 0, 9, 0, 7, 1, 8, 6, 4, 9, 4, 2, 3, 1, 9, 6, 1, 5, 6, 7,
		9, 4, 5, 2, 0, 8, 0, 9, 5, 1, 4, 6, 5, 5, 0, 2, 2, 5, 2, 3, 1, 6,
		0, 3, 8, 8, 1, 9, 3, 0, 1, 4, 2, 0, 9, 3, 7, 6, 2, 1, 3, 7, 8, 5,
		5, 9, 5, 6, 6, 3, 8, 9, 3, 7, 7, 8, 7, 0, 8, 3, 0, 3, 9, 0, 6, 9,
		7, 9, 2, 0, 7, 7, 3, 4, 6, 7, 2, 2, 1, 8, 2, 5, 6, 2, 5, 9, 9, 6,
		6, 1, 5, 0, 1, 4, 2, 1, 5, 0, 3, 0, 6, 8, 0, 3, 8, 4, 4, 7, 7, 3,
		4, 5, 4, 9, 2, 0, 2, 6, 0, 5, 4, 1, 4, 6, 6, 5, 9, 2, 5, 2, 0, 1,
		4, 9, 7, 4, 4, 2, 8, 5, 0, 7, 3, 2, 5, 1, 8, 6, 6, 6, 0, 0, 2, 1,
		3, 2, 4, 3, 4, 0, 8, 8, 1, 9, 0, 7, 1, 0, 4, 8, 6, 3, 3, 1, 7, 3,
		4, 6, 4, 9, 6, 5, 1, 4, 5, 3, 9, 0, 5, 7, 9, 6, 2, 6, 8, 5, 6, 1,
		0, 0, 5, 5, 0, 8, 1, 0, 6, 6, 5, 8, 7, 9, 6, 9, 9, 8, 1, 6, 3, 5,
		7, 4, 7, 3, 6, 3, 8, 4, 0, 5, 2, 5, 7, 1, 4, 5, 9, 1, 0, 2, 8, 9,
		7, 0, 6, 4, 1, 4, 0, 1, 1, 0, 9, 7, 1, 2, 0, 6, 2, 8, 0, 4, 3, 9,
		0, 3, 9, 7, 5, 9, 5, 1, 5, 6, 7, 7, 1, 5, 7, 7, 0, 0, 4, 2, 0, 3,
		3, 7, 8, 6, 9, 9, 3, 6, 0, 0, 7, 2, 3, 0, 5, 5, 8, 7, 6, 3, 1, 7,
		6, 3, 5, 9, 4, 2, 1, 8, 7, 3, 1, 2, 5, 1, 4, 7, 1, 2, 0, 5, 3, 2,
		9, 2, 8, 1, 9, 1, 8, 2, 6, 1, 8, 6, 1, 2, 5, 8, 6, 7, 3, 2, 1, 5,
		7, 9, 1, 9, 8, 4, 1, 4, 8, 4, 8, 8, 2, 9, 1, 6, 4, 4, 7, 0, 6, 0,
		9, 5, 7, 5, 2, 7, 0, 6, 9, 5, 7, 2, 2, 0, 9, 1, 7, 5, 6, 7, 1, 1,
		6, 7, 2, 2, 9, 1, 0, 9, 8, 1, 6, 9, 0, 9, 1, 5, 2, 8, 0, 1, 7, 3,
		5, 0, 6, 7, 1, 2, 7, 4, 8, 5, 8, 3, 2, 2, 2, 8, 7, 1, 8, 3, 5, 2,
		0, 9, 3, 5, 3, 9, 6, 5, 7, 2, 5, 1, 2, 1, 0, 8, 3, 5, 7, 9, 1, 5,
		1, 3, 6, 9, 8, 8, 2, 0, 9, 1, 4, 4, 4, 2, 1, 0, 0, 6, 7, 5, 1, 0,
		3, 3, 4, 6, 7, 1, 1, 0, 3, 1, 4, 1, 2, 6, 7, 1, 1, 1, 3, 6, 9, 9,
		0, 8, 6, 5, 8, 5, 1, 6, 3, 9, 8, 3, 1, 5, 0, 1, 9, 7, 0, 1, 6, 5,
		1, 5, 1, 1, 6, 8, 5, 1, 7, 1, 4, 3, 7, 6, 5, 7, 6, 1, 8, 3, 5, 1,
		5, 5, 6, 5, 0, 8, 8, 4, 9, 0, 9, 9, 8, 9, 8, 5, 9, 9, 8, 2, 3, 8,
		7, 3, 4, 5, 5, 2, 8, 3, 3, 1, 6, 3, 5, 5, 0, 7, 6, 4, 7, 9, 1, 8,
		5, 3, 5, 8, 9, 3, 2, 2, 6, 1, 8, 5, 4, 8, 9, 6, 3, 2, 1, 3, 2, 9,
		3, 3, 0, 8, 9, 8, 5, 7, 0, 6, 4, 2, 0, 4, 6, 7, 5, 2, 5, 9, 0, 7,
		0, 9, 1, 5, 4, 8, 1, 4, 1, 6, 5, 4, 9, 8, 5, 9, 4, 6, 1, 6, 3, 7,
		1, 8, 0, 2, 7, 0, 9, 8, 1, 9, 9, 4, 3, 0, 9, 9, 2, 4, 4, 8, 8, 9,
		5, 7, 5, 7, 1, 2, 8, 2, 8, 9, 0, 5, 9, 2, 3, 2, 3, 3, 2, 6, 0, 9,
		7, 2, 9, 9, 7, 1, 2, 0, 8, 4, 4, 3, 3, 5, 7, 3, 2, 6, 5, 4, 8, 9,
		3, 8, 2, 3, 9, 1, 1, 9, 3, 2, 5, 9, 7, 4, 6, 3, 6, 6, 7, 3, 0, 5,
		8, 3, 6, 0, 4, 1, 4, 2, 8, 1, 3, 8, 8, 3, 0, 3, 2, 0, 3, 8, 2, 4,
		9, 0, 3, 7, 5, 8, 9, 8, 5, 2, 4, 3, 7, 4, 4, 1, 7, 0, 2, 9, 1, 3,
		2, 7, 6, 5, 6, 1, 8, 0, 9, 3, 7, 7, 3, 4, 4, 4, 0, 3, 0, 7, 0, 7,
		4, 6, 9, 2, 1, 1, 2, 0, 1, 9, 1, 3, 0, 2, 0, 3, 3, 0, 3, 8, 0, 1,
		9, 7, 6, 2, 1, 1, 0, 1, 1, 0, 0, 4, 4, 9, 2, 9, 3, 2, 1, 5, 1, 6,
		0, 8, 4, 2, 4, 4, 4, 8, 5, 9, 6, 3, 7, 6, 6, 9, 8, 3, 8, 9, 5, 2,
		2, 8, 6, 8, 4, 7, 8, 3, 1, 2, 3, 5, 5, 2, 6, 5, 8, 2, 1, 3, 1, 4,
		4, 9, 5, 7, 6, 8, 5, 7, 2, 6, 2, 4, 3, 3, 4, 4, 1, 8, 9, 3, 0, 3,
		9, 6, 8, 6, 4, 2, 6, 2, 4, 3, 4, 1, 0, 7, 7, 3, 2, 2, 6, 9, 7, 8,
		0, 2, 8, 0, 7, 3, 1, 8, 9, 1, 5, 4, 4, 1, 1, 0, 1, 0, 4, 4, 6, 8,
		2, 3, 2, 5, 2, 7, 1, 6, 2, 0, 1, 0, 5, 2, 6, 5, 2, 2, 7, 2, 1, 1,
		1, 6, 6, 0, 3, 9, 6, 6, 6, 5, 5, 7, 3, 0, 9, 2, 5, 4, 7, 1, 1, 0,
		5, 5, 7, 8, 5, 3, 7, 6, 3, 4, 6, 6, 8, 2, 0, 6, 5, 3, 1, 0, 9, 8,
		9, 6, 5, 2, 6, 9, 1, 8, 6, 2, 0, 5, 6, 4, 7, 6, 9, 3, 1, 2, 5, 7,
		0, 5, 8, 6, 3, 5, 6, 6, 2, 0, 1, 8, 5, 5, 8, 1, 0, 0, 7, 2, 9, 3,
		6, 0, 6, 5, 9, 8, 7, 6, 4, 8, 6, 1, 1, 7, 9, 1, 0, 4, 5, 3, 3, 4,
		8, 8, 5, 0, 3, 4, 6, 1, 1, 3, 6, 5, 7, 6, 8, 6, 7, 5, 3, 2, 4, 9,
		4, 4, 1, 6, 6, 8, 0, 3, 9, 6, 2, 6, 5, 7, 9, 7, 8, 7, 7, 1, 8, 5,
		5, 6, 0, 8, 4, 5, 5, 2, 9, 6, 5, 4, 1, 2, 6, 6, 5, 4, 0, 8, 5, 3,
		0, 6, 1, 4, 3, 4, 4, 4, 3, 1, 8, 5, 8, 6, 7, 6, 9, 7, 5, 1, 4, 5,
		6, 6, 1, 4, 0, 6, 8, 0, 0, 7, 0, 0, 2, 3, 7, 8, 7, 7, 6, 5, 9, 1,
		3, 4, 4, 0, 1, 7, 1, 2, 7, 4, 9, 4, 7, 0, 4, 2, 0, 5, 6, 2, 2, 3,
		0, 5, 3, 8, 9, 9, 4, 5, 6, 1, 3, 1, 4, 0, 7, 1, 1, 2, 7, 0, 0, 0,
		4, 0, 7, 8, 5, 4, 7, 3, 3, 2, 6, 9, 9, 3, 9, 0, 8, 1, 4, 5, 4, 6,
		6, 4, 6, 4, 5, 8, 8, 0, 7, 9, 7, 2, 7, 0, 8, 2, 6, 6, 8, 3, 0, 6,
		3, 4, 3, 2, 8, 5, 8, 7, 8, 5, 6, 9, 8, 3, 0, 5, 2, 3, 5, 8, 0, 8,
		9, 3, 3, 0, 6, 5, 7, 5, 7, 4, 0, 6, 7, 9, 5, 4, 5, 7, 1, 6, 3, 7,
		7, 5, 2, 5, 4, 2, 0, 2, 1, 1, 4, 9, 5, 5, 7, 6, 1, 5, 8, 1, 4, 0,
		0, 2, 5, 0, 1, 2, 6, 2, 2, 8, 5, 9, 4, 1, 3, 0, 2, 1, 6, 4, 7, 1,
		5, 5, 0, 9, 7, 9, 2, 5, 9, 2, 3, 0, 9, 9, 0, 7, 9, 6, 5, 4, 7, 3,
		7, 6, 1, 2, 5, 5, 1, 7, 6, 5, 6, 7, 5, 1, 3, 5, 7, 5, 1, 7, 8, 2,
		9, 6, 6, 6, 4, 5, 4, 7, 7, 9, 1, 7, 4, 5, 0, 1, 1, 2, 9, 9, 6, 1,
		4, 8, 9, 0, 3, 0, 4, 6, 3, 9, 9, 4, 7, 1, 3, 2, 9, 6, 2, 1, 0, 7,
		3, 4, 0, 4, 3, 7, 5, 1, 8, 9, 5, 7, 3, 5, 9, 6, 1, 4, 5, 8, 9, 0,
		1, 9, 3, 8, 9, 7, 1, 3, 1, 1, 1, 7, 9, 0, 4, 2, 9, 7, 8, 2, 8, 5,
		6, 4, 7, 5, 0, 3, 2, 0, 3, 1, 9, 8, 6, 9, 1, 5, 1, 4, 0, 2, 8, 7,
		0, 8, 0, 8, 5, 9, 9, 0, 4, 8, 0, 1, 0, 9, 4, 1, 2, 1, 4, 7, 2, 2,
		1, 3, 1, 7, 9, 4, 7, 6, 4, 7, 7, 7, 2, 6, 2, 2, 4, 1, 4, 2, 5, 4,
		8, 5, 4, 5, 4, 0, 3, 3, 2, 1, 5, 7, 1, 8, 5, 3, 0, 6, 1, 4, 2, 2,
		8, 8, 1, 3, 7, 5, 8, 5, 0, 4, 3, 0, 6, 3, 3, 2, 1, 7, 5, 1, 8, 2,
		9, 7, 9, 8, 6, 6, 2, 2, 3, 7, 1, 7, 2, 1, 5, 9, 1, 6, 0, 7, 7, 1,
		6, 6, 9, 2, 5, 4, 7, 4, 8, 7, 3, 8, 9, 8, 6, 6, 5, 4, 9, 4, 9, 4,
		5, 0, 1, 1, 4, 6, 5, 4, 0, 6, 2, 8, 4, 3, 3, 6, 6, 3, 9, 3, 7, 9,
		0, 0, 3, 9, 7, 6, 9, 2, 6, 5, 6, 7, 2, 1, 4, 6, 3, 8, 5, 3, 0, 6,
		7, 3, 6, 0, 9, 6, 5, 7, 1, 2, 0, 9, 1, 8, 0, 7, 6, 3, 8, 3, 2, 7,
		1, 6, 6, 4, 1, 6, 2, 7, 4, 8, 8, 8, 8, 0, 0, 7, 8, 6, 9, 2, 5, 6,
		0, 2, 9, 0, 2, 2, 8, 4, 7, 2, 1, 0, 4, 0, 3, 1, 7, 2, 1, 1, 8, 6,
		0, 8, 2, 0, 4, 1, 9, 0, 0, 0, 4, 2, 2, 9, 6, 6, 1, 7, 1, 1, 9, 6,
		3, 7, 7, 9, 2, 1, 3, 3, 7, 5, 7, 5, 1, 1, 4, 9, 5, 9, 5, 0, 1, 5,
		6, 6, 0, 4, 9, 6, 3, 1, 8, 6, 2, 9, 4, 7, 2, 6, 5, 4, 7, 3, 6, 4,
		2, 5, 2, 3, 0, 8, 1, 7, 7, 0, 3, 6, 7, 5, 1, 5, 9, 0, 6, 7, 3, 5,
		0, 2, 3, 5, 0, 7, 2, 8, 3, 5, 4, 0, 5, 6, 7, 0, 4, 0, 3, 8, 6, 7,
		4, 3, 5, 1, 3, 6, 2, 2, 2, 2, 4, 7, 7, 1, 5, 8, 9, 1, 5, 0, 4, 9,
		5, 3, 0, 9, 8, 4, 4, 4, 8, 9, 3, 3, 3, 0, 9, 6, 3, 4, 0, 8, 7, 8,
		0, 7, 6, 9, 3, 2, 5, 9, 9, 3, 9, 7, 8, 0, 5, 4, 1, 9, 3, 4, 1, 4,
		4, 7, 3, 7, 7, 4, 4, 1, 8, 4, 2, 6, 3, 1, 2, 9, 8, 6, 0, 8, 0, 9,
		9, 8, 8, 8, 6, 8, 7, 4, 1, 3, 2, 6, 0, 4, 7, 2}
	numacc1 = stats.Float64Data{10000001, 10000003, 10000002}
	numacc2 = make(stats.Float64Data, 1001)
	numacc3 = make(stats.Float64Data, 1001)
	numacc4 = make(stats.Float64Data, 1001)
)

func init() {
	numacc2[0] = 1.2
	numacc3[0] = 1000000.2
	numacc4[0] = 10000000.2
	for i := 1; i < 1000; i += 2 {
		numacc2[i] = 1.1
		numacc2[i+1] = 1.3
		numacc3[i] = 1000000.1
		numacc3[i+1] = 1000000.3
		numacc4[i] = 10000000.1
		numacc4[i+1] = 10000000.3
	}
}

func TestLewData(t *testing.T) {
	r, e := stats.Mean(lew)
	test("Lew Mean", r, -177.435000000000, 1e-15, e, t)

	r, e = stats.StandardDeviationSample(lew)
	test("Lew Standard Deviation", r, 277.332168044316, 1e-15, e, t)

	r, e = stats.AutoCorrelation(lew, 1)
	test("Lew AutoCorrelate1", r, -0.307304800605679, 1e-14, e, t)
}

func TestLotteryData(t *testing.T) {
	r, e := stats.Mean(lottery)
	test("Lottery Mean", r, 518.958715596330, 1e-15, e, t)

	r, e = stats.StandardDeviationSample(lottery)
	test("Lottery Standard Deviation", r, 291.699727470969, 1e-15, e, t)

	r, e = stats.AutoCorrelation(lottery, 1)
	test("Lottery AutoCorrelate1", r, -0.120948622967393, 1e-14, e, t)
}

func TestMavroData(t *testing.T) {
	r, e := stats.Mean(mavro)
	test("Mavro Mean", r, 2.00185600000000, 1e-15, e, t)

	r, e = stats.StandardDeviationSample(mavro)
	test("Mavro Standard Deviation", r, 0.000429123454003053, 1e-13, e, t)

	r, e = stats.AutoCorrelation(mavro, 1)
	test("Mavro AutoCorrelate1", r, 0.937989183438248, 1e-13, e, t)
}

func TestMichelsonData(t *testing.T) {
	r, e := stats.Mean(michelson)
	test("Michelson Mean", r, 299.852400000000, 1e-15, e, t)

	r, e = stats.StandardDeviationSample(michelson)
	test("Michelson Standard Deviation", r, 0.0790105478190518, 1e-13, e, t)

	r, e = stats.AutoCorrelation(michelson, 1)
	test("Michelson AutoCorrelate1", r, 0.535199668621283, 1e-13, e, t)
}

func TestPidigitsData(t *testing.T) {
	r, e := stats.Mean(pidigits)
	test("Pidigits Mean", r, 4.53480000000000, 1e-14, e, t)

	r, e = stats.StandardDeviationSample(pidigits)
	test("Pidigits Standard Deviation", r, 2.86733906028871, 1e-14, e, t)

	r, e = stats.AutoCorrelation(pidigits, 1)
	test("Pidigits AutoCorrelate1", r, -0.00355099287237972, 1e-13, e, t)
}

func TestNumacc1Data(t *testing.T) {
	r, e := stats.Mean(numacc1)
	test("numacc1 Mean", r, 10000002.0, 1e-14, e, t)

	r, e = stats.StandardDeviationSample(numacc1)
	test("numacc1 Standard Deviation", r, 1.0, 1e-13, e, t)

	r, e = stats.AutoCorrelation(numacc1, 1)
	test("Lew AutoCorrelateNumacc1", r, -0.5, 1e-15, e, t)

}

func TestNumacc2Data(t *testing.T) {
	r, e := stats.Mean(numacc2)
	test("numacc2 Mean", r, 1.2, 1e-10, e, t)

	r, e = stats.StandardDeviationSample(numacc2)
	test("numacc2 Standard Deviation", r, 0.1, 1e-10, e, t)

	r, e = stats.AutoCorrelation(numacc2, 1)
	test("Lew AutoCorrelateNumacc2", r, -0.999, 1e-10, e, t)
}

func TestNumacc3Data(t *testing.T) {
	r, e := stats.Mean(numacc3)
	test("numacc3 Mean", r, 1000000.2, 1e-15, e, t)

	r, e = stats.StandardDeviationSample(numacc3)
	test("numacc3 Standard Deviation", r, 0.1, 1e-9, e, t)

	r, e = stats.AutoCorrelation(numacc3, 1)
	test("Lew AutoCorrelateNumacc3", r, -0.999, 1e-10, e, t)
}

func TestNumacc4Data(t *testing.T) {
	r, e := stats.Mean(numacc4)
	test("numacc4 Mean", r, 10000000.2, 1e-10, e, t)

	r, e = stats.StandardDeviationSample(numacc4)
	test("numacc4 Standard Deviation", r, 0.1, 1e-7, e, t)

	r, e = stats.AutoCorrelation(numacc4, 1)
	test("Lew AutoCorrelateNumacc4", r, -0.999, 1e-7, e, t)
}

// nistVarianceDataset holds NIST certified values for variance testing.
// Certified sample variance = stddev^2 from the NIST reference.
type nistVarianceDataset struct {
	name      string
	data      stats.Float64Data
	stddev    float64
	stddevTol float64
}

var nistVarianceDatasets = []nistVarianceDataset{
	{"Lew", lew, 277.332168044316, 1e-13},
	{"Lottery", lottery, 291.699727470969, 1e-15},
	{"Mavro", mavro, 0.000429123454003053, 1e-11},
	{"Michelson", michelson, 0.0790105478190518, 1e-13},
	{"Pidigits", pidigits, 2.86733906028871, 1e-14},
	{"NumAcc1", numacc1, 1.0, 1e-13},
	{"NumAcc2", numacc2, 0.1, 1e-10},
	{"NumAcc3", numacc3, 0.1, 1e-9},
	{"NumAcc4", numacc4, 0.1, 1e-7},
}

func TestNistSampleVariance(t *testing.T) {
	for _, ds := range nistVarianceDatasets {
		r, e := stats.SampleVariance(ds.data)
		certifiedVar := ds.stddev * ds.stddev
		test(ds.name+" SampleVariance", r, certifiedVar, ds.stddevTol, e, t)
	}
}

func TestNistPopulationVariance(t *testing.T) {
	for _, ds := range nistVarianceDatasets {
		r, e := stats.PopulationVariance(ds.data)
		n := float64(ds.data.Len())
		certifiedPopVar := ds.stddev * ds.stddev * (n - 1) / n
		test(ds.name+" PopulationVariance", r, certifiedPopVar, ds.stddevTol, e, t)
	}
}

// NIST StRD Linear Regression: Norris dataset
// https://www.itl.nist.gov/div898/strd/lls/data/Norris.shtml
// Certified values: B0 = -0.262323073774029, B1 = 1.00211681802045, R^2 = 0.999993745883712
var (
	norrisX = stats.Float64Data{
		0.2, 337.4, 118.2, 884.6, 10.1, 226.5, 666.3, 996.3, 448.6, 777.0,
		558.2, 0.4, 0.6, 775.5, 666.9, 338.0, 447.5, 11.6, 556.0, 228.1,
		995.8, 887.6, 120.2, 0.3, 0.3, 556.8, 339.1, 887.2, 999.0, 779.0,
		11.1, 118.3, 229.2, 669.1, 448.9, 0.5,
	}
	norrisY = stats.Float64Data{
		0.1, 338.8, 118.1, 888.0, 9.2, 228.1, 668.5, 998.5, 449.1, 778.9,
		559.2, 0.3, 0.1, 778.1, 668.8, 339.3, 448.9, 10.8, 557.7, 228.3,
		998.0, 888.8, 119.6, 0.3, 0.6, 557.6, 339.3, 888.0, 998.5, 778.9,
		10.2, 117.6, 228.9, 668.4, 449.2, 0.2,
	}
)

func TestNistLinearRegressionNorris(t *testing.T) {
	coords := make([]stats.Coordinate, len(norrisX))
	for i := range norrisX {
		coords[i] = stats.Coordinate{X: norrisX[i], Y: norrisY[i]}
	}

	regressions, err := stats.LinearRegression(coords)
	if err != nil {
		t.Fatalf("LinearRegression: %v", err)
	}

	// Certified coefficients
	b0 := -0.262323073774029
	b1 := 1.00211681802045

	for i, c := range regressions {
		expected := b0 + b1*norrisX[i]
		test("Norris LinearRegression", c.Y, expected, 1e-10, nil, t)
	}
}

func TestNistPearsonNorris(t *testing.T) {
	// Certified R^2 = 0.999993745883712, so R = sqrt(R^2)
	certifiedR := math.Sqrt(0.999993745883712)

	r, e := stats.Pearson(norrisX, norrisY)
	test("Norris Pearson", r, certifiedR, 1e-10, e, t)

	r2, e := stats.Correlation(norrisX, norrisY)
	test("Norris Correlation", r2, certifiedR, 1e-10, e, t)
}

func TestNistCovarianceNorris(t *testing.T) {
	// Covariance(X,Y) = Pearson(X,Y) * StdDev(X) * StdDev(Y)
	// Using NIST certified values:
	// R = sqrt(0.999993745883712), StdDev(X) = certified from data, StdDev(Y) = certified from data
	certifiedR := math.Sqrt(0.999993745883712)
	sdX, _ := stats.StandardDeviationSample(norrisX)
	sdY, _ := stats.StandardDeviationSample(norrisY)
	certifiedCov := certifiedR * sdX * sdY

	r, e := stats.Covariance(norrisX, norrisY)
	test("Norris Covariance", r, certifiedCov, 1e-10, e, t)

	// Population covariance = sample covariance * (n-1)/n
	n := float64(norrisX.Len())
	certifiedPopCov := certifiedCov * (n - 1) / n

	r2, e := stats.CovariancePopulation(norrisX, norrisY)
	test("Norris CovariancePopulation", r2, certifiedPopCov, 1e-10, e, t)
}

func TestNistSpearmanNorris(t *testing.T) {
	// The Norris dataset has a near-perfect monotonic relationship.
	// Spearman rank correlation should be very close to 1.
	r, e := stats.Spearman(norrisX, norrisY)
	if e != nil {
		t.Fatalf("Spearman: %v", e)
	}
	// Near-perfect linear relationship but with some noise, so Spearman
	// is high but not as close to 1 as Pearson.
	if r < 0.99 {
		t.Errorf("Norris Spearman: got %v, want > 0.99", r)
	}
}

// TestNistMedian validates Median against datasets where the answer
// can be independently verified from the sorted data.
func TestNistMedian(t *testing.T) {
	// Pidigits: 5000 digits of pi
	r, e := stats.Median(pidigits)
	test("Pidigits Median", r, 5, 1e-15, e, t)

	// Lew: 200 values, all integers
	r, e = stats.Median(lew)
	test("Lew Median", r, -162, 1e-15, e, t)

	// Mavro: 50 values, median of sorted[24] and sorted[25]
	r, e = stats.Median(mavro)
	test("Mavro Median", r, 2.0018, 1e-15, e, t)
}

// TestNistMinMax validates Min and Max against NIST datasets.
func TestNistMinMax(t *testing.T) {
	r, e := stats.Min(lew)
	test("Lew Min", r, -579, 1e-15, e, t)
	r, e = stats.Max(lew)
	test("Lew Max", r, 300, 1e-15, e, t)

	r, e = stats.Min(lottery)
	test("Lottery Min", r, 4, 1e-15, e, t)
	r, e = stats.Max(lottery)
	test("Lottery Max", r, 999, 1e-15, e, t)

	r, e = stats.Min(mavro)
	test("Mavro Min", r, 2.00130, 1e-15, e, t)
	r, e = stats.Max(mavro)
	test("Mavro Max", r, 2.00270, 1e-15, e, t)

	r, e = stats.Min(michelson)
	test("Michelson Min", r, 299.62, 1e-15, e, t)
	r, e = stats.Max(michelson)
	test("Michelson Max", r, 300.07, 1e-15, e, t)
}

// TestNistQuartiles validates quartile calculations against NIST datasets.
func TestNistQuartiles(t *testing.T) {
	q, err := stats.Quartile(lew)
	if err != nil {
		t.Fatalf("Lew Quartile: %v", err)
	}
	test("Lew Q1", q.Q1, -453, 1e-15, nil, t)
	test("Lew Q2", q.Q2, -162, 1e-15, nil, t)
	test("Lew Q3", q.Q3, 94, 1e-15, nil, t)
}

// TestNistGeometricMean validates geometric mean on a dataset of positive values.
func TestNistGeometricMean(t *testing.T) {
	// Michelson data is all positive values close to 299.x.
	// Certified by computing exp(mean(log(x))) independently.
	r, e := stats.GeometricMean(michelson)
	if e != nil {
		t.Fatalf("Michelson GeometricMean: %v", e)
	}

	var logSum float64
	for _, v := range michelson {
		logSum += math.Log(v)
	}
	expected := math.Exp(logSum / float64(michelson.Len()))
	test("Michelson GeometricMean", r, expected, 1e-14, nil, t)

	// Mavro data: all values close to 2.001x
	r, e = stats.GeometricMean(mavro)
	if e != nil {
		t.Fatalf("Mavro GeometricMean: %v", e)
	}

	logSum = 0
	for _, v := range mavro {
		logSum += math.Log(v)
	}
	expected = math.Exp(logSum / float64(mavro.Len()))
	test("Mavro GeometricMean", r, expected, 1e-14, nil, t)
}

// TestNistHarmonicMean validates harmonic mean on a dataset of positive values.
func TestNistHarmonicMean(t *testing.T) {
	r, e := stats.HarmonicMean(lottery)
	if e != nil {
		t.Fatalf("Lottery HarmonicMean: %v", e)
	}

	// Compute expected: n / sum(1/x)
	var recipSum float64
	for _, v := range lottery {
		recipSum += 1.0 / v
	}
	expected := float64(lottery.Len()) / recipSum
	test("Lottery HarmonicMean", r, expected, 1e-14, nil, t)
}

// TestNistSum validates Sum against NIST datasets.
// Certified sum = mean * n for each dataset.
func TestNistSum(t *testing.T) {
	type sumCase struct {
		name string
		data stats.Float64Data
		mean float64
	}
	cases := []sumCase{
		{"Lew", lew, -177.435000000000},
		{"Lottery", lottery, 518.958715596330},
		{"Mavro", mavro, 2.00185600000000},
		{"Michelson", michelson, 299.852400000000},
	}
	for _, c := range cases {
		r, e := stats.Sum(c.data)
		expected := c.mean * float64(c.data.Len())
		test(c.name+" Sum", r, expected, 1e-10, e, t)
	}
}

func bench(d stats.Float64Data) {
	_, _ = stats.Mean(d)
	_, _ = stats.StdDevS(d)
	_, _ = stats.AutoCorrelation(d, 1)
}

func BenchmarkNistLew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(lew)
	}
}

func BenchmarkNistLottery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(lottery)
	}
}

func BenchmarkNistMavro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(mavro)
	}
}

func BenchmarkNistMichelson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(michelson)
	}
}

func BenchmarkNistPidigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(pidigits)
	}
}

func BenchmarkNistNumacc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(numacc1)
	}
}

func BenchmarkNistNumacc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(numacc2)
	}
}

func BenchmarkNistNumacc3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(numacc3)
	}
}

func BenchmarkNistNumacc4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(numacc4)
	}
}

func BenchmarkNistAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(lew)
		bench(lottery)
		bench(mavro)
		bench(michelson)
		bench(pidigits)
		bench(numacc1)
		bench(numacc2)
		bench(numacc3)
		bench(numacc4)
	}
}

func test(d string, r, a, v float64, e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}

	var failure bool
	if math.IsNaN(r) || math.IsNaN(a) {
		failure = math.IsNaN(r) != math.IsNaN(a)
	} else if math.IsInf(r, 0) || math.IsInf(a, 0) {
		failure = math.IsInf(r, 0) != math.IsInf(a, 0)
	} else if a != 0 {
		failure = math.Abs(r-a)/math.Abs(a) > v
	} else {
		failure = math.Abs(r) > v
	}

	if failure {
		t.Errorf("%s => %v != %v", d, r, a)
	}
}
