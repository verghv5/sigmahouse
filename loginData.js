(function() {

function authorizeUser(username, password) {
  // API call
  return { message: "NothingEntered" }
}

return {
  "login": function(chunk, context, bodies) {
    var username = context.get("username"),
        password = context.get("password"),
        status = authorizeUser(username, password);

    switch(status.message) {
      case "OK":
        return true;
      case "InvalidUserName":
        return chunk.render(bodies.usernameError, context);
      case "InvalidPassword":
        return chunk.render(bodies.passwordError, context;
      case "NothingEntered":
        return chunk.render(bodies.nothingError, context);
    }

    return false;
  }
};

})();
