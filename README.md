#### hashtable

Hashtables map a unique key to a value to give a O(1) lookup performance. 
This implementation assumes that keys and values in a hash are non-negative integers only.

The hashing function used calculates the index within the slice at which the key-value record is stored.
Collisions are handled using a linked list (separate chaining technique).
Based on the loadfactor, hash size is updated to remain within the range 0.25 and 0.75.
The cost of this operation is O(n) for n number of keys in hashtable. However, this cost amortizes over time.
