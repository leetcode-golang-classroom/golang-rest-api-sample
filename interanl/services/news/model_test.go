package news_test

import (
	"testing"

	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/services/news"
)

func TestNewsPostBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         news.PostReqBody
		expectedErr bool
	}{
		{
			name:        "author empty",
			req:         news.PostReqBody{},
			expectedErr: true,
		},
		{
			name: "title empty",
			req: news.PostReqBody{
				Author: "test-author",
			},
			expectedErr: true,
		},
		{
			name: "summary invalid",
			req: news.PostReqBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectedErr: true,
		},
		{
			name: "time invalid",
			req: news.PostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "Invalid",
			},
			expectedErr: true,
		},
		{
			name: "source invalid",
			req: news.PostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
			},
			expectedErr: true,
		},
		{
			name: "tags is empty",
			req: news.PostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://tests-news.com",
			},
			expectedErr: true,
		},
		{
			name: "validate",
			req: news.PostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://tests-news.com",
				Tags:      []string{"test-tag"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()
			if tc.expectedErr && err == nil {
				t.Fatal("expected error but got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Fatal("expected nil but got error")
			}
		})
	}
}
