TCP allows tou to reliably stream data between nodes on a network. This chapter take a deeper dive into the protocol, focusing on the aspects directly influenced by the code we'll write to establish TCP connections and transmit data over those connections. 

## what make TCP reliable?
- TCP is reliable because:
+ it overcomes the effects of packet loss or receiving package out of order due to transmission error or network congestion.

+ TCP adapt its data transfer rate to make sure it transmits data as fast as possible while keeping dropped package to minimum called flow control.

+ TCP also keeps track of received packet and retransmits unacknowledged packets, as necessary.

## working with TCP session

- a TCP session allows tou to deliver a stream of data of any size to a recipient adn receive confirmation that re recipient received data.

### establishing a session with TCP handshake
- a TCP connection uses a three-way handshake to introduce the client to  the server and the server to client. The handshake create an established TCP session over wi=hich the client and server exchange data.

+ before it can establish a TCP session, the server must listen for incoming connections.

+ as the first step of the handshake, the client sends a packet with the synchronize (SYN) flag to the server. this SYN packet informs the server of the client's capabilities and preferred window setting for the rest of the conversion.

+ next, the server responds with its own packet, with both the acknowledgement(ACK) and SYN flags set. the ACK flag tells the client that the server acknowledges receipt of the client's SYN packet. The server's SYN packet tells the client what setting it's agreed to for the duration of conversion.

+ finally, the client replies with ACK packet to acknowledge the server's SYN packet

- completion of the three-way handshake process establishes the TCP session, and nodes may then exchange data.

### acknowledging recept of packets by using their sequence numbers

- each TCP packet contains a sequence number, which the receiver uses to acknowledge receipt of each packet and properly order the packets for presentation to your go application.

+ the client's operating system  determines the initial sequence number and sends it to the server in the client's SYN packet during the handshake. the server acknowledges receipt of the packet by including this sequence number in it ACK packet to client. likewise, the server share its  generated sequence number Y in its SYN packet to the client. The client replies with it ACK to the server

+ an ACK packet uses the sequence number to the sender "I've received all packets up to and including the packet with this sequence number"


### receive buffers and window sizes
+ since TCP allows a single ACK packet to acknowledge the receipt of more than oen incoming packet, the receiver must advertise


### gracefully terminating TCP connection

### handing less graceful termination

## establishing a TCP connection by using Go's standard library