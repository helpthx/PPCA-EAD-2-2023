def array_sum(arr, n):
    output = []
    for i in range(0, n-1, 2):
        output.append(arr[i] + arr[i + 1])
    
    if n == 1:
        output.append(arr[0] + arr[1])
    
    return output

def insertionSort(arr):
 
    # Traverse through 1 to len(arr)
    for i in range(1, len(arr)):
 
        key = arr[i]
 
        # Move elements of arr[0..i-1], that are
        # greater than key, to one position ahead
        # of their current position
        j = i-1
        while j >= 0 and key < arr[j] :
                arr[j + 1] = arr[j]
                j -= 1
        arr[j + 1] = key


size = int(input())
i = 0
elements = []
for i in range(size):
    elements.append(int(input()))
    #elements = elements + (int(input()),)


if not ((size % 2) == 0):
    #elements = elements + (1000000000, )
    elements.append(1000000000)

#insertionSort(lis(elements))
elements = sorted(set(elements))
output = array_sum(elements, len(elements))
elements = sorted(set(elements) | set(output))

n = len(elements)

for i in range(0, n, 4):
     print(elements[i])

print("Elementos:", n)


 