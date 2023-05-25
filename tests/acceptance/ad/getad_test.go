package ad

// TODO
//func TestFindAllAds(t *testing.T) {
//	ad1 := factory.AnAd()
//	ad2 := factory.AnAd()
//	ad1, _ = dependencies.Ads.SaveAd(context.Background(), ad1)
//	ad2, _ = dependencies.Ads.SaveAd(context.Background(), ad2)
//
//	var response []ad.JSONFindAllAdsResponse
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ads", nil)
//	srv.Engine.ServeHTTP(w, req)
//	json.Unmarshal(w.Body.Bytes(), &response)
//
//	assert.Equal(t, 200, w.Code)
//	assert.True(t, len(response) == 2)
//	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: ad1.Id.String()})
//	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: ad2.Id.String()})
//}
//
//func TestFindAdById(t *testing.T) {
//	adInput := factory.AnAd()
//	adInput, _ = dependencies.Ads.SaveAd(context.Background(), adInput)
//
//	var response ad.JSONFindAdByIdResponse
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ads/"+adInput.Id.String(), nil)
//	srv.Engine.ServeHTTP(w, req)
//	json.Unmarshal(w.Body.Bytes(), &response)
//
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, response, ad.JSONFindAdByIdResponse{
//		Id:          adInput.Id.String(),
//		Title:       adInput.Title.String(),
//		Description: adInput.Description.String(),
//		Price:       adInput.Price.Number(),
//		Date:        adInput.Date.String(),
//	})
//}
