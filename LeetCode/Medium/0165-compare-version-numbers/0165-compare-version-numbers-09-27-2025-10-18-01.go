func compareVersion(version1 string, version2 string) int {
	// Split the versions by "."
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	// Find the length of the longer version
	maxLength := len(v1Parts)
	if len(v2Parts) > maxLength {
		maxLength = len(v2Parts)
	}

	// Compare each part of the version
	for i := 0; i < maxLength; i++ {
		// If parts are shorter, treat them as 0
		v1 := 0
		v2 := 0

		if i < len(v1Parts) {
			v1, _ = strconv.Atoi(v1Parts[i])
		}
		if i < len(v2Parts) {
			v2, _ = strconv.Atoi(v2Parts[i])
		}

		// Compare the parts
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	// If all parts are equal
	return 0
}

