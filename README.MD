# Going Places API with ChatGPT

This project demonstrates the power of Go for building a RESTful API server. It integrates with a local Redis server for data storage and OpenAI's ChatGPT for enhanced functionality.

## Prerequisites
Go (version 1.16 or higher)
Docker and Docker Compose
An OpenAI API key

## Setting Up

1. Clone the Repository:

```
git clone https://github.com/yourusername/going-places-api.git
cd going-places-api
```

2. Start Redis Server:
```
docker-compose -f docker-compose-redis-server.yaml up 
```

3. Set OpenAI API Key (optional): 
```go 
// update config/config.go
// TODO: put this in environment
const OPENAI_API_KEY="xxxxxxx"
```

4. Run the application
```
go clean
go build
go run main.go

```

The server will start on port 1323.

## Usage

All requests are GET requests and can thus be run from a browser or your REST client of choice.

- Add a location
```
curl -X GET 'http://localhost:1323/add-word?word=Japan'
```
Example output:
```
{
  "message": "thank you for your word",
  "reason": "Experience the perfect blend of tradition and modernity, from breathtaking ancient temples and historic landmarks to futuristic technology and vibrant city life.",
  "word": "Japan"
}
```

- Remove a random location
```
curl -X GET 'http://localhost:1323/remove-word'
```

Example output:
```
{
  "message": "We are happy to give you back a word",
  "reason": "Beware of Godzilla attacks in Tokyo!",
  "word": "Japan - Experience the perfect blend of tradition and modernity, from breathtaking ancient temples and historic landmarks to futuristic technology and vibrant city life."
}
```

## Extra notes

- To create a similar project structure, use script included in project: go-project-creator.sh 
