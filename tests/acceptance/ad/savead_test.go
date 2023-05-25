package ad

//TODO
//func TestSaveAd(t *testing.T) {
//	objectRequest, jsonRequest := factory.AJSONSaveAdRequest()
//	var response ad.JSONSaveAdResponse
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodPost, "/api/v1/ads", bytes.NewReader(jsonRequest))
//	acceptance.Server.Engine.ServeHTTP(w, req)
//	json.Unmarshal(w.Body.Bytes(), &response)
//
//	assert.Equal(t, 201, w.Code)
//	result := acceptance.FindAdById(response.Id)
//	assert.Equal(t, objectRequest.Description, result.Description.String())
//	assert.Equal(t, objectRequest.Title, result.Title.String())
//	assert.Equal(t, objectRequest.Price, result.Price.Number())
//	assert.NotNil(t, result.Date)
//	assert.NotNil(t, result.Id)
//}
