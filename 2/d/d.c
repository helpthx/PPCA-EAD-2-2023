#include <stdio.h>
#include <stdlib.h>

#define ull unsigned long long

#define troca(x, y) \
    {               \
        Cell temp = x; \
        x = y;      \
        y = temp;      \
    }

#define min(a, b) (a < b ? a : b)
#define max(a, b) (b < a ? a : b)

#define key(x) (x)
#define menor(a, b) (a.v == b.v ? a.k < b.k : a.v < b.v)

#define cmptroca(a, b) \
    if (menor(b, a))   \
    troca(a, b)

typedef struct
{
    int k, v;
} Cell;

int partition(Cell *v, int l, int r)
{
    Cell pivot = v[r];
    int j = l;
    for (int k = l; k < r; k++)
        if (menor(v[k], pivot))
        {
            troca(v[k], v[j]);
            j++;
        }
    troca(v[j], v[r]);
    return j;
}

void quick_sort_merge(Cell *v, int l, int r)
{
    if (r - l <= 32)
        return;

    troca(v[(l + r) / 2], v[r - 1]);
    cmptroca(v[l], v[r - 1]);
    cmptroca(v[l], v[r]);
    cmptroca(v[r - 1], v[r]);

    int m = partition(v, l, r);
    quick_sort_merge(v, l, m - 1);
    quick_sort_merge(v, m + 1, r);
}

void insertion_sort(Cell *v, int l, int r)
{
    for (int i = l + 1, j; i <= r; i++)
    {
        Cell t = v[i];
        for (j = i; j > 0 && menor(t, v[j - 1]); j--)
            v[j] = v[j - 1];
        v[j] = t;
    }
}

void quick_sort(Cell *v, int l, int r)
{
    quick_sort_merge(v, l, r);
    insertion_sort(v, l, r);
}

int main()
{
    char s[1010];
    while (scanf(" %s", s) == 1)
    {
        int fr[130];
        Cell v[130];

        for (int i = 0; i < 130; i++)
            fr[i] = 0;

        for (int i = 0; s[i] != '\0'; i++)
            fr[s[i]]++;
        int k = 0;
        for (int i = 0; i < 130; i++)
            if (fr[i])
                v[k++] = (Cell) {.k = i, .v=fr[i]};
        quick_sort(v, 0, k-1);
        for (int i = 0; i < k; i++)
            printf("%d %d\n", v[i].k, v[i].v);
        printf("\n");
    }

    return 0;
}