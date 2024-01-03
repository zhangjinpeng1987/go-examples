### bytes.Buffer memory usage

There is no way to shrink the memory usage for existing bytes.Buffer, neither Reset nor Truncate.

The output of this should be 
```
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 64
bytes buffer capacity after reset 64
bytes buffer capacity 128
bytes buffer capacity after reset 128
bytes buffer capacity 256
bytes buffer capacity after reset 256
bytes buffer capacity 512
bytes buffer capacity after reset 512
bytes buffer capacity 1024
bytes buffer capacity after reset 1024
bytes buffer capacity 2048
bytes buffer capacity after reset 2048
bytes buffer capacity 4096
bytes buffer capacity after reset 4096
bytes buffer capacity 8192
bytes buffer capacity after reset 8192
bytes buffer capacity 16384
bytes buffer capacity after reset 16384
bytes buffer capacity 32768
bytes buffer capacity after reset 32768
bytes buffer capacity 65536
bytes buffer capacity after reset 65536
bytes buffer capacity 131072
bytes buffer capacity after reset 131072
bytes buffer capacity 262144
bytes buffer capacity after reset 262144
bytes buffer capacity 524288
bytes buffer capacity after reset 524288
bytes buffer capacity 1048576
bytes buffer capacity after reset 1048576
bytes buffer capacity 2097152
bytes buffer capacity after reset 2097152
bytes buffer capacity 4194304
bytes buffer capacity after reset 4194304
bytes buffer capacity 4194304
...
bytes buffer capacity after reset 4194304
bytes buffer capacity 4194304
bytes buffer capacity after reset 4194304
```
