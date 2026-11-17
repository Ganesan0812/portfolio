package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Education struct {
	Institution string `json:"institution"`
	Location    string `json:"location"`
	Degree      string `json:"degree"`
	Year        string `json:"year"`
}

type SkillGroup struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

type Profile struct {
	Name        string       `json:"name"`
	Title       string       `json:"title"`
	Location    string       `json:"location"`
	About       string       `json:"about"`
	SkillGroups []SkillGroup `json:"skillGroups"`
	Education   []Education  `json:"education"`
	GitHub      string       `json:"github"`
	LinkedIn    string       `json:"linkedin"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
}

type Experience struct {
	ID       int      `json:"id"`
	Role     string   `json:"role"`
	Company  string   `json:"company"`
	Location string   `json:"location"`
	Period   string   `json:"period"`
	Points   []string `json:"points"`
}

type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tech        []string `json:"tech"`
	GitHub      string   `json:"github"`
	Live        string   `json:"live"`
}

type Message struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

var profile = Profile{
	Name:     "Ganesan D",
	Title:    "Golang Developer",
	Location: "Chennai, Tamil Nadu",
	About:    "Golang developer at Flattrade Broking Private Limited, building scalable RESTful APIs and backend microservices. Experienced in concurrent processing with goroutines and channels, database optimization with MySQL/PostgreSQL, and containerization with Docker. Passionate about turning data and ideas into reliable, production-grade products.",
	Phone:    "6381550595",
	Email:    "dganesan124@gmail.com",
	GitHub:   "https://github.com/Ganesan0812",
	LinkedIn: "https://linkedin.com/in/ganesan-d-b85546318/",
	SkillGroups: []SkillGroup{
		{Category: "Languages", Items: []string{"Go (Golang)", "SQL"}},
		{Category: "Frameworks", Items: []string{"Vue.js 2"}},
		{Category: "Databases", Items: []string{"MySQL", "PostgreSQL", "DBMS"}},
		{Category: "Tools", Items: []string{"Git", "GitHub", "Docker", "Postman", "Claude"}},
		{Category: "IDEs", Items: []string{"VS Code", "Antigravity"}},
	},
	Education: []Education{
		{
			Institution: "Dr.M.G.R. Educational and Research Institute",
			Location:    "Chennai, Tamil Nadu",
			Degree:      "B.Tech Information Technology | CGPA 8.02/10",
			Year:        "May 2025",
		},
		{
			Institution: "KSR Matric Public Hr. Sec. School",
			Location:    "Tirupattur, Tamil Nadu",
			Degree:      "Higher Secondary Certificate | 81.00%",
			Year:        "March 2021",
		},
	},
}

var experiences = []Experience{
	{
		ID:       1,
		Role:     "Golang Developer",
		Company:  "Flattrade Broking Private Limited",
		Location: "Chennai, Tamil Nadu",
		Period:   "June 2025 – Present",
		Points: []string{
			"Designed and developed scalable RESTful APIs and backend microservices in Go (Golang).",
			"Implemented concurrent processing using goroutines and channels, improving response time significantly.",
			"Built and optimized database layers with PostgreSQL/MySQL, writing efficient queries and reducing latency.",
			"Wrote unit and integration tests, increasing code coverage and reducing production bugs.",
			"Containerized services with Docker and collaborated on CI/CD pipelines using Git and GitHub Actions.",
			"Participated in code reviews, debugging and production support in an Agile team.",
		},
	},
}

var projects = []Project{
	{
		ID:    1,
		Title: "Plant Disease Prediction (CNN)",
		Description: "Deep learning model identifying plant diseases from 1,74,000+ images with 98% accuracy " +
			"using TensorFlow & Keras. Interactive GUI built with Streamlit for real-time predictions.",
		Tech:   []string{"Python", "TensorFlow", "Keras", "Streamlit", "OpenCV"},
		GitHub: "https://github.com/Ganesan0812",
		Live:   "",
	},
	{
		ID:    2,
		Title: "Loan Status Prediction",
		Description: "ML pipeline applying Random Forest, SVC, and Logistic Regression for loan status prediction. " +
			"Random Forest achieved the highest accuracy with a custom prediction GUI.",
		Tech:   []string{"Python", "Scikit-learn", "Random Forest", "Pandas"},
		GitHub: "https://github.com/Ganesan0812",
		Live:   "",
	},
	{
		ID:    3,
		Title: "Supply Chain Analysis",
		Description: "End-to-end supply chain data analysis across sourcing, manufacturing, transportation, inventory, and sales. " +
			"Data preprocessing in MySQL, analysis in Python, visualization in Power BI.",
		Tech:   []string{"MySQL", "Python", "Power BI", "Pandas"},
		GitHub: "https://github.com/Ganesan0812",
		Live:   "",
	},
}

var certifications = []string{
	"Generative AI Fundamentals – Databricks",
	"Data Analytics & Visualization Job Simulation – Accenture North America",
	"Artificial Intelligence Job Simulation – Cognizant",
	"Data Analysis Essential – Cisco Networking Academy",
	"Excel Skill Job Simulation – JPMorgan Chase & Co.",
	"SQL – HackerRank",
	"TCS ION Career Edge – Young Professional",
}

var (
	mu       sync.Mutex
	messages []Message
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func cors(next http.Handler) http.Handler {
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		origin = "http://localhost:5173"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	var m Message
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&m); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	m.Name, m.Email, m.Message = strings.TrimSpace(m.Name), strings.TrimSpace(m.Email), strings.TrimSpace(m.Message)
	if m.Name == "" || !strings.Contains(m.Email, "@") || m.Message == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name, valid email and message are required"})
		return
	}
	m.CreatedAt = time.Now()
	mu.Lock()
	messages = append(messages, m)
	mu.Unlock()
	log.Printf("new message from %s <%s>", m.Name, m.Email)
	writeJSON(w, http.StatusCreated, map[string]string{"status": "received"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/profile", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, profile)
	})
	mux.HandleFunc("GET /api/projects", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, projects)
	})
	mux.HandleFunc("GET /api/experience", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, experiences)
	})
	mux.HandleFunc("GET /api/certifications", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, certifications)
	})
	mux.HandleFunc("POST /api/contact", contactHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("API running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, cors(mux)))
}
