export const getTransactionUrl = (txId: string) => {
  return `${process.env.REACT_APP_EXPLORER_BASE_URL}/tx/${txId}`
}

export const getAddressUrl = (address: string) => {
  return `${process.env.REACT_APP_EXPLORER_BASE_URL}/address/${address}`
}
