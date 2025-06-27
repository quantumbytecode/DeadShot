# DeadShot 🎯  
### **Smart HTTP Request Logging & Replay for Debugging APIs**  

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org/) 
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/licenses/MIT) 
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)](https://github.com/quantumbytecode/DeadShot/pulls)  

A lightweight **logging API** designed to capture, store, and replay HTTP request/response cycles for effortless debugging. Perfect for diagnosing production issues with full context.  

---

## ✨ **Features**  
✅ **Full Request/Response Logging** – Headers, query params, body, status codes  
✅ **Request Replay** – Reproduce bugs instantly by replaying captured traffic  
✅ **Multi-SDK Support** – Go and C# SDKs (more coming soon!)  
✅ **Centralized Logging** – Send logs to a single DeadShot instance for all services  
✅ **Error Context** – Attach custom tags, source system info, and error details  

---

## 🚀 **Quick Start**  

### **1. Install the SDK**  
```sh
# Go SDK
go get github.com/quantumbytecode/DeadShotGoLib
