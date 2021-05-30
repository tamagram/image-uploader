import axios from "axios"

export const imageFetch = async () => {
  const result = await axios.get<any>('http://localhost:5000/images')
  return result.data
}