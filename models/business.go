package models

type Business struct {
	Name               string `json:"name"`
	Industry           string `json:"industry"`
	BusinessModel      string `json:"business_model"`
	Location           string `json:"location"`
	SizeStructure      string `json:"size_structure"`
	YearsOperation     string `json:"years_operation"`
	ProductsServices   string `json:"products_services"`
	TargetAudience     string `json:"target_audience"`
	USP                string `json:"usp"`
	Competitors        string `json:"competitors"`
	Challenges         string `json:"challenges"`
	PerformanceMetrics string `json:"performance_metrics"`
	GrowthObjectives   string `json:"growth_objectives"`
	MarketingSales     string `json:"marketing_sales"`
	FinancialSituation string `json:"financial_situation"`
	Email              string `json:"email"`
	Password           string `json:"password"`
}
