from socket import *

serverSocket = socket(AF_INET, SOCK_STREAM)


serverPort = 8888

serverSocket.bind(('',serverPort))

serverSocket.listen(1)


while True:
    print('Ready to serve...')
    connectionSocket, addr = serverSocket.accept()
    try:
        message = connectionSocket.recv(1024)
        filepath = message.split()[1]
        f = open(filepath[1:])
        outputdata = f.readlines()
        connectionSocket.send("HTTP/1.1 200 OK\r\n\n")
        for i in range(0, len(outputdata)):
            connectionSocket.send(outputdata[i])
        connectionSocket.send("\r\n")
        connectionSocket.close()
    except IOError:
        connectionSocket.send("HTTP/1.1 404 NOT FOUND\r\n\n")
        connectionSocket.send("<html><head></head><body><h1>404 Not Found</h1></body></html>\r\n")
        connectionSocket.close()
serverSocket.close()
