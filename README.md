# Eutamias

```
 ______     __  __     ______   ______     __    __     __     ______     ______
 /\  ___\   /\ \/\ \   /\__  _\ /\  __ \   /\ "-./  \   /\ \   /\  __ \   /\  ___\ 
 \ \  __\   \ \ \_\ \  \/_/\ \/ \ \  __ \  \ \ \-./\ \  \ \ \  \ \  __ \  \ \___  \ 
  \ \_____\  \ \_____\    \ \_\  \ \_\ \_\  \ \_\ \ \_\  \ \_\  \ \_\ \_\  \/\_____\ 
    \/_____/   \/_____/     \/_/   \/_/\/_/   \/_/  \/_/   \/_/   \/_/\/_/   \/_____/
```
a IM schema writed in golang
with four server:

### Server
+ Dispatcher(D)
+ Notificer(N)
+ Switcher(S)
+ HttpServer(H)

### Client
+ Messenger(M)

# 初步机制

1. M connect to D, get a IP of N, close this conn
2. M connect to N, hold this conn, do some operation
3. M send msg to N, M get IP of S from N(this notify also send to peer of M)
4. M and peer of M exchange msg in this S

# 各个模块的设想

1. D: all of short connection / load balance
2. N: all of long connection, need send notify msg to each other, load balance of S
3. S: send msg to each other
4. H: http interface for getting some system infomation
5. M: client for test
