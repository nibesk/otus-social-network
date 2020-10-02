export const routes = {
  index: `/`,
  login: `/login`,
  register: `/register`,
  flow: `/flow`,
  friends: `/friends`,

  service_users: {
    friends: `/users/friends`,
    removeFriends: `/users/friends/remove`,
    availableFriends: `/users/availableFriends`,

    login: "/users/login",
    getUser: "/users/getUser",
    register: "/users/register",
    logout: "/users/logout"
  }
};
