package bible

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Input    string
	Expected []BibelVerse
}

func TestParseBibleVerses(t *testing.T) {
	testCases := []TestCase{
		{
			Input: "路加福音8:26",
			Expected: []BibelVerse{
				{"路", 8, 26, 26},
			},
		},
		{
			Input: "路加福音8:26-39",
			Expected: []BibelVerse{
				{"路", 8, 26, 39},
			},
		},
		{
			Input: "诗篇38篇",
			Expected: []BibelVerse{
				{"诗", 38, 1, -1},
			},
		},
		{
			Input: "提摩太后书2:22-3:9",
			Expected: []BibelVerse{
				{"提后", 2, 22, -1},
				{"提后", 3, 1, 9},
			},
		},
		{
			Input: "士师记8:30-9:22、53-57",
			Expected: []BibelVerse{
				{"士", 8, 30, -1},
				{"士", 9, 1, 22},
				{"士", 9, 53, 57},
			},
		},
		{
			Input: "马太福音19:28-20:16",
			Expected: []BibelVerse{
				{"太", 19, 28, -1},
				{"太", 20, 1, 16},
			},
		},
		{
			Input: "约翰福音4:43-54；路加福音4:14-15",
			Expected: []BibelVerse{
				{"约", 4, 43, 54},
				{"路", 4, 14, 15},
			},
		},
		{
			Input: "彼得后书1:16-21，3:1-2、9-18",
			Expected: []BibelVerse{
				{"彼后", 1, 16, 21},
				{"彼后", 3, 1, 2},
				{"彼后", 3, 9, 18},
			},
		},
		{
			Input: "彼得后书1:16-21，3:1-2、9-18，5:4，6",
			Expected: []BibelVerse{
				{"彼后", 1, 16, 21},
				{"彼后", 3, 1, 2},
				{"彼后", 3, 9, 18},
				{"彼后", 5, 4, 4},
				{"彼后", 6, 1, -1},
			},
		},
		{
			Input:    "可怜我们",
			Expected: []BibelVerse{},
		},
		{
			Input:    "5",
			Expected: []BibelVerse{},
		},
		{
			Input: "提摩太后书2:22-5:9",
			Expected: []BibelVerse{
				{"提后", 2, 22, -1},
				{"提后", 3, 1, -1},
				{"提后", 4, 1, -1},
				{"提后", 5, 1, 9},
			},
		},
		{
			Input: "民数记10:11-13，10:33-11:3",
			Expected: []BibelVerse{
				{"民", 10, 11, 13},
				{"民", 10, 33, -1},
				{"民", 11, 1, 3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Input, func(t *testing.T) {
			result := ParseBibleVerses(testCase.Input)
			expectedJSON, _ := json.Marshal(testCase.Expected)
			resultJSON, _ := json.Marshal(result)
			require.Equal(t, string(expectedJSON), string(resultJSON))
		})

	}
}
