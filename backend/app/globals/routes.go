package globals

const ViewIndexRoute = "/"

const ApiLoginRoute = "/api/login"
const ApiGetUserRoute = "/api/getUser"
const ApiRegisterRoute = "/api/register"
const ApiLogoutRoute = "/api/logout"
const ApiFriendsRoute = "/api/friends"
const ApiRemoveFriendsRoute = "/api/friends/remove"
const ApiAvailableFriendRoute = "/api/availableFriends"

//List of endpoints that require auth
var NonAuthorizedOnlyRoutes = map[string]bool{
	ApiLoginRoute:    true,
	ApiRegisterRoute: true,
	ApiGetUserRoute:  true,

	ViewIndexRoute: true,
}
