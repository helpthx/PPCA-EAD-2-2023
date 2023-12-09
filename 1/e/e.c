#include <stdio.h>
#include <stdlib.h>

typedef struct Cell {
    int x; 
    int y;
}Cell;

void main() {
    int m;
    int n; 
    int area_torre = 0; 

    int h; 
    int acc = 0;

    scanf("%d %d", &n, &m);
    Cell tower[n*m];
    char field[n][m], aux;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++)  {
            scanf("%c", &aux);
            if  (aux == '\n') 
                scanf("%c", &aux);
            field[i][j] = aux;
            if  (aux == 't') 
                tower[area_torre++] = (Cell){i, j};
        }
    }


    for (int i = 0; i < area_torre; i++) {
        scanf("%d", &h);
        for (int j = tower[i].x - h; j <= tower[i].x + h; j++)
            for (int k = tower[i].y - h; k <= tower[i].y + h; k++)
                if (j >= 0 && j < n && k >= 0 && k < m && field[j][k] == '#') {
                    acc++;
                    field[j][k] = '.';
                }
    }
    printf("%d\n", acc);
    for (int i = 0; i < n; i++) {
        for(int j = 0; j < m; j++)
            printf("%c", field[i][j]);        
        printf("\n");
    }
}