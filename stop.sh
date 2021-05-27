lsof -i:2379|awk '{print $2}'|xargs -n1 -I {} kill -9 {}
lsof -i:8888|awk '{print $2}'|xargs -n1 -I {} kill -9 {}

for (( i=8080; i<8090; i=i+1 )); do
  echo $i
  lsof -i:$i|awk '{print $2}'|xargs -n1 -I {} kill -9 {}
done