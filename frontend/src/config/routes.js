export const routes = {
  index: `/`,
  login: `/login`,
  register: `/register`,
  flow: `/flow`,
  friends: `/friends`,
  chat: `/chat/:userId`,

  service_users: {
    friends: `/api/users/friends`,
    removeFriends: `/api/users/friends/remove`,
    availableFriends: `/api/users/availableFriends`,
    getUserById: function (userId) {
        return `/api/users/getUser/${userId}`
    },

    login: "/api/users/login",
    getUser: "/api/users/getUser",
    register: "/api/users/register",
    logout: "/api/users/logout"
  },

  service_chat: {
      ws: `/api/chat/ws`,
      getMessages: function (userId) {
          return `/api/chat/messages/${userId}`
      }
  },
};
