import axios from "@/util/request.js";

const auth = {
  login: "auth/login",
};

export function AuthLogin(data) {
  return axios({
    method: "post",
    data: data,
    url: auth.login,
  });
}

