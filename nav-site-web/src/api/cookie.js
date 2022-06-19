import Cookies from "js-cookie";

const TokenName = "token"

export function CookieSet(token, expire) {
  expire = expire / 86400 ;
  Cookies.set(TokenName, token || "", { expires: expire });
}

export function CookieGetToken(){
    let token = Cookies.get(TokenName)
    return token || ""
}

export function CookieRemoveToken(){
  Cookies.remove(TokenName)
}