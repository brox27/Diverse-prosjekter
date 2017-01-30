from sys import stdin, stderr

def best_path(nm, sans):
    # SKRIV DIN KODE HER
    n = len(nm)
    results = [0]*n
    visited = [0]*n
    results[0] = sans[0]
    for i in range(n):
        for j in range(len(nm[i])):
            if nm[i][j] == 1:
                if results[j] < (results[i]*sans[j]):
                    results[j] = (results[i]*sans[j])
                    #print("Result: " + str(j) + " satt til: " + str(results[i]*sans[j] ))
                    visited[j] = i
    #print(results)
    #print(visited)
    if results[len(results)-1] == 0:
        return 0
    end = n-1
    path = []
    while 1:
        path.append(end)
        if end == 0: break
        end = visited[end]
    path.reverse()
    return "-".join(str(y) for y in path)


n = int(stdin.readline())
probabilities = [float(x) for x in stdin.readline().split()]
neighbour_matrix = []
for line in stdin:
    neighbour_row = [0] * n
    neighbours = [int(neighbour) for neighbour in line.split()]
    for neighbour in neighbours:
        neighbour_row[neighbour] = 1
    neighbour_matrix.append(neighbour_row)
print (best_path(neighbour_matrix, probabilities))
