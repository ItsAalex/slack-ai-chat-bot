# Integration of AI chatbot by using programming language Go
This project integrates the Wolfram Alpha chatbot into Slack and gives you the ability to chat with it through Slack.
This project is created as a final thesis project + writing a scientific paper in the Knowledge Capital of future journal.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- Wolfram Alpha chatbot integration with Slack
- Chat with Wolfram Alpha chatbot through Slack.

## Installation

### Prerequisites
- Slack installed
- Go installed

### Setup

```bash
# Clone the repository
git clone https://github.com/ItsAalex/slack-ai-chat-bot.git

# Navigate to project directory
cd slack-ai-chat-bot
```

## Usage

### Basic Usage

```bash
go run main.go
```

## Configuration

Explain any configuration files, environment variables, or settings:

```bash
# Create a .env file with:
SLACK_BOT_TOKEN="#"
SLACK_APP_TOKEN="#"
WIT_AI_TOKEN="#"
WOLFRAM_APP_ID="#"
```

## Examples

### Example 1: Basic Task

Describe what this example does:

```bash
1. go run main.go
2. tag bot in chat room. Slack will ask you to add it to a room.
3. To interact with a chatbot, use this prompt template: "@bot query - Who is the president of the USA?"
```

Expected output:
```
"The current president of the United States is Donald Trump." (or something like that)
```

## Project Structure

```
project-root/
├── .env           # Environment file with all the tokens
├── go.mod         # Go dependency management file for defining a project with a module path.
├── go.mod         # Go dependency management file to ensure downloading the exact same code each time.
├── main.go        # Main file
└── README.md      # This file
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Testing

```bash
# Run tests
npm test
# or
pytest
```

## Acknowledgments
Big thanks to the good people:
- Mentor Dr.Slavimir Stošović
- Nikola Vukotić

## Contact

Aleksandar Cvetkovic - linkedin.com/in/aleksandar-cvetkovic-7a880b257 - aleksandarcvetkovic756@gmail.com

---

### Screenshots

COMING SOON
