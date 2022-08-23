import axios from 'axios'
let fetchedAddress = '0x268287f8Ad03Aa9aFA6D00FAd8ea7091ca0fD8E9'

export const getContractAddress = () => {
  axios
    .get('id.env')
    .then((res) => {
      fetchedAddress = res.data
    })
    .catch((err) => console.log(err))
  return fetchedAddress
}
