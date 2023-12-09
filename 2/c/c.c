#include <stdio.h>
#include <stdlib.h>

void count(int n)
{
    if (n == 0)
        return;

    int ans = 0, counter = 0;
    for (int i = 0; i < n; i++)
    {
        int x;
        scanf(" %d",&x);
        if (counter == 0)
        {
            ans = x;
            counter++;
        }
        else
            counter += (ans == x ? 1 : -1);
    }
    printf("%d\n", ans);
    scanf(" %d", &n);
    count(n);
}

int main()
{
    int n;
    scanf(" %d", &n);
    count(n);

    return 0;
}