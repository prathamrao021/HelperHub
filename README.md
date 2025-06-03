# HelperHub

<div align="center">
  <img src="https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black" alt="React" />
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Vite-646CFF?style=for-the-badge&logo=vite&logoColor=white" alt="Vite" />
  <img src="https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white" alt="TypeScript" />
  <img src="https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white" alt="Tailwind CSS" />
</div>

<br>

<div align="center">
  <i>HelperHub is a platform that bridges the gap between volunteers and organizations by creating meaningful connections based on location and expertise.</i>
</div>

## 📋 About the Project

HelperHub streamlines the volunteering process by matching passionate individuals with causes that matter to them. This project was focusing on creating an intuitive platform that connects volunteers with organizations based on skills, interests, and location.

## ✨ Features

- **Role-Based User System** - Separate flows for volunteers and organizations
- **Opportunity Discovery** - Advanced filtering by location, date, and keywords
- **Application Management** - Track application status and history
- **Project Dashboard** - Organizations can manage their volunteer projects
- **Secure Authentication** - JWT-based authentication with role-specific permissions
- **Responsive Design** - Optimized for desktop and mobile devices
- **Interactive UI** - Modern, clean interface with light/dark mode support

## 🛠️ Technology Stack

### Frontend
- **Framework**: React 18 with TypeScript
- **Build Tool**: Vite
- **UI Components**: ShadCN (based on Tailwind CSS)
- **State Management**: React Context API
- **Form Handling**: React Hook Form + Zod
- **Testing**: Cypress for E2E, Jest for unit tests

### Backend
- **Framework**: Go Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **Documentation**: Swagger/OpenAPI
- **Authentication**: JWT

## 🚀 Getting Started

### Prerequisites

- Node.js (v18+)
- Go (v1.20+)
- PostgreSQL (v14+)
- Git

### Installation

#### Clone the Repository

```bash
git clone https://github.com/Dhruv-mak/HelperHub.git
cd HelperHub
```

#### Backend Setup

1. Navigate to the backend directory:

```bash
cd backend
```

2. Install Go dependencies:

```bash
go mod tidy
```

3. Set up the PostgreSQL database:

```bash
# Update the connection string in main.go with your PostgreSQL credentials
# dsn := "host=localhost user=postgres password=admin dbname=Helperhub port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
```

4. Start the backend server:

```bash
go run main.go
```

#### Frontend Setup

1. Navigate to the frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm run dev
```

4. The application will be available at `http://localhost:5173`

## 📚 API Documentation

### Generating API Documentation

The project uses Swagger/OpenAPI for API documentation. To generate updated documentation:

1. Make sure you have Swag installed:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate Swagger documentation:

```bash
cd backend
swag init -g main.go -o docs
```

3. Access the documentation through:

```
http://localhost:8080/swagger/index.html
```

## 🧪 Testing

### Backend Testing

To run the backend tests:

```bash
cd backend
go test ./... -v
```

### Frontend Testing

To run the frontend unit tests:

```bash
cd frontend
npm run test
```

For E2E testing with Cypress:

```bash
cd frontend
npm cypress run  # Interactive UI
# OR
npm cypress open   # Headless mode
```

## 👥 Team Members

| Name | Student ID | Role |
|------|------------|------|
| Dhruv Makwana | 67272938 | Frontend Developer |
| Pratham Rao | 43695122 | Backend Developer |
| Akash Balaji | 73539997 | Frontend Developer |
| Nikhil Dinesan | 23060474 | Backend Developer |

## 📝 Project Structure

```
HELPERHUB/
├── backend/              # Go backend code
│   ├── controllers/      # API controllers
│   ├── docs/             # Swagger documentation
│   ├── models/           # Database models
│   ├── routes/           # API routes
│   └── main.go           # Entry point
├── frontend/             # React frontend code
│   ├── cypress/          # cypress files
│   │   ├── e2e/          # e2e test modules
│   │   ├── support/      # config files
│   ├── src/              # Source files
│   │   ├── assets/       # Static assets
│   │   ├── components/   # Reusable components
│   │   ├── contexts/     # React contexts
│   │   ├── hooks/        # Custom React hooks
│   │   ├── lib/          # Utility functions
│   │   ├── pages/        # Page components
│   │   └── App.tsx       # Main application
│   └── index.html        # HTML entry
└── docs/                 # Project documentation
    ├── sprint1/          # Sprint 1 documentation
    ├── sprint2/          # Sprint 2 documentation
    └── sprint3/          # Sprint 3 documentation
```

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Course Instructor: Dr. Alin Dobra
- All the open-source libraries and frameworks that made this project possible
