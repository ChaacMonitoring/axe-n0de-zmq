package checks

var MemoryCheck = map[string]string{
    "mem-total": "cat /proc/meminfo | grep MemTotal | awk '{print $2,$3}'",
    "mem-free": "cat /proc/meminfo | grep MemFree | awk '{print $2,$3}'",
}
