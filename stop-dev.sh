# find the services
services="node|sus|guard|backend|atlas" 
pids=$(ss -plunt | grep pid | grep -E "$services" | grep -Eo 'pid=[0-9]*,' | grep -Eo '[0-9]*')

# send signals to gracefully end processes
for pid in "${pids[@]}"; do
    kill -15 $pid 2> /dev/null
done