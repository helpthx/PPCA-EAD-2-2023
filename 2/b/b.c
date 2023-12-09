#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ull unsigned long long
#define Item unsigned
#define swap(a, b)  \
    {               \
        Item t = a; \
        a = b;      \
        b = t;      \
    }
#define compare_and_swap(a, b) \
    if (less(b, a))   \
    swap(a, b)

#define key(x) (x)
#define less(a, b) (a < b)

#define minimum(a, b) (a < b ? a : b)
#define max(a, b) (b < a ? a : b)

int partition(Item *v, int l, int r)
{
    Item pivot = v[r];
    int j = l;
    for (int k = l; k < r; k++)
        if (less(v[k], pivot))
        {
            swap(v[k], v[j]);
            j++;
        }
    swap(v[j], v[r]);
    return j;
}

void quick_sort_alternative(Item *v, int l, int r)
{
    if (r - l <= 32)
        return;

    swap(v[(l + r) / 2], v[r - 1]);
    compare_and_swap(v[l], v[r - 1]);
    compare_and_swap(v[l], v[r]);
    compare_and_swap(v[r - 1], v[r]);

    int m = partition(v, l, r);
    quick_sort_alternative(v, l, m - 1);
    quick_sort_alternative(v, m + 1, r);
}

void insertion_sort(Item *v, int l, int r)
{
    for (int i = l + 1, j; i <= r; i++)
    {
        Item t = v[i];
        for (j = i; j > 0 && less(t, v[j - 1]); j--)
            v[j] = v[j - 1];
        v[j] = t;
    }
}

void quick_sort(Item *v, int l, int r)
{
    quick_sort_alternative(v, l, r);
    insertion_sort(v, l, r);
}

void quick_selection(Item *v, int l, int r, int i)
{
    if (i < l || i > r)
        return;

    compare_and_swap(v[(l + r) / 2], v[r]);
    compare_and_swap(v[l], v[(l + r) / 2]);
    compare_and_swap(v[r], v[(l + r) / 2]);

    int m = partition(v, l, r);
    if (m > i)
        quick_selection(v, l, m - 1, i);
    else if (m < i)
        quick_selection(v, m + 1, r, i);
}

int main()
{
    int n, p, x;
    scanf(" %d %d %d", &n, &p, &x);
    unsigned *v = malloc(n * sizeof(unsigned));
    for (int i = 0; i < n; i++)
        scanf(" %u", v + i);

    quick_selection(v, 0, n-1, minimum(n-1, p*x));
    quick_selection(v, 0, n-1, minimum(n-1, (p+1)*x-1));
    quick_sort(v, minimum(n-1, p*x), minimum(n-1, (p+1)*x-1));

    for (int i = p*x; i < minimum(n, (p+1)*x); i++)
        printf("%u\n", v[i]);

    return 0;
}