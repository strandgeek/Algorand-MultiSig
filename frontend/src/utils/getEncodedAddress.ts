import { encodeAddress } from "algosdk"

export const getEncodedAddress = (addr?: Uint8Array): string => {
  if (!addr) {
    return ''
  }
  return encodeAddress(addr)
}
