def array_sum(arr, arr_dict, n):
    output = []
    for i in range(0, n-1, 2):
        arr_dict[arr[i] + arr[i + 1]] = True
    
    if n == 1:
        arr_dict[arr[0] + arr[1]]

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
elements_dict = {}
for i in range(size):
    elements_dict[(int(input()))] = True
    #elements = elements + (int(input()),)

#insertionSort(lis(elements))
elements = list(elements_dict.keys())

if not ((len(elements_dict.keys()) % 2) == 0):
    #elements = elements + (1000000000, )
    elements_dict[1000000000] = True
elements.sort()
array_sum(elements, elements_dict, len(elements))
elements = list(elements_dict.keys())
elements.sort()

n = len(elements)

for i in range(0, n, 4):
     print(elements[i])

print("Elementos:", n)


 