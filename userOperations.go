package main

var user UserModel

func (userModel UserModel) clusterAge(u UserModel) UserModel {
	if u.Age < 15 {
		u.AgeCluster = 0
	} else if 15 < u.Age && u.Age <= 25 {
		u.AgeCluster = 1
	} else if 25 < u.Age && u.Age <= 40 {
		u.AgeCluster = 2
	} else if 40 < u.Age && u.Age <= 50 {
		u.AgeCluster = 3
	} else if 50 < u.Age && u.Age <= 60 {
		u.AgeCluster = 4
	} else if 60 < u.Age {
		u.AgeCluster = 5
	}
	return u
}
