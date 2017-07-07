echo "------------------"
echo "[Adding new users]"

echo http http://127.0.0.1:3056/user id="user1" age=23
http http://127.0.0.1:3056/user id="user1" age=23

echo http http://127.0.0.1:3056/user id="user2" age=32
http http://127.0.0.1:3056/user id="user2" age=32

echo "------------------"

echo "[Getting recommendations for user]"
echo http http://127.0.0.1:3056/user1/3
http http://127.0.0.1:3056/r/user1/3


echo "[selecting item by user]"

echo http http://127.0.0.1:3056/selectItem/user1/item1
http http://127.0.0.1:3056/selectItem/user1/item1

echo "------------------"

echo "[Getting recommendations for user]"
echo http http://127.0.0.1:3056/user1/3
http http://127.0.0.1:3056/r/user1/3
