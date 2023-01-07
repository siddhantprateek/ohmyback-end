echo "BASH SCRIPT starting...."
j=1
for f in *.xml; do
    mv -- "$f" "${j}.xml"
    j=$((j+1))
done