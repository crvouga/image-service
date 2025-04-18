package api

// func TestApiImageResize(t *testing.T) {
// 	// Create a mock AppCtx with a mock ProjectDB
// 	ac := &appCtx.AppCtx{
// 		ProjectDB: &mockProjectDB{},
// 	}

// 	// Test cases
// 	tests := []struct {
// 		name           string
// 		url            string
// 		expectedStatus int
// 	}{
// 		{
// 			name:           "Valid request",
// 			url:            "/api/resize?url=https://example.com/image.jpg&width=300&height=200&projectID=test-project-id",
// 			expectedStatus: http.StatusOK,
// 		},
// 		{
// 			name:           "Missing URL",
// 			url:            "/api/resize?width=300&height=200&projectID=test-project-id",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "Missing projectID",
// 			url:            "/api/resize?url=https://example.com/image.jpg&width=300&height=200",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "Invalid width (too small)",
// 			url:            "/api/resize?url=https://example.com/image.jpg&width=0&height=200&projectID=test-project-id",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "Invalid width (too large)",
// 			url:            "/api/resize?url=https://example.com/image.jpg&width=3000&height=200&projectID=test-project-id",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "Invalid height (too small)",
// 			url:            "/api/resize?url=https://example.com/image.jpg&width=300&height=0&projectID=test-project-id",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "Invalid height (too large)",
// 			url:            "/api/resize?url=https://example.com/image.jpg&width=300&height=3000&projectID=test-project-id",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 	}

// 	// Run test cases
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req, err := http.NewRequest("GET", tc.url, nil)
// 			if err != nil {
// 				t.Fatalf("Failed to create request: %v", err)
// 			}

// 			rr := httptest.NewRecorder()
// 			handler := ApiImageResize(ac)
// 			handler.ServeHTTP(rr, req)

// 			if rr.Code != tc.expectedStatus {
// 				t.Errorf("Expected status %d, got %d", tc.expectedStatus, rr.Code)
// 			}
// 		})
// 	}
// }

// // Mock ProjectDB for testing
// type mockProjectDB struct{}

// func (m *mockProjectDB) GetByID(id projectID.ProjectID) (*project.Project, error) {
// 	// Return a mock project for testing
// 	return &project.Project{
// 		ID: id,
// 	}, nil
// }

// func (m *mockProjectDB) GetByCreatedByUserID(userID interface{}) ([]*project.Project, error) {
// 	// Not used in this test
// 	return nil, nil
// }
