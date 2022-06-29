package cspm

//func TestCspmClient_GetResource(t *testing.T) {
//	teardown := setup()
//	defer teardown()
//	cspmClient, err := NewCSPMClient(&ClientOptions{
//		ApiUrl:     server.URL,
//		SslVerify:  false,
//		Schema:     "http",
//		MaxRetries: 3,
//	})
//	assert.Nil(t, err)
//	mux.HandleFunc(resourceEndpoint, func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
//		w.Header().Set(internal.AuthHeader, "foo")
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(``))
//	})
//	resource, err := cspmClient.GetResource("foo")
//	assert.Nil(t, err)
//	assert.Equal(t)
//}
