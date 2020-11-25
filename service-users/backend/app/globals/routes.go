package globals

const ViewIndexRoute = "/"

const ApiLoginRoute = "/api/users/login"
const ApiGetUserRoute = "/api/users/getUser"
const ApiGetUserByIdRoute = "/api/users/getUser/{userId:[0-9]+}"
const ApiRegisterRoute = "/api/users/register"
const ApiLogoutRoute = "/api/users/logout"
const ApiFriendsRoute = "/api/users/friends"
const ApiRemoveFriendsRoute = "/api/users/friends/remove"
const ApiAvailableFriendRoute = "/api/users/availableFriends"

//List of endpoints that do not require auth
var RoutesWithoutAuth = map[string]bool{
	ApiLoginRoute:    true,
	ApiRegisterRoute: true,
	ApiGetUserRoute:  true,

	ViewIndexRoute: true,
}
