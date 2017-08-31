# dispatcher

1. accept client request of notificer
2. return the appropriate notificer to client (notificer has lowest curconn/maxconn)
3. for load balance info update, notificer notify dispatcher its infomation when it start
4. send heartbeat between notificer and dispatcher, so that we can know notificer is down
