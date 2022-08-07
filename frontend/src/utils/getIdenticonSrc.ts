import Identicon from "identicon.js";

const getHexCode = (hash : string = ''): string => {
  let result = '';
  for (var i=0; i<hash.length; i++) {
    result += hash.charCodeAt(i).toString(16);
  }
  return result;
}

export const getIdenticonSrc = (hash?: string) => {
  if (!hash || hash.length < 15) {
    return
  }
  var data = new Identicon((getHexCode(hash)), 420).toString();
  return `data:image/png;base64,${data}`
}
