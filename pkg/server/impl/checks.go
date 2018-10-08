package impl

func checkNamespacePermissions(isAdmin bool, namespace string, userNamespaces []string) string {
	if isAdmin {
		//Admin has access to all namespaces
		return namespace
	} else {
		//User only has access to list of namespaces received from headers
		for _, userNs := range userNamespaces {
			if userNs == namespace {
				return namespace
			}
		}
	}
	return ""
}
