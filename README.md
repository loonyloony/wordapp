# Wordapp
An app for word practicing vocabulary

# System Design

### Architecture Diagram

![elsa_homework](https://github.com/user-attachments/assets/56ea6356-8e3d-457c-86d9-f1cef0a221c3)

### Component Description
- Mobile: UI for users joining quiz, answering and checking leaderboard
- Load Balancers: distribute network or application traffic across mutiple servers
- Api Gateway: real-time routing and connection handling requirements
- Quiz Service: quiz management and participant status
- Answer Service: receive answers from participants and forward them to processing system real-time
- Score Service: calculating scores based on answers, logging results and updating leaderboard
- Message Queue: real-time messaging system to send events from Answer Service to Score Service
- Redis Cache: real-time temporary data storage to support scores and participant list
- Database: MongoDB - store questions, quiz information and results

### Data Flow
1. Login
2. Participating
   Frontend (mobile) -> api gateway -> quiz service (quiz/join) -> db -> redis cache
3. Answering
   - Frontend (mobile) -> api gateway (websocket) -> get questions (answer service) -> answer questions (answer service)
     -> produce answer (message queue)
   - Score service -> consume ansewers -> calculate score -> update redis cache (leaderboard) -> save db
4. Leaderboard
   Score service -> redis cache (quiz_xxx:score) -> api gateway (websocket) -> frontend

### Technologies and Tools
- Backend: Golang
- Mobile: kotlin, swift
- Cache: Redis
- Message queue: kafka
