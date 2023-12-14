count="${1:-1000}"
filename="../benchmark/normalization__large"

if [ -f "$filename" ]; then
    rm "$filename"
fi

for ((i=1; i<=count; i++))
do
    echo "k$i=v$i" >> "$filename"
done