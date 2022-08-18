import axios from 'axios'
let fetchedAddress = '0x87A061ED19dcA76EC5B01643b054f5eae2730a85'

axios
  .get('id.env')
  .then((res) => (fetchedAddress = res.data))
  .catch((err) => console.log(err))

export const contractAddress = fetchedAddress
