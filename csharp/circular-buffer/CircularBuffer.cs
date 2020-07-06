using System;
using System.Collections.Generic;

public class CircularBuffer<T>
{
    private T[] _buffer;
    private Queue<int> _transactionHistory;
    public CircularBuffer(int capacity)
    {
        _buffer = new T[capacity];
        _transactionHistory = new Queue<int>();
    }

    public T Read()
    {
        if (_buffer.Length == 0)
        {
            throw new InvalidOperationException();
        }

        T res = _buffer[_transactionHistory.Dequeue()];
        return res;
    }

    public void Write(T value)
    {
        throw new NotImplementedException("You need to implement this function.");
    }

    public void Overwrite(T value)
    {
        throw new NotImplementedException("You need to implement this function.");
    }

    public void Clear()
    {
        if (_buffer.Length == 0)
        {
            return;
        }
    }
}