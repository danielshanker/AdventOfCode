day=`ls -l | grep -c ^d`
dayNum=$(expr $day)
day=day$dayNum
mkdir $day
cp template/template.go $day/puzzle$dayNum.go
touch $day/sample.txt
sed -i "s/dayx/$day/g" $day/puzzle$dayNum.go
curl https://adventofcode.com/2022/day/$dayNum/input --cookie "session=53616c7465645f5f1ae04f98abf00617f18950eba2e83063e1c365f0ec2b998a0230ded0ff3c389564eb96fc34f955a067d887d67b5267a3830c0eafc0e23578" > $day/input.txt
