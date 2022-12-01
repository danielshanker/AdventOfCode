day=`ls -l | grep -c ^d`
dayNum=$(expr $day + 1)
day=day$dayNum
mkdir $day
cp template.go $day/puzzle$dayNum.go
touch $day/sample.txt
sed -i "s/dayx/$day/g" $day/puzzle$dayNum.go
curl https://adventofcode.com/2022/day/$dayNum/input --cookie "session=53616c7465645f5f1f3ded1c08eac5973d16e1e09b3f4231210d3d523aa1e530fa19245268f82c4c3faa360be509334e254dbca9ee7be617a942afcc31f41b70" > $day/input.txt
