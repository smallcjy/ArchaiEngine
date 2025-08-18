#include <fifo.h>

bool test1()
{
    FIFO<int> fifo(2);
    int value;

    // Test push
    if (!fifo.push(1)) return false;
    if (!fifo.push(2)) return false;
    if (fifo.push(3)) return false;  // Should trigger resize

    // Test pop
    if (!fifo.pop(value) || value != 1) return false;
    if (!fifo.pop(value) || value != 2) return false;

    return true;
}

int main()
{
    return 0;
}