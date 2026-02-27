package main

type RecommendationsDatabase struct {
	Model struct {
		Algorithm        string  `json:"algorithm"`
		CfWeight         float64 `json:"cf_weight"`
		ContentWeight    float64 `json:"content_weight"`
		SimilarityMetric string  `json:"similarity_metric"`
		RmseCf           float64 `json:"rmse_cf"`
		RmseHybrid       float64 `json:"rmse_hybrid"`
	} `json:"model"`
	Bookworms []struct {
		UserID          string `json:"user_id"`
		Name            string `json:"name"`
		Recommendations []struct {
			Rank           int      `json:"rank"`
			BookID         string   `json:"book_id"`
			Title          string   `json:"title"`
			Author         string   `json:"author"`
			Genres         []string `json:"genres"`
			Year           int      `json:"year"`
			Pages          int      `json:"pages"`
			AvgRating      float64  `json:"avg_rating"`
			PredictedScore float64  `json:"predicted_score"`
		} `json:"recommendations"`
	} `json:"bookworms"`
}
