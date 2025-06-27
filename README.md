🔫 DeadShot – Smart Request Logging & Replay for Debugging APIs
DeadShot is a lightweight logging API designed to capture, store, and replay HTTP request/response cycles to help developers debug issues faster and with full context.

It acts as a centralized receiver for logs sent from different services via SDKs (currently available in Go and C#). You can log:

Full request metadata (headers, query params, body)

Responses (headers, body, status code)

Tags, source system info, error details

And later replay these requests for easier reproduction and debugging

🚀 Features
✅ Centralized log ingestion via API

✅ Replay captured HTTP requests

✅ Lightweight SDKs (Go, C#) to integrate in seconds

✅ Ideal for debugging production and staging issues

📦 How It Works
Your app uses the DeadShot SDK to send detailed HTTP logs to the DeadShot API.

DeadShot stores the logs with metadata.

You can search or replay logged requests later via DeadShot's interface or endpoint.


import deadshot "github.com/quantumbytecode/DeadShotGoLib"

log := deadshot.LogModel{
	Method:     "POST",
	URL:        "/api/login",
	Headers:    "Authorization: Bearer ...",
	Body:       `{"user":"hamid"}`,
	StatusCode: 500,
	Source:     "AuthService",
	Error:      "Invalid credentials",
}

client := deadshot.DeadShot{
	EndPoint: "http://deadshot.yourdomain.com/log",
}

_ = client.Send(log)
