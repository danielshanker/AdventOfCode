day=`ls -l | grep -c ^d`
dayNum=$(expr $day)
day=day$dayNum
mkdir $day
cp template/template.go $day/puzzle$dayNum.go
touch $day/sample.txt
sed -i "s/dayx/$day/g" $day/puzzle$dayNum.go
curl https://adventofcode.com/2023/day/$dayNum/input --cookie "session=53616c7465645f5fef15a69fbcc39f4edbfa3f87c71d39296ac77725e2cb311f98e9a99a19719258ae8e21ab8b4d5fb4752e06f5a37c33e06be6c496652473be" > $day/input.txt
git ignore $day/input.txt
