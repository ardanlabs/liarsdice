import axios from 'axios'
let fetchedAddress = '0x5c63E35252A807815E1dF566b0eF04DeB1372464'

axios
  .get('id.env')
  .then((res) => (fetchedAddress = res.data))
  .catch((err) => console.log(err))

export const contractAddress = fetchedAddress
