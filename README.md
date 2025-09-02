# LeetCode Assistant Backend API

A Go-based REST API that serves as the backend for the **LeetCode Assistant Chrome Extension**. This service is powered by Google's Gemini AI and provides intelligent hints, solution analysis, code reviews, and chat functionality to help users learn and improve their problem-solving skills directly within their browser.

🔗 **Chrome Extension**: [LeetCode Assistant Extension](https://leetcodeassistant.vercel.app)

## Features

- **Chat Assistant**: Interactive chat with AI for general programming questions and guidance
- **Hint Generation**: Get progressive hints for LeetCode problems without spoiling solutions
- **Solution Analysis**: Analyze submitted solutions for correctness and optimization opportunities
- **Code Review**: Comprehensive code evaluation with scoring across multiple dimensions
- **CORS Support**: Cross-origin resource sharing enabled for frontend integration

## Tech Stack

- **Language**: Go 1.23
- **Framework**: Gin (HTTP web framework)
- **AI Service**: Google Gemini AI
- **Architecture**: Clean architecture with layered design

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── constants/
│   │   └── prompts.go           # AI prompts and constants
│   ├── handlers/
│   │   ├── chat_handler.go      # Chat endpoint handler
│   │   ├── hints_handler.go     # Hints endpoint handler
│   │   ├── review_handler.go    # Code review handler
│   │   └── solution_handler.go  # Solution analysis handler
│   ├── middleware/
│   │   └── middleware.go        # CORS and other middleware
│   ├── routes/
│   │   └── router.go            # Route definitions
│   └── services/
│       ├── genai_service.go     # GenAI interface
│       └── gemini_service.go    # Gemini AI implementation
├── pkg/
│   ├── request/
│   │   └── request.go           # Request DTOs
│   └── response/
│       └── response.go          # Response DTOs
├── Dockerfile                   # Container configuration
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
├── .env.example                 # Environment variables template
└── README.md                    # Project documentation
```

## API Endpoints

### Health Check
- **GET** `/health` - Service health status

### Core Features
- **POST** `/api/v1/chat` - Interactive chat with AI assistant
- **POST** `/api/v1/hints` - Get hints for a problem
- **POST** `/api/v1/codereview` - Get comprehensive code review
- **POST** `/api/v1/solution` - Analyze solution correctness and optimization

## Request/Response Examples

### Chat
```bash
POST /api/v1/chat
{
    "question": "What is the time complexity of binary search?"
}
```

Response:
```json
{
    "success": true,
    "answer": "Binary search has O(log n) time complexity..."
}
```

### Get Hints
```bash
POST /api/v1/hints
{
    "question": "Two Sum problem"
}
```

Response:
```json
{
    "success": true,
    "hints": "Start by thinking about what data structure..."
}
```

### Solution Analysis
```bash
POST /api/v1/solution
{
    "question": "Two Sum problem",
    "solution": "def twoSum(nums, target): ..."
}
```

Response:
```json
{
    "success": true,
    "feedback": "Your solution is correct! The time complexity is..."
}
```

### Code Review
```bash
POST /api/v1/codereview
{
    "code": "function twoSum(nums, target) { ... }",
    "question": "Two Sum problem"
}
```

Response:
```json
{
    "success": true,
    "review": "Summary: Well-structured solution with clear logic...\nScore: 85/100\n..."
}
```

## Setup and Installation

### Prerequisites
- Go 1.23 or higher
- Google Gemini API key

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd lc-assistant-backend-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your actual values
   ```

4. **Configure environment variables**
   ```env
   GEMINI_API_KEY=your_gemini_api_key_here
   APP_ENV=development
   PORT=8080
   ```

5. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```

The server will start on `http://localhost:8080`

### Docker Deployment

1. **Build the Docker image**
   ```bash
   docker build -t lc-assistant-backend .
   ```

2. **Run the container**
   ```bash
   docker run -p 8080:8080 \
     -e GEMINI_API_KEY=your_api_key \
     -e APP_ENV=production \
     lc-assistant-backend
   ```

## Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `GEMINI_API_KEY` | Google Gemini AI API key | Yes | - |
| `APP_ENV` | Application environment | No | development |
| `PORT` | Server port | No | 8080 |

## Getting Gemini API Key

1. Visit [Google AI Studio](https://ai.google.dev/gemini-api/docs/api-key)
2. Create a new project or select existing one
3. Generate an API key
4. Add the key to your `.env` file

## Development

### Code Structure

The application follows clean architecture principles:

- **Handlers**: Handle HTTP requests and responses
- **Services**: Business logic and external API integration  
- **Middleware**: Cross-cutting concerns like CORS
- **Config**: Configuration management
- **DTOs**: Request/response data transfer objects

### AI Integration

The service uses Google's Gemini AI with carefully crafted prompts for:
- Educational guidance without giving away solutions
- Progressive hint generation
- Comprehensive code analysis
- Technical interview-style code reviews

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions:
- Create an issue in the repository
- Check existing documentation
- Review the API examples above

---

**Note**: This is a teaching assistant API designed to guide learning, not provide direct solutions to problems. Happy Learning :)
