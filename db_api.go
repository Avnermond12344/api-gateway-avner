package main

func getSchools() []School {
	return GetSchools() // Mock_db. We should Later change to DB.getSchoolsList()
}

func getGrades(schoolID string) []Grade {
	return GetGradesBySchoolID(schoolID) // Mock_db. We should Later change to DB.getGradesList(schoolID)
}

func getEquipment(schoolID string, gradeID string) []Equipment {
	return GetEquipmentList(schoolID, gradeID) // Mock_db. We should Later change to DB.getEquipmentList(schoolID, gradeID)
}

func getUserIDByCredentials(userName, password string) string {
	// Mock_db. We should Later change to DB.getUserID(userName, password)
	for _, user := range MockUsers {
		if user.Username == userName && user.Password == password {
			return user.UserID
		}
	}
	return ""
}

func getUsernameFromUserID(userID string) string {
	// Mock_db. We should Later change to DB.IsValidUserID(userID)
	for _, user := range MockUsers {
		if user.UserID == userID {
			return user.Username
		}
	}
	return ""
}

func getCartByUserID(userID string) []CartEntry {
	//Mock_db. We should later change to DB.GetCartContent(userID)
	cart, ok := MockCarts[userID]
	if !ok {
		return []CartEntry{}
	}
	return cart
}

func saveCart(userID string, cart []CartEntry) {
	MockCarts[userID] = cart // Mock_db. We should later change to DB.setNewCart(userID, cart)
}
