package globals

const ViewIndexRoute = "/"

const ApiLoginRoute = "/users/login"
const ApiGetUserRoute = "/users/getUser"
const ApiRegisterRoute = "/users/register"
const ApiLogoutRoute = "/users/logout"
const ApiFriendsRoute = "/users/friends"
const ApiRemoveFriendsRoute = "/users/friends/remove"
const ApiAvailableFriendRoute = "/users/availableFriends"

//List of endpoints that require auth
var NonAuthorizedOnlyRoutes = map[string]bool{
	ApiLoginRoute:    true,
	ApiRegisterRoute: true,
	ApiGetUserRoute:  true,

	ViewIndexRoute: true,
}
