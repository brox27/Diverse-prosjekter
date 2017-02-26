import time
from socket import *

host = "10.22.72.105"
port = 12000
timeout = 1


clientsocket = socket(AF_INET, SOCK_DGRAM)
clientsocket.settimeout(1)

adress = (host, port)

ptime = 0

while ptime<10:
    ptime += 1

    data2 = "ping"
    data = data2.encode('utf-8')

    try:
        sendTime = time.time()
        clientsocket.sendto(data, adress)
        data3, adress = clientsocket.recvfrom(1024)
        dataS = data3.decode('utf-8').upper()
        recTime = time.time()
        print ("msg: ", dataS)
        print ("time difference: ", (sendTime-recTime))
    except:
        print("timed out")
clientsocket.close()
