import time
from socket import *

host = "10.22.43.26"
port = 9999
timeout = 1


clientsocket = socket(AF_INET, SOCK_DGRAM)
clientsocket.settimeout(1)

ptime = 0

while ptime<10:
    ptime += 1

    data = clientsocket.recvfrom(1024)

    try:
        sendTime = time.time()
        clientsocket.sendto(data, adress)
        data, adress = socket.recvfrom(1024)
        recTime = time.time()
        print ("msg: ", data)
        print ("time difference: ", (sendTime-recTime))
    except:
        print("timed out")
clientsocket.close()
