package gatewaymodels

import "github.com/gin-gonic/gin"

const BasicAuthAccount = "admin"
const BasicAuthPassword = "bwW8JQD2vrUMrSVd"

var BasicAuthAccounts = gin.Accounts{
	BasicAuthAccount: BasicAuthPassword,
}
