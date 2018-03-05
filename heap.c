#include <stdio.h>

#define MAX_SIZE 1024

// testing

typedef enum {false, true} bool;

typedef struct Heap 
{
    int elements[MAX_SIZE];
    int size;
} Heap;


void Swap(int *a, int *b)
{
    int aux;
    aux = *a;
    *a = *b;
    *b = aux;
}

bool InsertNode(Heap* heap, int value)
{
    if (heap->size == MAX_SIZE)
        return false;


    heap->size++;
    heap->elements[heap->size] = value;

    int index = heap->size;
    int parent_index = heap->size / 2;
    while (parent_index != 0 && heap->elements[index] > heap->elements[parent_index])
    {
        Swap(&heap->elements[index], &heap->elements[parent_index]);
        index = parent_index;
        parent_index = index / 2;
    }

    return true;
}


void PrintHeap(Heap* heap)
{
    int i;

    for (i = 1; i <= heap->size; i++)
        printf("%d ", heap->elements[i]);
    printf("\n");
}

void TopEquilibrate(Heap* heap)
{
    int index = 1;
    int child_index;
    bool swaped;

    do
    {
        swaped = false;
        child_index = index * 2;
        if (heap->elements[child_index] < heap->elements[index * 2 + 1]
            && (index * 2 + 1) <= heap->size)
        {
            child_index = index * 2 + 1;
        }

        if (heap->elements[child_index] > heap->elements[index] &&
            child_index <= heap->size)
        {
            Swap(&heap->elements[index], &heap->elements[child_index]);

            index = child_index;
            swaped = true;
        }

    } while(swaped == true);
}


bool DeleteNode(Heap* heap)
{
    if (heap->size == 0)
        return false;

    heap->elements[1] = heap->elements[heap->size];
    heap->size--;

    TopEquilibrate(heap);

    return true;
}


void HeapSort(Heap* heap)
{
    int i;
    int current_size = heap->size;
    for (i = 1; i <= current_size; i++)
    {
        Swap(&heap->elements[1], &heap->elements[heap->size]);
        heap->size--;
        TopEquilibrate(heap);
    }
    heap->size = current_size;
}
/*
Heap* Heapify(int* array, int n)
{
    int i;
    Heap *heap = new Heap();

    for (i = 0; i < n; i++)
    {
        heap->elements[i+1] = array[i];
        heap->size++;
    }

    for 

    return heap;
}*/

int main()
{   
    int i;
    Heap heap;
    InsertNode(&heap, 10);
    InsertNode(&heap, 11);
    InsertNode(&heap, 13);
    InsertNode(&heap, 12);
    InsertNode(&heap, 12);
    InsertNode(&heap, 16);
    InsertNode(&heap, 11);
    InsertNode(&heap, 21);
    InsertNode(&heap, 64);
    InsertNode(&heap, 33);
    InsertNode(&heap, 11);
    InsertNode(&heap, 2);
    InsertNode(&heap, 44);
    InsertNode(&heap, 88);

    PrintHeap(&heap);

    for (i = 0; i < 1000; i++)
        InsertNode(&heap, i);
    PrintHeap(&heap);    

    for (i = 0; i < 3; i++)
    {
        DeleteNode(&heap);
        PrintHeap(&heap);
    }

    HeapSort(&heap);
    PrintHeap(&heap);

    return 1;
}