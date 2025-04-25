package prompts

import (
	"fmt"
	"mini_BA/models"
)

func GeneratePrompt(b models.Business) string {
	return fmt.Sprintf(`
Act as an expert business analyst. Based on the data below, generate a structured, insight-rich, and fully action-oriented JSON report to help the business grow and compete. Based on the detailed information below, generate a structured, insightful, and **action-oriented** business report in **valid JSON format**.

The purpose of this report is to **help the business owner understand their current position, discover untapped opportunities, and identify concrete ways to grow, compete, and improve performance** across all areas.

Rules:
- Output only valid, raw JSON (no markdown, no explanation).
- JSON must be syntactically valid and ready for parsing.

Return a JSON object with the following top-level structure:

{
  "Summary": {
   "Business Strengths": "",
   "Biggest Risk": "",
   "Growth Rating": "",
   "Top Priority": ""
  },
  "Analysis": {
    "Market and Industry Insight": "Provide deep analysis of the market trends, demand, and industry trajectory relevant to the business context.",
    "Competitive Positioning": "Explain where this business stands among its competitors, including any strengths or vulnerabilities.",
    "Internal Operations": "Evaluate internal business structure, processes, team, or operational efficiency.",
    "Financial Health": "Analyze financial situation, revenue trends, and any concerns or strengths in financial sustainability."
  },
  "Opportunities": {
    "Growth Potential": "Identify real, data-informed opportunities for business expansion or scaling.",
    "Marketing & Sales Improvements": "Suggest practical marketing, branding, or sales channel improvements.",
    "Service Optimization": "Provide actionable suggestions to enhance products, services, or customer experience."
  },
  "Recommendations": {
    "Short-term Actions": "List immediate, impactful steps the business can take within 1–3 months.",
    "Mid-term Strategy": "Describe strategic moves for the next 6–12 months that improve operations or market position.",
    "Long-term Vision": "Outline visionary ideas that align with sustainable growth and transformation over 1–3 years."
  },
  "Risks and Mitigations": {
    "Market Risks": "Describe external risks such as regulation, demand shifts, or competitors—and how to mitigate them.",
    "Operational Risks": "Flag potential issues in delivery, staffing, or supply—and how to manage or avoid them.",
    "Financial Risks": "Identify financial vulnerabilities and suggest smart controls or risk-reduction strategies."
  }
}

Also include a detailed SWOT section summarizing strengths, weaknesses, opportunities, and threats based on the full report above.

Guidelines:
- You must use **all the provided data** to form insights. Some fields may be absent.
- Focus on **real analysis, not repetition**. Extract patterns, trends, or implications from the data.
- All content should be **actionable** and **specific to this business**.


## Business Overview:
- Business Name: %s
- Industry/Sector: %s
- Business Model: %s
- Location & Market: %s
- Business Size and Structure: %s
- Years in Operation: %s

## Offerings & Market:
- Products/Services Offered: %s
- Target Audience: %s
- Unique Selling Proposition (USP): %s
- Main Competitors: %s

## Internal & External Insights:
- Current Challenges: %s
- Recent Performance Metrics: %s
- Growth Objectives: %s
- Marketing and Sales Channels: %s
- Financial Situation: %s
`,
		b.Name,
		b.Industry,
		b.BusinessModel,
		b.Location,
		b.SizeStructure,
		b.YearsOperation,
		b.ProductsServices,
		b.TargetAudience,
		b.USP,
		b.Competitors,
		b.Challenges,
		b.PerformanceMetrics,
		b.GrowthObjectives,
		b.MarketingSales,
		b.FinancialSituation)
}

func GenerateRecommendationsPrompt(b models.Business) string {
	return fmt.Sprintf(`
You are a senior business consultant.

Based on the following business information, give 10 practical recommendations.

Each recommendation must have:
- A concise title
- A short explanation why it's important
- The impact it will have on the business

Return as valid JSON array:
[
  {
    "Title": "___",
    "Explanation": "___",
	"Impact" : "___"
  },
  ...
]

Avoid headings or bullet points. Just return JSON.

Business:
Name: %s
Industry: %s
Model: %s
Location: %s
Size & Structure: %s
Years: %s
Products: %s
Audience: %s
USP: %s
Competitors: %s
Challenges: %s
Metrics: %s
Goals: %s
Sales: %s
Finance: %s
`,
		b.Name,
		b.Industry,
		b.BusinessModel,
		b.Location,
		b.SizeStructure,
		b.YearsOperation,
		b.ProductsServices,
		b.TargetAudience,
		b.USP,
		b.Competitors,
		b.Challenges,
		b.PerformanceMetrics,
		b.GrowthObjectives,
		b.MarketingSales,
		b.FinancialSituation)
}
