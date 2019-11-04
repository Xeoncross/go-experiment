# bbolt Update Races

What happens when two different goroutines try to get-or-create a nonexistent
key inside two different transactions at the same time?

It seems like they both have the same view of the world when they start. However,
bolt/bbolt does not allow two `.Update()` calls to run at the same time. Only
one `.Update()` runs at a time and will always see the results of the last
`.Update()`.

This means that in a read-modify-write scenario only the first `.Update()` call
will fail to find the key. Every other `.Update` will see the value from the
last `.Update` call (well, that succeeded anyway).

https://github.com/etcd-io/bbolt#transactions
