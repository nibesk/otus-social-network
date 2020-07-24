package params

const ViewIndexRoute = "/"
const ViewLoginRoute = "/login"
const ViewRegisterRoute = "/register"
const ViewFlowRoute = "/flow"

const ApiLoginRoute = "/api/login"
const ApiRegisterRoute = "/api/register"
const ApiLogoutRoute = "/api/logout"
const ApiFriendRoute = "/api/friends"

//List of endpoints that require auth
var AuthorizedOnlyRoutes = map[string]bool{
	ViewFlowRoute:  true,
	ApiFriendRoute: true,
}
