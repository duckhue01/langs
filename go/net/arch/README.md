## choosing a network topology

- the organization of nodes in a network is called its topology. A network's topology can be as simple as a single connection two nodes or as complex as a layout of nodes that don't share a direct connection but are nonetheless able to exchange data.

## bandwidth vs latency
- network bandwidth is the amount of data we can send over a network connection in an interval of time

- network latency is a measure of the time that passes between sending a network resource request and receiving a response

## the open system interconnection reference model

- the hierarchal layer of the osi reference model

+ layer 7 - application layer: your network applications and libraries most often interact with the application layer, which is responsible for identifying hosts and retrieving resource.
+ layer 6 presentation layer: The presentation layer prepares data for the network layer when that data is moving down the stack, and it presents data to the application layer when that data moves up the stack. Encryption, decryption, and data encoding are examples of Layer 6 functions.
+ layer 5 - session layer: the session layer manages the connection life cycle between nodes on a network. it's responsible for establishing the connection, managing connection time-outs, coordinating the mode of operation, and terminating the connection.
+ layer 4 - transport layer: the transport layer controls and coordinates the transfer of data between two nodes while maintaining the reliability of the transfer
+ layer 3 - network layer: the network layer is responsible for transmitting data between nodes. It allows you to send data to a network address without having a direct point-to-point connection to the remote node.
+ layer 2 -data link layer: the data link layer handles data transfers between two directly connected node
+ layer 1 - physical layer: The physical layer converts bits from the network stack to electrical, optic, or radio signals suitable for the underlying physical medium and from the physical medium back into bits.