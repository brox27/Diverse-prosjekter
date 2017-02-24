import random
from socket import *

serverSocket = socket(AF_INET, SOCK_DGRAM)

serverSocket.bind(("10.22.43.26",12000))

print("ready for pings")

while True:
        rand = random.randint(0,10)
        print(rand)

        message, address = serverSocket.recvfrom(1024)

        messageS = message.decode('utf-8').upper()

        if rand < 4:
                continue

        serverSocket.sendto(messageS.encode('utf-8'), address)
        print (messageS)
