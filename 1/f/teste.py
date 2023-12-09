

n = 10
query = [[1, 5, 3], [4, 8, 7], [6, 9, 1]]
init_list = []

for _ in range(n):
    init_list.append(0)

for line in query:
    init = line[0]
    end = line[1]
    k = line[2]
    for i in range(n):
        if i + 1 >= init and i + 1 <= end:
            init_list[i] = init_list[i] + k
    print(init_list)

print(max(init_list))


