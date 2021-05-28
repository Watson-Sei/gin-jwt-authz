package utils

func InterfaceSliceConversion(interfaceSlice []interface{}) []string {
	stringSlice := make([]string, len(interfaceSlice))
	for i, v := range interfaceSlice {
		stringSlice[i] = v.(string)
	}
	return stringSlice
}

func Every(expectedScopes []string, userScopes []string) bool {
	check := false
	for _, v := range expectedScopes {
		if !includes(userScopes, v) {
			check = false
			break
		} else {
			check = true
		}
	}
	return check
}

func Some(expectedScopes []string, userScopes []string) bool {
	check := false
	for _, v := range expectedScopes {
		if includes(userScopes, v) {
			check = true
		}
	}
	return check
}

func includes(userScopes []string, scope string) bool {
	for _, v := range userScopes {
		if v == scope {
			return true
		}
	}

	return false
}
