# Eutamias
a IM schema writed in golang

```
 ______     __  __     ______   ______     __    __     __     ______     ______
 /\  ___\   /\ \/\ \   /\__  _\ /\  __ \   /\ "-./  \   /\ \   /\  __ \   /\  ___\ 
 \ \  __\   \ \ \_\ \  \/_/\ \/ \ \  __ \  \ \ \-./\ \  \ \ \  \ \  __ \  \ \___  \ 
  \ \_____\  \ \_____\    \ \_\  \ \_\ \_\  \ \_\ \ \_\  \ \_\  \ \_\ \_\  \/\_____\ 
    \/_____/   \/_____/     \/_/   \/_/\/_/   \/_/  \/_/   \/_/   \/_/\/_/   \/_____/
```
![logo](/doc/logo/eutamias_logo.jpg)

## Server
+ Dispatcher(D)
+ Notificer(N)
+ Switcher(S)
+ HttpServer(H) (in the next future, for http interface of system, now don not have one)

## Client
+ Client(C)

## Mechanism

1. dispatcher: load-balance of notificer
2. notificer: register in dispatcher ,
3. switcher: register in all notificer, maintain user map of all user, so switcher can switch msg between themn

//1. C login
//    1. connect to D, get a IP of N, close this conn
//    2. C connect to N, hold this conn, do some operation
//    3. get login response from N
//        1. have Sip/port, connect to the S, notify S peerC is online
//        2. have unread msg, pull unreadMsg
//2. C send msg
//    1. C send create session request to N
//    2. if peer C is Online
//        1. N notify peerC with theS info
//        2. peerC connect to theS
//        3. reponse IP/port of theS to C
//        4. C connect to theS
//        5. C & peerC send msg to each other
//        6. C & peerC send msgack to theS
//        7. theS maintain msg that don not get an ack
//        8. if theS do not get hb from C for HEARTBEATTIMEOUT
//            1. theS close the connection
//            2. save the unread msg
//    3. if peer C is Not Online
//        1. reponse IP/port of theS to C
//        2. C connect to theS
//        3. C send msg to theS
//            1. if peerC is still Offline, S save unread msg
//            2. if peerC is Online, send to peerC
//3. Eutamias does not save msg, only N connect to DB

## conception

1. D: all of short connection / load balance
2. N: all of long connection, need send notify msg to each other, load balance of S
3. S: send msg to each other
4. H: http interface for getting some system infomation
5. C: client of this im system

## NOTE
/usr/local/protobuf/bin/protoc -I rpc/ --go_out=plugins=grpc:rpc rpc/BaseDefine.proto rpc/Clientmsg.proto rpc/DNmsg.proto rpc/RpcService.proto
