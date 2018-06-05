#### hashtable

Hashtables map a unique key to a value. In this implementation, a hash is implemented using a slice and both the key and the value of the hash are integers. A hashing function calculates the index within the slice at which the key-value record is stored. Collisions are handled using a linked list.
