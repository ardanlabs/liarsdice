import axios from 'axios'
let fetchedAddress = '0x0BA052bAeb8925Ac8b480a291F75Ff0dD2c4297c'

axios
  .get('id.env')
  .then((res) => {
    fetchedAddress = res.data
  })
  .catch((err) => console.log(err))

export const contractAddress = fetchedAddress
