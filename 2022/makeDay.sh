day=`ls -l | grep -c ^d`
dayNum=$(expr $day)
day=day$dayNum
mkdir $day
cp template/template.go $day/puzzle$dayNum.go
touch $day/sample.txt
sed -i "s/dayx/$day/g" $day/puzzle$dayNum.go
curl https://adventofcode.com/2022/day/$dayNum/input --cookie "session=" > $day/input.txt
