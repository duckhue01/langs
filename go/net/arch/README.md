## Architecture
### Connection model
1. connection oriented
- a single connection is established for the session. Two-way communication flow along the connection. When the session is over, the connection is broken (tcp)
2. connection-less
- in a connection-less system, message are sent independant of each other
- connection-less messages may arrived out of order (ip)
- connection oriented transports may be established on top of connection-less ones (tcp/ip)
- connection-less transports may be established on top of connection oriented ones (http/tcp)
  
### Communication models
1. message passing
2. remote procedure call

### Distributed Computing Model
1. peer-to-peer
2. filter
3. client-server

