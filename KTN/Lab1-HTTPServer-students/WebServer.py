from socket import *

serverSocket = socket(AF_INET, SOCK_STREAM)


serverPort = 9999

serverSocket.bind(('',serverPort))

serverSocket.listen(1)


while True:
    print('Ready to serve...')
    connectionSocket, addr = serverSocket.accept()
    try:
        message = connectionSocket.recv(1024)
        filename = message.split()[1]
        f = open(filename[1:])
        outputdata = f.read()
        #OK = "HTTP/1.1 200 OK\r\n\r\n"
        OK = "HTTP/1.1 200 OK\r\n\r\n"
        connectionSocket.send(bytes(OK, 'UTF-8'))
        for i in range(0, len(outputdata)):
            connectionSocket.send(bytes(outputdata[i],'UTF-8'))
        connectionSocket.send(bytes("\r\n", 'UTF-8'))
        connectionSocket.close()
    except IOError:

        connectionSocket.send(bytes("HTTP/1.1 404 NOT FOUND\r\n\r\n", 'UTF-8'))
        connectionSocket.send(bytes("<html><head></head><body><h1>404 Not Found</h1></body></html>\r\n", 'UTF-8'))
        #outd
        f = open("error.html")
        outputdata = f.readlines()
        for i in range (0, len(outputdata)):
            print(outdata[i])
            connectionSocket.send(outputdata[i])
        connectionSocket.close()
serverSocket.close()
