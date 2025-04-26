# miniBA — AI-powered Mini Business Analyst

![miniBA Logo](assets/favicon-32x32.png)  

A lightweight web app that helps small and medium‑sized business owners get fast, practical insights — powered by OpenAI’s GPT models.

---

## 🚀 Live Demo

> _(optional)_

Try it out: [https://miniba.yourdomain.com](https://miniba.yourdomain.com)

---

## 🎯 Motivation

Small businesses often need quick, data‑driven feedback without the time or budget for a full consultant. miniBA streamlines the process:

1. Multi‑step form to capture your business overview, offerings, and challenges.
2. GPT‑powered analysis, recommendations, & growth priorities.
3. Export PDF report and securely retrieve past reports via email & password.

---

## 🔑 Key Features

- **Multi‑step form** with inline validation & progress indicator
- **GPT Analysis** (Summary, Analysis, Opportunities, Risks)
- **Recommendations** endpoint for targeted action items
- **PDF export** via html2pdf.js
- **Persistent storage** using SQLite  
- **Secure retrieval** of past reports by email & password

---

## 🛠️ Tech Stack

- **Backend:** Go (`net/http`), SQLite (`github.com/mattn/go-sqlite3`)
- **Frontend:** Vanilla JavaScript, HTML, CSS
- **PDF generation:** [html2pdf.js](https://github.com/eKoopmans/html2pdf.js)
- **AI integration:** OpenAI GPT-4 API

---

## 📦 Getting Started

### Prerequisites

- Go (>= 1.21)  
- OpenAI API key (set `OPENAI_API_KEY` env var)  

### Installation & Run

```bash
# 1. Clone the repo
git clone https://github.com/nurkhannury/miniba.git
cd miniba

# 2. Install dependencies
go mod download

# 3. Build & start
go build -o miniBA .
./miniBA
```

Then open your browser to `http://localhost:8080`.

---

## 📷 Screenshots

1. **Form entry**  
   ![Form screenshot](static/img/form.png)

2. **Analysis result**  
   ![Analysis screenshot](static/img/analysis.png)

3. **History retrieval**  
   ![Retrieve screenshot](static/img/retrieve.png)

---

## 🔮 Future Improvements

- User authentication & encrypted passwords
- Deployable Docker image + CI/CD pipeline
- Additional export formats (Word, PowerPoint)
- Customizable analysis templates
- Role‑based access for teams

---

## 📜 License

This project is licensed under the MIT License — see [LICENSE](LICENSE) for details.

